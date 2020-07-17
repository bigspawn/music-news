package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

type Notifier struct {
	store *Store
	bot   *TelegramBot
	links *LinksApi
	gauge prometheus.Gauge
}

func NewNotifier(s *Store, bot *TelegramBot, links *LinksApi) *Notifier {
	return &Notifier{
		store: s,
		bot:   bot,
		links: links,
		gauge: promauto.NewGauge(prometheus.GaugeOpts{Name: "notified_news_gauge"}),
	}
}

func (n *Notifier) Notify() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := n.store.GetWithNotifyFlag(ctx)
	if err != nil {
		return err
	}

	count := 0
	for _, item := range items {
		Lgr.Logf("[INFO] prepare notification: title=%s", item.Title)

		link, err := n.links.GetLink(ctx, item.Title)
		if err != nil {
			Lgr.Logf("[ERROR] getting link: title=%s, err=%v", item.Title, err)
			continue
		}
		if err := n.bot.SendRelease(item, link); err != nil {
			Lgr.Logf("[ERROR] sending: title=%s, err=%v", item.Title, err)
			continue
		}
		if err := n.store.UpdateNotifyFlag(ctx, item); err != nil {
			Lgr.Logf("[ERROR] update notify flag: title=%s, err=%v", item.Title, err)
			continue
		}
		count++

		time.Sleep(time.Second)
	}
	n.gauge.Set(float64(count))

	return nil
}
