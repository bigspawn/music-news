package internal

import (
	"context"
	"time"

	"github.com/go-pkgz/lgr"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Publisher struct {
	lgr   lgr.L
	ch    <-chan []News
	bot   *TelegramBot
	store *Store
}

func NewPublisher(l lgr.L, s *Store, b *TelegramBot, ch <-chan []News) *Publisher {
	return &Publisher{
		lgr:   l,
		ch:    ch,
		bot:   b,
		store: s,
	}
}

func (p *Publisher) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case items := <-p.ch:
			err := p.publish(ctx, items)
			if err != nil {
				p.lgr.Logf("[ERROR] publishing %v", err)
			}
		}
	}
}

func (p *Publisher) publish(ctx context.Context, items []News) error {
	for _, item := range items {
		if err := p.store.Insert(ctx, item); err != nil {
			p.lgr.Logf("[ERROR] insert to db: %v", err)
			continue
		}
	}

	for _, item := range items {
		id, err := p.bot.SendImage(ctx, item)
		if err != nil {
			if bErr, ok := err.(tgbotapi.Error); ok {
				time.Sleep(time.Duration(bErr.RetryAfter) * time.Second)

				id, err = p.bot.SendImage(ctx, item)
				if err != nil {
					p.lgr.Logf("[ERROR] send image: %v", err)
					continue
				}
			} else {
				p.lgr.Logf("[ERROR] send image: %v", err)
				continue
			}
		}

		if err := p.bot.SendNews(ctx, item); err != nil {
			if bErr, ok := err.(tgbotapi.Error); ok {
				time.Sleep(time.Duration(bErr.RetryAfter) * time.Second)

				err = p.bot.SendNews(ctx, item)
				if err != nil {
					p.lgr.Logf("[ERROR] send news: %v", err)
					_, _ = p.bot.BotAPI.DeleteMessage(tgbotapi.NewDeleteMessage(p.bot.ChatId, id))
					continue
				}
			} else {
				p.lgr.Logf("[ERROR] send news: %v", err)
				_, _ = p.bot.BotAPI.DeleteMessage(tgbotapi.NewDeleteMessage(p.bot.ChatId, id))
				continue
			}
		}

		if err := p.store.SetPosted(ctx, item.Title); err != nil {
			p.lgr.Logf("[ERROR] can't set posted: item=%v, err=%v", item, err)
		}

		p.lgr.Logf("[INFO] item was send [%s]", item.Title)

		time.Sleep(postItemTimeout)
	}
	return nil
}
