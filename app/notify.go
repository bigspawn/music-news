package main

import (
	"context"
	"github.com/bigspawn/music-news/app/api"
	"regexp"
	"time"
)

var yearRegexp = regexp.MustCompile(`\(\d+\)`)

type Notifier struct {
	store *Store
	bot   *TelegramBot
	key   string
}

func NewNotifier(s *Store, bot *TelegramBot, songAPIKey string) *Notifier {
	return &Notifier{store: s, bot: bot, key: songAPIKey}
}

func (n *Notifier) Notify(ctx context.Context) error {
	items, err := n.store.GetWithNotifyFlag(ctx)
	if err != nil {
		return err
	}

	for _, item := range items {
		title := yearRegexp.ReplaceAllString(item.Title, "")
		Lgr.Logf("[INFO] album title = %s", title)

		resp, err := api.GetByTitle(title, n.key)
		if err != nil {
			if err == api.ErrITunesNotFound {
				Lgr.Logf("[WARN] itunes not found: %s", title)
				continue
			}
			return err
		}
		Lgr.Logf("[INFO] album links %v", resp)
		if err := n.bot.SendRelease(item, resp.PageUrl); err != nil {
			return err
		}
		if err := n.store.UpdateNotifyFlag(ctx, item); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}
	return nil
}
