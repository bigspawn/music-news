package internal

import (
	"context"
	"github.com/go-pkgz/lgr"
	"time"
)

const (
	parsingTimeout  = 30 * time.Minute
	postItemTimeout = 3 * time.Second
)

type MusicScraper interface {
	Scrape(ctx context.Context) error
}

func NewMusicScraper(
	parser RssFeedParser,
	lgr lgr.L,
	ch chan<- []News,
	s *Store,
) MusicScraper {
	return &scraper{
		parser: parser,
		lgr:    lgr,
		ch:     ch,
		store:  s,
	}
}

type scraper struct {
	parser RssFeedParser
	lgr    lgr.L
	ch     chan<- []News
	store  *Store
}

func (s scraper) Scrape(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, parsingTimeout)
	defer cancel()

	items, err := s.parser.Parse(ctx)
	if err != nil {
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

func Merge(current []*News, unpublished map[string]*News) {
	if len(unpublished) == 0 {
		return
	}

	for _, c := range current {
		if _, ok := unpublished[c.Title]; ok {
			delete(unpublished, c.Title)
		}
	}

	for _, v := range unpublished {
		current = append(current, v)
	}
}
