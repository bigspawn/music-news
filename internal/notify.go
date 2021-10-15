package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pkgz/lgr"
	tbapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const notifiesTimeout = time.Second

var platforms = []Platform{TidalPlatform, SpotifyPlatform, ItunesPlatform, YandexPlatform}

type Notifier struct {
	Store  *Store
	BotAPI *TelegramBot
	Links  *LinksApi
	Lgr    lgr.L
}

func (n *Notifier) Notify(ctx context.Context) error {
	items, err := n.Store.GetWithNotifyFlag(ctx)
	if err != nil {
		return err
	}

	for _, item := range items {
		n.Lgr.Logf("[INFO] prepare notification: title=%s", item.Title)

		releaseLink, linksByPlatform, err := n.Links.GetLinks(ctx, item.Title)
		if err != nil {
			n.Lgr.Logf("[ERROR] getting link: title=%s, err=%v", item.Title, err)
			continue
		}

		n.Lgr.Logf("[INFO] platforms=%s", linksByPlatform)

		links, err := validatePlatforms(linksByPlatform)
		if err != nil {
			n.Lgr.Logf("[ERROR] validate platforms: title=%s, err=%v", item.Title, err)
			continue
		}

		if err := n.BotAPI.SendReleaseWithButtons(item, releaseLink, links); err != nil {
			if bErr, ok := err.(tbapi.Error); ok {
				time.Sleep(time.Duration(bErr.RetryAfter) * time.Second)

				err = n.BotAPI.SendReleaseWithButtons(item, releaseLink, links)
				if err != nil {
					n.Lgr.Logf("[ERROR] sending: title=%s, err=%v", item.Title, err)

					continue
				}
			} else {
				n.Lgr.Logf("[ERROR] sending: title=%s, err=%v", item.Title, err)
				continue
			}
		}

		if err := n.Store.UpdateNotifyFlag(ctx, item); err != nil {
			n.Lgr.Logf("[ERROR] update notify flag: title=%s, err=%v", item.Title, err)
			continue
		}

		time.Sleep(notifiesTimeout)
	}

	return nil
}

func validatePlatforms(byPlatform map[Platform]string) (map[Platform]string, error) {
	links := make(map[Platform]string)
	for _, p := range platforms {
		if l, ok := byPlatform[p]; ok {
			links[p] = l
			continue
		}
		if p == TidalPlatform || p == SpotifyPlatform || p == ItunesPlatform {
			return nil, fmt.Errorf("link for platform=%s not found", p)
		}
	}
	return links, nil
}
