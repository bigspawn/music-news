package internal

import (
	"context"
	"net"
	"net/url"
	"time"

	"github.com/go-pkgz/lgr"
)

const (
	parsingTimeout  = 30 * time.Minute
	postItemTimeout = 3 * time.Second
)

type MusicScraper interface {
	Scrape(ctx context.Context) error
}

type Scraper struct {
	parser    *Parser
	lgr       lgr.L
	ch        chan<- []News
	store     *Store
	withDelay bool
	name      string
}

func (s Scraper) Scrape(ctx context.Context) error {
	if s.withDelay {
		sec := RandBetween(10*60, 1)
		duration := time.Duration(sec) * time.Second
		s.lgr.Logf("[INFO] %s: sleep %s", s.name, duration)

		t := time.NewTimer(duration)
		defer t.Stop()

		select {
		case <-t.C:
			// go
		case <-ctx.Done():
			return nil
		}
	}

	ctx, cancel := context.WithTimeout(ctx, parsingTimeout)
	defer cancel()

	items, err := s.parser.Parse(ctx)
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			return dnsErr
		}
		if urlErr, ok := err.(*url.Error); ok {
			return urlErr
		}

		s.lgr.Logf("[ERROR] can't parse: err=%v", err)

		return nil
	}

	for _, item := range items {
		if err := s.store.Insert(ctx, item); err != nil {
			s.lgr.Logf("[ERROR] insert to db: %v", err)
			continue
		}
	}

	go func() {
		s.ch <- items
	}()

	return nil
}
