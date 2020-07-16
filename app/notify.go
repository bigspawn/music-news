package main

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/bigspawn/music-news/app/api"
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
		title := strings.TrimSpace(yearRegexp.ReplaceAllString(item.Title, ""))
		Lgr.Logf("[INFO] album title = %s", title)

		resp, err := api.GetByTitle(title, n.key)
		if err != nil {
			Lgr.Logf("[ERROR] getting info %s: %s", err.Error(), title)
			continue
		}
		Lgr.Logf("[INFO] album links %v", resp)
		if err := n.bot.SendRelease(item, resp.PageUrl); err != nil {
			Lgr.Logf("[ERROR] sending %s: %s", err.Error(), title)
			continue
		}
		if err := n.store.UpdateNotifyFlag(ctx, item); err != nil {
			Lgr.Logf("[ERROR] update flag %s: %s", err.Error(), title)
			continue
		}

		time.Sleep(time.Second)
	}
	return nil
}
