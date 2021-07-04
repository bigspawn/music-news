package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
	"github.com/pkg/errors"
)

type Options struct {
	User       string `long:"user" env:"USER"`
	Passwd     string `long:"passwd" env:"PASSWD"`
	ProxyURL   string `long:"proxy_url" env:"PROXY_URL"`
	Notify     bool   `short:"n" long:"notify" env:"NOTIFY"`
	BotID      string `short:"b" long:"bot_id" env:"BOT_ID" required:"true"`
	ChatID     int64  `short:"c" long:"chat_id" env:"CHAT_ID" required:"true"`
	DbURL      string `short:"d" long:"db_url" env:"DB_URL" required:"true"`
	SongAPIKey string `long:"song_api_key" env:"SONG_API_KEY"`
}

type App struct {
	store     *Store
	lgr       lgr.L
	scheduler *gocron.Scheduler
	ch        chan []News
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

	ch := make(chan []News)
	scheduler := gocron.NewScheduler(time.Now().Location())

	app := &App{
		store:     store,
		lgr:       lgr,
		scheduler: scheduler,
		ch:        ch,
	}

	if opt.Notify {
		err = runNotifier(ctx, opt, lgr, store, bot, scheduler)
	} else {

		go NewPublisher(lgr, store, bot, ch).Start(ctx)

		err = runScrapers(ctx, lgr, store, scheduler, ch)
	}
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Start() {
	a.scheduler.StartAsync()
}

func (a *App) Stop() {
	if err := a.store.db.Close(); err != nil {
		a.lgr.Logf("[ERROR] stopped application: %w", err)
	}
	a.scheduler.Stop()
	a.scheduler.Clear()
	close(a.ch)
}

func runNotifier(
	ctx context.Context,
	opt *Options,
	lgr lgr.L,
	store *Store,
	bot *TelegramBot,
	scheduler *gocron.Scheduler,
) error {
	linkApi := NewLinkApi(opt.SongAPIKey, lgr)
	notifier := NewNotifier(store, bot, linkApi, lgr)

	_, err := scheduler.Every(3).Hour().Do(func() {
		if err := notifier.Notify(ctx); err != nil {
			lgr.Logf("[ERROR] notifier %v", err)
		}

		_, next := scheduler.NextRun()
		lgr.Logf("[INFO] job next start %s", next)
	})
	return err
}

func runScrapers(
	ctx context.Context,
	lgr lgr.L,
	store *Store,
	scheduler *gocron.Scheduler,
	ch chan []News,
) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	var err error
	var itemParser ItemParser
	itemParser = NewAlterportalParser(lgr, client)
	alterportalRssFeedParser := NewRssFeedParser(AlterportalRSSFeedURL, store, lgr, itemParser)
	alterportalScr := NewMusicScraper(alterportalRssFeedParser, lgr, ch, store)
	alterportalJob := NewJob(alterportalScr, scheduler, "alterportal", lgr)
	_, err = scheduler.Every(30).Minutes().Do(alterportalJob.Do, ctx)
	if err != nil {
		return err
	}

	itemParser = NewGetRockMusicParser(lgr, client)
	getrockmusicRssFeedParser := NewRssFeedParser(GetRockMusicRss, store, lgr, itemParser)
	getrockmusicScr := NewMusicScraper(getrockmusicRssFeedParser, lgr, ch, store)
	getrockmusicJob := NewJob(getrockmusicScr, scheduler, "getrockmusic", lgr)
	_, err = scheduler.Every(32).Minutes().Do(getrockmusicJob.Do, ctx)
	if err != nil {
		return err
	}

	return nil
}
