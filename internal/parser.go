package internal

import (
	"context"
	"errors"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/url"
)

const siteLabel = "site"

var errorCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "music_news",
	Subsystem: "parser",
	Name:      "parsed_error_total",
	Help:      "Total count of parsed item errors",
}, []string{siteLabel})

var (
	ErrAccountIsDisabled = errors.New("account is disabled")
	ErrSkipItem          = errors.New("skip item")
)

type ItemParser interface {
	Parse(ctx context.Context, item *gofeed.Item) (*News, error)
}

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
