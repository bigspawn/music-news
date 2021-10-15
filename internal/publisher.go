package internal

import (
	"context"
	"time"

	"github.com/go-pkgz/lgr"
	tbapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Publisher struct {
	Lgr    lgr.L
	NewsCh <-chan []News
	BotAPI *TelegramBot
	Store  *Store
}

func (p *Publisher) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case items := <-p.NewsCh:
			err := p.publish(ctx, items)
			if err != nil {
				p.Lgr.Logf("[ERROR] publishing %v", err)
			}
		}
	}
}

func (p *Publisher) publish(ctx context.Context, items []News) error {
	for _, item := range items {
		id, err := p.BotAPI.SendImage(ctx, item)
		if err != nil {
			if bErr, ok := err.(tbapi.Error); ok {
				time.Sleep(time.Duration(bErr.RetryAfter) * time.Second)

				id, err = p.BotAPI.SendImage(ctx, item)
				if err != nil {
					p.Lgr.Logf("[ERROR] send image: %v", err)
					continue
				}
			} else {
				p.Lgr.Logf("[ERROR] send image: %v", err)
				continue
			}
		}

		if err := p.BotAPI.SendNews(ctx, item); err != nil {
			if bErr, ok := err.(tbapi.Error); ok {
				time.Sleep(time.Duration(bErr.RetryAfter) * time.Second)

				err = p.BotAPI.SendNews(ctx, item)
				if err != nil {
					p.Lgr.Logf("[ERROR] send news: %v", err)
					_, _ = p.BotAPI.BotAPI.DeleteMessage(tbapi.NewDeleteMessage(p.BotAPI.ChatId, id))
					continue
				}
			} else {
				p.Lgr.Logf("[ERROR] send news: %v", err)
				_, _ = p.BotAPI.BotAPI.DeleteMessage(tbapi.NewDeleteMessage(p.BotAPI.ChatId, id))
				continue
			}
		}

		if err := p.Store.SetPosted(ctx, item.Title); err != nil {
			p.Lgr.Logf("[ERROR] can't set posted: item=%v, err=%v", item, err)
		}

		p.Lgr.Logf("[INFO] item was send [%s]", item.Title)

		time.Sleep(postItemTimeout)
	}
	return nil
}
