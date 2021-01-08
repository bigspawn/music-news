package internal

import (
	"context"
	"github.com/go-pkgz/lgr"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

const (
	parsingTimeout  = 30 * time.Minute
	postItemTimeout = 3 * time.Second
)

type MusicScraper interface {
	Scrape(ctx context.Context) error
}

func NewMusicScraper(bot *TelegramBot, parser RssFeedParser, lgr lgr.L, store *Store) MusicScraper {
	return &scraper{
		bot:    bot,
		parser: parser,
		lgr:    lgr,
		store:  store,
	}
}

type scraper struct {
	bot    *TelegramBot
	parser RssFeedParser
	lgr    lgr.L
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

	for _, item := range items {
		id, err := s.bot.SendImage(ctx, item)
		if err != nil {
			if bErr, ok := err.(tgbotapi.Error); ok {
				time.Sleep(time.Duration(bErr.RetryAfter) * time.Second)

				id, err = s.bot.SendImage(ctx, item)
				if err != nil {
					s.lgr.Logf("[ERROR] send image: %v", err)
					continue
				}
			} else {
				s.lgr.Logf("[ERROR] send image: %v", err)
				continue
			}
		}

		if err := s.bot.SendNews(ctx, item); err != nil {
			if bErr, ok := err.(tgbotapi.Error); ok {
				time.Sleep(time.Duration(bErr.RetryAfter) * time.Second)

				err = s.bot.SendNews(ctx, item)
				if err != nil {
					s.lgr.Logf("[ERROR] send news: %v", err)
					_, _ = s.bot.BotAPI.DeleteMessage(tgbotapi.NewDeleteMessage(s.bot.ChatId, id))
					continue
				}
			} else {
				s.lgr.Logf("[ERROR] send news: %v", err)
				_, _ = s.bot.BotAPI.DeleteMessage(tgbotapi.NewDeleteMessage(s.bot.ChatId, id))
				continue
			}
		}

		if err := s.store.SetPosted(ctx, item.Title); err != nil {
			s.lgr.Logf("[ERROR] can't set posted: item=%v, err=%v", item, err)
		}

		s.lgr.Logf("[INFO] item was send [%s]", item.Title)

		time.Sleep(postItemTimeout)
	}

	//unpublished, err := s.store.GetUnpublished(ctx)
	//if err != nil {
	//	return err
	//}
	//
	//Merge(items, unpublished)

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
