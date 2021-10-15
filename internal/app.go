package internal

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"golang.org/x/net/proxy"
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

		err = runScrapers(ctx, lgr, store, scheduler, ch, opt)
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

func runScrapers(ctx context.Context, lgr lgr.L, store *Store, scheduler *gocron.Scheduler,
	ch chan []News, opt *Options) error {

	d, err := proxy.SOCKS5("tcp", opt.ProxyURL, &proxy.Auth{
		User:     opt.User,
		Password: opt.Passwd,
	}, &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	})
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
				return d.Dial(network, addr)
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	feedParser := gofeed.NewParser()
	feedParser.Client = client

	var itemParser ItemParser
	itemParser = NewAlterportalParser(lgr)
	alterportalRssFeedParser := NewRssFeedParser(AlterportalRSSFeedURL, store, lgr, itemParser, feedParser)
	alterportalScr := NewMusicScraper(alterportalRssFeedParser, lgr, ch, store, false)
	alterportalJob := NewJob(alterportalScr, scheduler, "alterportal", lgr)
	_, err = scheduler.Every(15).Minutes().Do(alterportalJob.Do, ctx)
	if err != nil {
		return err
	}

	itemParser = NewGetRockMusicParser(lgr, client)
	getrockmusicRssFeedParser := NewRssFeedParser(GetRockMusicRss, store, lgr, itemParser, feedParser)
	getrockmusicScr := NewMusicScraper(getrockmusicRssFeedParser, lgr, ch, store, true)
	getrockmusicJob := NewJob(getrockmusicScr, scheduler, "getrockmusic", lgr)
	_, err = scheduler.Every(35).Minutes().Do(getrockmusicJob.Do, ctx)
	if err != nil {
		return err
	}

	return nil
}
