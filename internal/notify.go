package internal

import (
	"context"
	"fmt"
	"github.com/go-pkgz/lgr"
	"time"
)

const notifiesTimeout = time.Second

var platforms = []Platform{TidalPlatform, SpotifyPlatform, ItunesPlatform, YandexPlatform}

type Notifier struct {
	store *Store
	bot   *TelegramBot
	links *LinksApi
	lgr   lgr.L
}

func NewNotifier(s *Store, bot *TelegramBot, links *LinksApi) *Notifier {
	return &Notifier{
		store: s,
		bot:   bot,
		links: links,
	}
}

func (n *Notifier) Notify(ctx context.Context) error {
	items, err := n.store.GetWithNotifyFlag(ctx)
	if err != nil {
		return err
	}

	for _, item := range items {
		n.lgr.Logf("[INFO] prepare notification: title=%s", item.Title)

		releaseLink, linksByPlatform, err := n.links.GetLinks(ctx, item.Title)
		if err != nil {
			n.lgr.Logf("[ERROR] getting link: title=%s, err=%v", item.Title, err)
			continue
		}

		links, err := validatePlatforms(linksByPlatform)
		if err != nil {
			n.lgr.Logf("[ERROR] validate platforms: title=%s, err=%v", item.Title, err)
			continue
		}

		if err := n.bot.SendReleaseWithButtons(item, releaseLink, links); err != nil {
			n.lgr.Logf("[ERROR] sending: title=%s, err=%v", item.Title, err)
			continue
		}

		if err := n.store.UpdateNotifyFlag(ctx, item); err != nil {
			n.lgr.Logf("[ERROR] update notify flag: title=%s, err=%v", item.Title, err)
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
