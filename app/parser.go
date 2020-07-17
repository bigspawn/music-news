package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ExcludeWords     = []string{"Leaked\n"}
	ExcludeLastWords = []string{"Download\n", "Downloads\n", "Total length:", "Download", "Support! Facebook / iTunes"}
	ExcludeGenders   = []string{"pop", "rap", "folk", "synthpop", "r&b", "thrash metal", "J-Core", "R&amp;B"}

	imageCssSelectorPath      = "html.wf-roboto-n3-active.wf-roboto-n4-active.wf-roboto-i4-active.wf-roboto-n7-active.wf-roboto-i3-active.wf-roboto-i7-active.wf-active body.ipsApp.ipsApp_front.ipsJS_has.ipsClearfix.ipsApp_noTouch main#ipsLayout_body.ipsLayout_container.v-nav-wrap div#ipsLayout_contentArea div#ipsLayout_contentWrapper div#ipsLayout_mainArea div.cTopic.ipsClear.ipsSpacer_top div.ipsAreaBackground_light form article#elComment_212133.cPost.ipsBox.ipsComment.ipsComment_parent.ipsClearfix.ipsClear.ipsColumns.ipsColumns_noSpacing.ipsColumns_collapsePhone div.ipsColumn.ipsColumn_fluid div#comment-212133_wrap.ipsComment_content.ipsType_medium.ipsFaded_withHover div.cPost_contentWrap.ipsPad div.ipsType_normal.ipsType_richText.ipsContained p a img.ipsImage"
	imageCssSelectorPathThumb = imageCssSelectorPath + ".ipsImage_thumbnailed"
	imgSelector               = "img.ipsImage.ipsImage_thumbnailed"
	imgSelector2              = "img.ipsImage"
)

var (
	newsParsed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "parsed_news_total",
		Help: "The total number of parsed news",
	})

	errSkipItem = errors.New("skip item")
)

type SiteParser struct {
	FeedParser             *gofeed.Parser
	Store                  *Store
	URL                    string
	SentGauge, ParsedGauge prometheus.Gauge
}

func NewParser(FeedParser *gofeed.Parser, Store *Store, URL string) *SiteParser {
	return &SiteParser{
		FeedParser:  FeedParser,
		Store:       Store,
		URL:         URL,
		SentGauge:   promauto.NewGauge(prometheus.GaugeOpts{Name: "sent_news_gauge"}),
		ParsedGauge: promauto.NewGauge(prometheus.GaugeOpts{Name: "parsed_news_gauge"}),
	}
}

func (p *SiteParser) Parse() ([]*News, error) {
	feed, err := p.FeedParser.ParseURL(p.URL)
	if err != nil {
		return nil, err
	}

	var news []*News
	for _, item := range feed.Items {
		if item == nil {
			continue
		}

		newsParsed.Inc()

		exist, err := p.Store.Exist(item.Title)
		if err != nil {
			return nil, err
		}
		if exist {
			continue
		}

		Lgr.Logf("[INFO] parsing: title=%s, link=%s", item.Title, item.Link)

		n, err := parseItem(item)
		if err != nil {
			if errors.Is(err, errSkipItem) {
				Lgr.Logf("[INFO] skip: title=%v, link=%s", item.Title, item.Link)
				continue
			}

			Lgr.Logf("[ERROR] parsing: item=%v, err=%v", item, err)
			continue
		}

		if err = p.Store.Insert(n); err != nil {
			Lgr.Logf("[ERROR] saving: item=%v, err=%v", item, err)
			continue
		}
		news = append(news, n)
	}

	p.ParsedGauge.Set(float64(len(news)))

	return news, nil
}

func parseItem(item *gofeed.Item) (*News, error) {
	u, err := url.Parse(item.Link)
	if err != nil {
		return nil, err
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	var topic string
	for k, _ := range q {
		if strings.Contains(k, "/forums/topic/") {
			topic = k
			break
		}
	}
	if topic == "" {
		return nil, errors.New("empty topic")
	}

	u.RawQuery = topic
	item.Link = u.String()

	resp, err := http.Get(item.Link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile("[\\n\\t\\s]{2,}")
	desc := re.ReplaceAllString(item.Description, "\n")

	if containText(desc, ExcludeGenders) {
		return nil, errSkipItem
	}

	downloadLink, err := extractDownload(doc)
	if err != nil {
		return nil, err
	}

	imageLink, err := extractImage(doc)
	if err != nil {
		return nil, fmt.Errorf("can't find image link: %v", err)
	}

	n := &News{
		Title:        item.Title,
		DateTime:     item.PublishedParsed,
		Text:         normalize(desc),
		PageLink:     item.Link,
		ImageLink:    imageLink,
		DownloadLink: downloadLink,
	}
	return n, nil
}

func extractDownload(doc *goquery.Document) ([]string, error) {
	link := doc.
		Find("div.ipsType_normal.ipsType_richText.ipsContained").
		First().
		Find("a[rel~='external']").
		First()

	val, exists := link.Attr("href")
	if !exists {
		return nil, errors.New("can't find download link")
	}
	return []string{val}, nil
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
