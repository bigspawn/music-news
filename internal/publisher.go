package internal

import (
	"context"
	"errors"
	"time"

	"github.com/go-pkgz/lgr"
	tb "gopkg.in/telebot.v3"
)

type PublisherParams struct {
	Lgr    lgr.L
	NewsCh <-chan []News
	BotAPI *RetryableBotApi
	Store  *Store
}

func (p *PublisherParams) Validate() error {
	if p.Lgr == nil {
		return errors.New("lgr is required")
	}
	if p.NewsCh == nil {
		return errors.New("news channel is required")
	}
	if p.BotAPI == nil {
		return errors.New("bot api is required")
	}
	if p.Store == nil {
		return errors.New("store is required")
	}
	return nil
}

type Publisher struct {
	PublisherParams
}

func NewPublisher(params PublisherParams) (*Publisher, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &Publisher{PublisherParams: params}, nil
}

func (p *Publisher) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case items := <-p.NewsCh:
			err := p.publish(ctx, items)
			if err != nil {
				p.Lgr.Logf("[ERROR] publishing %v", err)
			}
		}
	}
}

func (p *Publisher) Stop() {}

func (p *Publisher) publish(ctx context.Context, items []News) error {
	for _, item := range items {
		if err := p.BotAPI.SendNews(ctx, item); err != nil {
			p.Lgr.Logf("[ERROR] can't send news: item=%v, err=%v", item, err)

			var bErr *tb.Error
			if errors.As(err, &bErr) && bErr.Code == 400 && bErr.Message == "wrong type of the web page content" {
				_ = p.Store.SetPostedAndNotified(ctx, item.ID)
			}

			continue
		}

		if err := p.Store.SetPostedByID(ctx, item.ID); err != nil {
			p.Lgr.Logf("[ERROR] can't set posted: item=%v, err=%v", item, err)
			continue
		}

		p.Lgr.Logf("[INFO] news was send [%s]", item.Title)

		duration := time.Duration(RandBetween(10_000, 1)) * time.Millisecond

		p.Lgr.Logf("[INFO] sleep between next send [%s]", duration)

		WaitUntil(ctx, duration)
	}

	return nil
}
