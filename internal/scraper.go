package internal

import (
	"context"
	"errors"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/go-pkgz/lgr"
)

const parsingTimeout = 30 * time.Minute

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

func (s *Scraper) Scrape(ctx context.Context) error {
	if s.withDelay {
		if err := s.wait(ctx); err != nil {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(ctx, parsingTimeout)
	defer cancel()

	items, err := s.parser.Parse(ctx)
	if err != nil {
		s.lgr.Logf("[ERROR] can't parse: err=%v", err)

		if netErr(err) || strings.Contains(err.Error(), "network is unreachable") {
			return err
		}
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

func (s *Scraper) wait(ctx context.Context) error {
	sec := RandBetween(10*60, 1)
	duration := time.Duration(sec) * time.Second

	s.lgr.Logf("[INFO] %s: sleep %s", s.name, duration)

	t := time.NewTimer(duration)
	defer t.Stop()

	select {
	case <-t.C:
		// go
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

//nolint: errorlint // ok
func netErr(err error) bool {
	if ok := errors.Is(err, &net.DNSError{}); ok {
		return ok
	}
	if ok := errors.Is(err, &net.OpError{}); ok {
		return ok
	}
	if ok := errors.Is(err, &net.AddrError{}); ok {
		return ok
	}
	if ok := errors.Is(err, &net.ParseError{}); ok {
		return ok
	}
	if ok := errors.Is(err, &net.DNSConfigError{}); ok {
		return ok
	}
	if ok := errors.Is(err, &url.Error{}); ok {
		return ok
	}
	if _, ok := err.(net.InvalidAddrError); ok {
		return ok
	}
	if _, ok := err.(net.UnknownNetworkError); ok {
		return ok
	}
	if _, ok := err.(url.EscapeError); ok {
		return ok
	}
	if _, ok := err.(url.InvalidHostError); ok {
		return ok
	}
	return false
}
