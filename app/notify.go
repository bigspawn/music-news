package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
	gauge prometheus.Gauge
}

func NewNotifier(s *Store, bot *TelegramBot, songAPIKey string) *Notifier {
	gauge := promauto.NewGauge(prometheus.GaugeOpts{Name: "notified_news_gauge"})
	return &Notifier{store: s, bot: bot, key: songAPIKey, gauge: gauge}
}

func (n *Notifier) Notify(ctx context.Context) error {
	items, err := n.store.GetWithNotifyFlag(ctx)
	if err != nil {
		return err
	}

	count := 0
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
	n.gauge.Set(float64(count))

	return nil
}
