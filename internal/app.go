package internal

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
	"github.com/pkg/errors"
	"time"
)

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

func NewApp(ctx context.Context, opt *Options, lgr lgr.L) (*App, error) {
	bot, err := NewTelegramBotAPI(opt, lgr)
	if err != nil {
		return nil, errors.Wrap(err, "init telegram bot api")
	}

	store, err := NewNewsStore(opt.DbURL, lgr)
	if err != nil {
		return nil, errors.Wrap(err, "init store")
	}

	scheduler := gocron.NewScheduler(time.Now().Location())

	if opt.Notify {
		linkApi := NewLinkApi(opt.SongAPIKey, lgr)
		n := NewNotifier(store, bot, linkApi)
		task := func() {
			if err := n.Notify(ctx); err != nil {
				lgr.Logf("[ERROR] notifier %v", err)
			}

			_, next := scheduler.NextRun()
			lgr.Logf("[INFO] job next start %s", next)
		}
		scheduler.Every(3).Hour().Do(task)
	} else {
		itemParser := NewAlterportalParser(lgr)
		rssParser := NewRssFeedParser(opt.FeedURL, store, lgr, itemParser)
		s := NewMusicScraper(bot, rssParser, lgr, store)
		task := func() {
			if err := s.Scrape(ctx); err != nil {
				lgr.Logf("[ERROR] scraper %v", err)
			}

			_, next := scheduler.NextRun()
			lgr.Logf("[INFO] job next start %s", next)
		}
		scheduler.Every(10).Minutes().Do(task)
	}

	return &App{
		store:     store,
		lgr:       lgr,
		scheduler: scheduler,
	}, nil
}

type App struct {
	store     *Store
	lgr       lgr.L
	scheduler *gocron.Scheduler
}

func (a App) Start() {
	a.scheduler.StartAsync()

}

func (a App) Stop() {
	if err := a.store.db.Close(); err != nil {
		a.lgr.Logf("[ERROR] stopped application: %w", err)
	}
	a.scheduler.Stop()
	a.scheduler.Clear()
}
