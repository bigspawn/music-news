package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pkgz/lgr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

var (
	ExcludeWords     = []string{"Leaked\n"}
	ExcludeLastWords = []string{"Download\n", "Downloads\n", "Total length:", "Download", "Support! Facebook / iTunes"}
)

var ErrSkipItem = errors.New("skip item")

const siteLabel = "site"

var errorCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "music_news",
	Subsystem: "parser",
	Name:      "parsed_error_total",
	Help:      "Total count of parsed item errors",
}, []string{siteLabel})

type RssFeedParser interface {
	Parse(ctx context.Context) ([]News, error)
}

func NewRssFeedParser(rssHost string, store *Store, lgr lgr.L, itemParser ItemParser) RssFeedParser {
	link, err := url.Parse(rssHost)
	if err != nil {
		lgr.Logf("[FATAL] %w+", err)
	}

	return &parser{
		url:        rssHost,
		feedParser: gofeed.NewParser(),
		store:      store,
		lgr:        lgr,
		itemParser: itemParser,
		siteLabel:  link.Host,
	}
}

type parser struct {
	url        string
	feedParser *gofeed.Parser
	store      *Store
	lgr        lgr.L
	itemParser ItemParser
	siteLabel  string
}

func (p parser) Parse(ctx context.Context) ([]News, error) {
	feed, err := p.feedParser.ParseURL(p.url)
	if err != nil {
		return nil, err
	}

	news := make([]News, 0, len(feed.Items))
	for _, item := range feed.Items {
		if item == nil {
			continue
		}

		exist, err := p.store.Exist(ctx, item.Title)
		if err != nil {
			return nil, err
		}
		if exist {
			continue
		}

		p.lgr.Logf("[INFO] parsing: title=%s, link=%s", item.Title, item.Link)

		n, err := p.itemParser.Parse(ctx, item)
		if err != nil {
			if errors.Is(err, ErrSkipItem) {
				p.lgr.Logf("[INFO] skip: title=%v, link=%s", item.Title, item.Link)
				continue
			}
			if errors.Is(err, ErrAccountIsDisabled) {
				p.lgr.Logf("[WARN] skip: %s: %w+", p.url, err)
				return nil, nil
			}

			errorCounter.With(prometheus.Labels{siteLabel: p.siteLabel}).Inc()

			p.lgr.Logf("[ERROR] failed parsing: item=%v, err=%v", item, err)
			continue
		}
		news = append(news, *n)
	}
	return news, err
}

type ItemParser interface {
	Parse(ctx context.Context, item *gofeed.Item) (*News, error)
}

func extractImage(document *goquery.Document) (string, error) {
	imageLink := document.
		Find("div[data-role=commentContent]").
		Find("img[class~=ipsImage]").
		Not("img[src~='https://kingdom-leaks.com/img/lastfm-logo.png']").
		First()

	if link, ok := imageLink.Attr("src"); ok && link != "" {
		return link, nil
	}
	if link, ok := imageLink.Attr("data-imageproxy-source"); ok && link != "" {
		return link, nil
	}
	if link, ok := imageLink.Attr("data-src"); ok && link != "" {
		return link, nil
	}

	return "", errors.New("image link not found")
}

func normalize(description string) string {
	for _, word := range ExcludeLastWords {
		index := strings.LastIndex(description, word)
		if index > 0 {
			description = description[:index]
		}
	}
	for _, word := range ExcludeWords {
		index := strings.Index(description, word)
		if index > -1 {
			description = description[:index] + description[index+utf8.RuneCountInString(word):]
		}
	}
	description = split(description)
	return description
}

const (
	genderTxt            = " Genre - "
	genderNextLineTxt    = "\nGenre - "
	qualityTxt           = " Quality - "
	qualityNextLineTxt   = "\nQuality - "
	trackListTxt         = " Tracklist: "
	trackListNextLineTxt = "\nTracklist:\n"
)

func split(desc string) string {
	isPrefixNeeds := strings.Contains(desc, "01. ")

	max := 100
	n := 1
	for {
		if n == max {
			break
		}

		prefix := ""
		if n < 10 && isPrefixNeeds {
			prefix = "0"
		}

		oldV := fmt.Sprintf(" %s%d. ", prefix, n)
		newV := fmt.Sprintf("\n%s%d. ", prefix, n)
		desc = strings.Replace(desc, oldV, newV, 1)

		n++
	}

	desc = strings.Replace(desc, genderTxt, genderNextLineTxt, 1)
	desc = strings.Replace(desc, qualityTxt, qualityNextLineTxt, 1)
	desc = strings.Replace(desc, trackListTxt, trackListNextLineTxt, 1)
	desc = strings.ReplaceAll(desc, "&nbsp;", "\n")

	return desc
}

func containText(text string, words []string) bool {
	for _, word := range words {
		if strings.Contains(strings.ToLower(text), strings.ToLower(word)) {
			return true
		}
	}
	return false
}
