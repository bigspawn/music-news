package internal

import (
	"context"
	"fmt"
	"time"

	goOdesli "github.com/bigspawn/go-odesli"
	"github.com/go-pkgz/lgr"
	"github.com/pkg/errors"
)

var platforms = []goOdesli.Platform{
	goOdesli.PlatformTidal,
	goOdesli.PlatformSpotify,
	goOdesli.PlatformItunes,
	goOdesli.PlatformYandex,
}

type Notifier struct {
	Store  *Store
	BotAPI *RetryableBotApi
	Links  *LinksApi
	Lgr    lgr.L
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

		duration := time.Duration(RandBetween(10_000, 1)) * time.Millisecond

		n.Lgr.Logf("[INFO] sleep between next notify [%s]", duration)

		WaitUntil(ctx, duration)
	}

	return nil
}

func (n *Notifier) notify(ctx context.Context, item News) error {
	releaseLink, linksByPlatform, err := n.Links.GetLinks(ctx, item.Title)
	if err != nil {
		return errors.Wrap(err, "get platform links")
	}

	links, err := validatePlatforms(linksByPlatform)
	if err != nil {
		return errors.Wrap(err, "validate platform links")
	}

	err = n.BotAPI.SendReleaseNews(ctx, ReleaseNews{
		News:          item,
		ReleaseLink:   releaseLink,
		PlatformLinks: links,
	})
	if err != nil {
		return errors.Wrap(err, "send notify")
	}

	if err := n.Store.UpdateNotifyFlag(ctx, item); err != nil {
		return errors.Wrap(err, "update store notify")
	}

	return nil
}

func validatePlatforms(byPlatform map[goOdesli.Platform]string) (map[goOdesli.Platform]string, error) {
	links := make(map[goOdesli.Platform]string)
	for _, p := range platforms {
		if l, ok := byPlatform[p]; ok {
			links[p] = l
			continue
		}
		if p == goOdesli.PlatformTidal || p == goOdesli.PlatformSpotify || p == goOdesli.PlatformItunes {
			return nil, fmt.Errorf("link for platform=%s not found", p)
		}
	}
	return links, nil
}
