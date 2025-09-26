package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pkgz/lgr"
)

type NotifierParams struct {
	Lgr    lgr.L
	Store  *Store
	BotAPI *RetryableBotApi
	Links  *LinksApi
}

func (p *NotifierParams) Validate() error {
	if p.Lgr == nil {
		return errors.New("lgr is nil")
	}
	if p.Store == nil {
		return errors.New("store is nil")
	}
	if p.BotAPI == nil {
		return errors.New("bot api is nil")
	}
	if p.Links == nil {
		return errors.New("links api is nil")
	}
	return nil
}

type Notifier struct {
	NotifierParams
}

func NewNotifier(params NotifierParams) (*Notifier, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &Notifier{
		NotifierParams: params,
	}, nil
}

func (n *Notifier) Notify(ctx context.Context) error {
	items, err := n.Store.GetWithNotifyFlag(ctx)
	if err != nil {
		return err
	}

	for _, item := range items {
		if err = n.notify(ctx, item); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				return err
			}

			n.Lgr.Logf("[ERROR] failed notify [%s] cause: %v", item.Title, err)

			continue
		}
		n.Lgr.Logf("[INFO] notify was send [%s]", item.Title)
	}

	return nil
}

func (n *Notifier) notify(ctx context.Context, item News) error {
	releaseLink, linksByPlatform, err := n.Links.GetLinks(ctx, item.Title)
	if err != nil {
		return fmt.Errorf("get links: %w", err)
	}

	links, err := CheckRequiredPlatforms(linksByPlatform)
	if err != nil {
		return fmt.Errorf("no required platforms found for [%s]: %w", item.Title, err)
	}

	if len(links) == 0 && releaseLink == "" {
		return fmt.Errorf("no links available for [%s]", item.Title)
	}

	err = n.BotAPI.SendReleaseNews(ctx, ReleaseNews{
		News:          item,
		ReleaseLink:   releaseLink,
		PlatformLinks: links,
	})
	if err != nil {
		return fmt.Errorf("send release news: %w", err)
	}

	if err := n.Store.UpdateNotifyFlag(ctx, item); err != nil {
		return fmt.Errorf("update notify flag: %w", err)
	}

	return nil
}
