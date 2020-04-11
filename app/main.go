package main

import (
	"context"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/go-pkgz/lgr"
	"github.com/jasonlvhit/gocron"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
)

const (
	period         = 30 // minutes
	timeoutBetween = 10 * time.Second
	notifyPeriod   = 3 // hours
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

	if opt.Notify {
		notifierRun(store, bot, opt.SongAPIKey)
	}

	parserRun(store, bot, opt)
}

func notifierRun(store *Store, bot *TelegramBot, songAPIKey string) {
	notifier := NewNotifier(store, bot, songAPIKey)

	gocron.Every(notifyPeriod).Hours().Do(doNotify, notifier)
	gocron.RunAll()

	_, next := gocron.NextRun()
	Lgr.Logf("[INFO] Next start %s", next)

	<-gocron.Start()
}

func parserRun(store *Store, bot *TelegramBot, opt *Options) {
	parser := &SiteParser{
		FeedParser: gofeed.NewParser(),
		Store:      store,
		URL:        opt.FeedURL,
	}

	gocron.Every(period).Minutes().Do(work, bot, parser)
	gocron.RunAll()

	_, next := gocron.NextRun()
	Lgr.Logf("[INFO] Next start %s", next)

	<-gocron.Start()
}

func work(b *TelegramBot, p *SiteParser) {
	items, err := p.Parse()
	if err != nil {
		Lgr.Logf("[ERROR] Error %s. Waiting for next execution", err.Error())
		return
	}

	for _, item := range items {
		time.Sleep(timeoutBetween)
		if err := b.SendImage(item); err != nil {
			continue
		}
		_ = b.SendNews(item)
		Lgr.Logf("[INFO] Item was send [%s]", item.Title)
	}
}

func doNotify(n *Notifier) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := n.Notify(ctx); err != nil {
		Lgr.Logf("[ERROR] notifier %v", err)
	}
}
