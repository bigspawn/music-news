package internal

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// bbCodeTags lists known BBCode tags to strip from titles.
var bbCodeTags = []string{"bb", "b", "i", "u", "s", "url", "img", "color", "size", "quote", "code"}

func stripBBCode(s string) string {
	for _, tag := range bbCodeTags {
		open := "[" + tag + "]"
		close := "[/" + tag + "]"
		if strings.Contains(s, open) && strings.Contains(s, close) {
			s = strings.ReplaceAll(s, open, "")
			s = strings.ReplaceAll(s, close, "")
		}
	}
	return strings.TrimSpace(s)
}

const siteLabel = "site"

var skipTitlePatterns = []string{
	"discography",
	"дискография",
	"лучшие альбомы",
	"итоги",
	"best of",
	"va -",
	"v.a. -",
	"v/a -",
	"various artists",
	"compilation",
	"сборник",
	"top albums",
	"top releases",
}

func shouldSkipTitle(title string) bool {
	lower := strings.ToLower(title)
	for _, p := range skipTitlePatterns {
		if strings.Contains(lower, p) {
			return true
		}
	}
	return false
}

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

type Parser struct {
	url        string
	feedParser *gofeed.Parser
	store      *Store
	lgr        lgr.L
	itemParser ItemParser
	siteLabel  string
	withDelay  bool
}

func (p *Parser) Parse(ctx context.Context) ([]News, error) {
	feed, err := p.feedParser.ParseURLWithContext(p.url, ctx)
	if err != nil {
		return nil, err
	}

	news := make([]News, 0, len(feed.Items))
	for _, item := range feed.Items {
		if item == nil {
			continue
		}

		item.Title = stripBBCode(item.Title)

		if shouldSkipTitle(item.Title) {
			p.lgr.Logf("[INFO] skip non-release title: %s", item.Title)
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

		if p.withDelay {
			duration := time.Duration(RandBetween(15, 1)) * time.Second

			p.lgr.Logf("[INFO] %s: sleep: sec=%s", p.siteLabel, duration)

			t := time.NewTimer(duration)
			select {
			case <-t.C:
				t.Stop()
			case <-ctx.Done():
				t.Stop()
				return nil, ctx.Err()
			}
		}
	}

	return news, nil
}
