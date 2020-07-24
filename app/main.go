package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pkgz/lgr"
	"github.com/jasonlvhit/gocron"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"github.com/mmcdole/gofeed"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	period       = 30 // minutes
	notifyPeriod = 3  // hours

	timeoutBetween = 3 * time.Second
	parsingTimeout = 5 * time.Minute
	notifyTimeout  = 5 * time.Minute
)

var Lgr = lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc)

type Options struct {
	User       string `long:"user" env:"USER"`
	Passwd     string `long:"passwd" env:"PASSWD"`
	ProxyURL   string `long:"proxy_url" env:"PROXY_URL"`
	FeedURL    string `short:"f" long:"feed_url" env:"FEED_URL"`
	Notify     bool   `short:"n" long:"notify" env:"NOTIFY"`
	BotID      string `short:"b" long:"bot_id" env:"BOT_ID" required:"true"`
	ChatID     int64  `short:"c" long:"chat_id" env:"CHAT_ID" required:"true"`
	DbURL      string `short:"d" long:"db_url" env:"DB_URL" required:"true"`
	SongAPIKey string `long:"song_api_key" env:"SONG_API_KEY"`
}

func main() {
	opt := &Options{}
	p := flags.NewParser(opt, flags.Default)
	if _, err := p.Parse(); err != nil {
		Lgr.Logf("[FATAL] parse flags %v", err)
	}

	Lgr.Logf("[INFO] %v", opt)

	bot, err := NewTelegramBotAPI(opt)
	if err != nil {
		Lgr.Logf("[FATAL] init bot %v", err)
	}

	store, err := NewNewsStore(opt.DbURL)
	if err != nil {
		Lgr.Logf("[FATAL] init store %v", err)
	}
	defer func() {
		_ = store.conn.Close()
	}()

	go metrics()

	if opt.Notify {
		notifierRun(store, bot, opt.SongAPIKey)
	}

	parserRun(store, bot, opt)
}

func notifierRun(store *Store, bot *TelegramBot, songAPIKey string) {
	links := NewLinkApi(songAPIKey)
	notifier := NewNotifier(store, bot, links)

	gocron.Every(notifyPeriod).Hours().Do(doNotify, notifier)
	gocron.RunAll()

	_, next := gocron.NextRun()
	Lgr.Logf("[INFO] Next start %s", next)

	<-gocron.Start()
}

func parserRun(store *Store, bot *TelegramBot, opt *Options) {
	parser := NewParser(gofeed.NewParser(), store, opt.FeedURL)

	gocron.Every(period).Minutes().Do(work, bot, parser)
	gocron.RunAll()

	_, next := gocron.NextRun()
	Lgr.Logf("[INFO] Next start %s", next)

	<-gocron.Start()
}

func work(b *TelegramBot, p *SiteParser) {
	ctx, cancel := context.WithTimeout(context.Background(), parsingTimeout)
	defer cancel()

	items, err := p.Parse(ctx)
	if err != nil {
		Lgr.Logf("[ERROR] can't parse: err=%v", err)
		return
	}

	items, err = p.MergeWithUnpublished(ctx, items)
	if err != nil {
		Lgr.Logf("[ERROR] can't merge with unpublished: err=%v", err)
		return
	}

	count := 0
	for _, item := range items {
		time.Sleep(timeoutBetween)

		if err := b.SendImage(ctx, item); err != nil {
			Lgr.Logf("[ERROR] send image: %v", err)
			continue
		}
		if err := b.SendNews(ctx, item); err != nil {
			Lgr.Logf("[ERROR] send news: %v", err)
			continue
		}
		if err := p.SetPosted(ctx, item); err != nil {
			Lgr.Logf("[ERROR] can't set posted: item=%v, err=%v", item, err)
		}

		Lgr.Logf("[INFO] Item was send [%s]", item.Title)

		count++
	}

	p.SentGauge.Set(float64(count))
}

func doNotify(n *Notifier) {
	if err := n.Notify(); err != nil {
		Lgr.Logf("[ERROR] notifier %v", err)
	}
}

func metrics() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":9091", nil); err != nil {
		Lgr.Logf("[ERROR] metrics handler: err=%v", err)
	}
}
