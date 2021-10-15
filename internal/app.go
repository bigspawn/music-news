package internal

import (
	"context"
	"net"
	"net/http"
	"net/url"
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
	bot       *TelegramBot
}

func NewApp(ctx context.Context, opt *Options, lgr lgr.L) (*App, error) {
	bot, err := NewTelegramBotAPI(opt, lgr)
	if err != nil {
		return nil, errors.Wrap(err, "init telegram BotAPI api")
	}

	store, err := NewNewsStore(opt.DbURL, lgr)
	if err != nil {
		return nil, errors.Wrap(err, "init Store")
	}

	app := &App{
		store:     store,
		bot:       bot,
		lgr:       lgr,
		scheduler: gocron.NewScheduler(time.Now().Location()),
		ch:        make(chan []News),
	}

	if opt.Notify {
		if err := app.runNotifier(ctx, opt); err != nil {
			return nil, err
		}

		return app, nil
	}

	publisher := &Publisher{
		Lgr:    lgr,
		NewsCh: app.ch,
		BotAPI: bot,
		Store:  store,
	}

	go publisher.Start(ctx)

	if err := app.runScrapers(ctx, opt); err != nil {
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

func (a *App) runNotifier(ctx context.Context, opt *Options) error {
	notifier := &Notifier{
		Store:  a.store,
		BotAPI: a.bot,
		Links: &LinksApi{
			Client: http.DefaultClient,
			Key:    opt.SongAPIKey,
			Lgr:    a.lgr,
		},
		Lgr: a.lgr,
	}

	jobFun := func() {
		if err := notifier.Notify(ctx); err != nil {
			a.lgr.Logf("[ERROR] notifier %v", err)
		}

		_, next := a.scheduler.NextRun()
		a.lgr.Logf("[INFO] job next start %s", next)
	}

	_, err := a.scheduler.Every(3).Hour().Do(jobFun)
	return err
}

func (a *App) runScrapers(ctx context.Context, opt *Options) error {
	//--- AlterPortal
	link, err := url.Parse(AlterportalRSSFeedURL)
	if err != nil {
		return err
	}

	job := Job{
		s: &Scraper{
			parser: &Parser{
				url:        AlterportalRSSFeedURL,
				feedParser: gofeed.NewParser(),
				store:      a.store,
				lgr:        a.lgr,
				itemParser: &AlterPortalParser{
					Lgr:    a.lgr,
					Client: http.DefaultClient,
				},
				siteLabel: link.Host,
			},
			lgr:       a.lgr,
			ch:        a.ch,
			store:     a.store,
			withDelay: false,
		},
		sch:  a.scheduler,
		name: "alterPortal",
		lgr:  a.lgr,
	}

	_, err = a.scheduler.Every(15).Minutes().Do(job.Do, ctx)
	if err != nil {
		return err
	}

	//--- GetRockMusic
	dialer, err := newDialer(opt)
	if err != nil {
		return err
	}

	client := newClient(dialer)

	feedParser := gofeed.NewParser()
	feedParser.Client = client

	link, err = url.Parse(GetRockMusicRss)
	if err != nil {
		return err
	}

	job = Job{
		s: &Scraper{
			parser: &Parser{
				url:        GetRockMusicRss,
				feedParser: gofeed.NewParser(),
				store:      a.store,
				lgr:        a.lgr,
				itemParser: &GetRockMusicParser{
					Lgr:    a.lgr,
					Client: client,
				},
				siteLabel: link.Host,
				withDelay: true,
			},
			lgr:       a.lgr,
			ch:        a.ch,
			store:     a.store,
			withDelay: true,
		},
		sch:  a.scheduler,
		name: "getRockMusic",
		lgr:  a.lgr,
	}

	_, err = a.scheduler.Every(35).Minutes().Do(job.Do, ctx)
	if err != nil {
		return err
	}

	return nil
}

func newDialer(opt *Options) (proxy.Dialer, error) {
	defaultDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	if opt.ProxyURL == "" {
		return defaultDialer, nil
	}

	dialer, err := proxy.SOCKS5("tcp", opt.ProxyURL, &proxy.Auth{
		User:     opt.User,
		Password: opt.Passwd,
	}, defaultDialer)
	if err != nil {
		return nil, err
	}

	return dialer, nil
}

func newClient(dialer proxy.Dialer) *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
				return dialer.Dial(network, addr)
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}
