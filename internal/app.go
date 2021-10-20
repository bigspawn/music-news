package internal

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/proxy"
	tb "gopkg.in/tucnak/telebot.v2"
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
	ProxyType  string `long:"proxy_type" env:"PROXY_TYPE"`
}

type ProxyType string

const (
	ProxyTypeNone   ProxyType = ""
	ProxyTypeSocks5 ProxyType = "SOCKS5"
	ProxyTypeHttp   ProxyType = "HTTP"
)

type App struct {
	store     *Store
	lgr       lgr.L
	scheduler *gocron.Scheduler
	ch        chan []News
	bot       *RetryableBotApi
}

func NewApp(ctx context.Context, opt *Options, lgr lgr.L) (*App, error) {
	db, err := sql.Open(driver, opt.DbURL)
	if err != nil {
		return nil, err
	}

	b, err := tb.NewBot(tb.Settings{
		Token:     opt.BotID,
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tb.ModeHTML,
	})
	if err != nil {
		return nil, err
	}

	app := &App{
		store: &Store{
			db:  db,
			lgr: lgr,
		},
		bot: &RetryableBotApi{
			Bot: BotAPI{
				Bot:     b,
				ChantID: tb.ChatID(opt.ChatID),
			},
			Lgr: lgr,
		},
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
		BotAPI: app.bot,
		Store:  app.store,
	}

	go publisher.Start(ctx)

	go processUnpublished(ctx, app, lgr)

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
	if err := runCoreRadio(ctx, a); err != nil {
		return err
	}

	if err := runGetRockMusic(ctx, a, opt); err != nil {
		return err
	}

	if err := runAlterPortal(ctx, a); err != nil {
		return err
	}

	return nil
}

func runAlterPortal(ctx context.Context, a *App) error {
	link, err := url.Parse(AlterPortalParserRssURL)
	if err != nil {
		return err
	}

	job := Job{
		s: &Scraper{
			parser: &Parser{
				url:        AlterPortalParserRssURL,
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
			name:      "alterPortal",
		},
		sch:  a.scheduler,
		name: "alterPortal",
		lgr:  a.lgr,
	}

	_, err = a.scheduler.Every(15).Minutes().Do(job.Do, ctx)
	if err != nil {
		return err
	}

	return nil
}

func runGetRockMusic(ctx context.Context, a *App, opt *Options) error {
	dialer, err := newDialer(opt)
	if err != nil {
		return err
	}

	client, err := newClient(dialer, opt)
	if err != nil {
		return err
	}

	feedParser := gofeed.NewParser()
	feedParser.Client = client

	link, err := url.Parse(GetRockMusicParserRssURL)
	if err != nil {
		return err
	}

	job := Job{
		s: &Scraper{
			parser: &Parser{
				url:        GetRockMusicParserRssURL,
				feedParser: feedParser,
				store:      a.store,
				lgr:        a.lgr,
				itemParser: &GetRockMusicParser{
					Lgr:    a.lgr,
					Client: client,
				},
				siteLabel: link.Host,
				withDelay: false,
			},
			lgr:       a.lgr,
			ch:        a.ch,
			store:     a.store,
			withDelay: false,
			name:      "getRockMusic",
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

func runCoreRadio(ctx context.Context, a *App) error {
	link, err := url.Parse(CoreRadioParserRssURL)
	if err != nil {
		return err
	}

	job := Job{
		s: &Scraper{
			parser: &Parser{
				url:        CoreRadioParserRssURL,
				feedParser: gofeed.NewParser(),
				store:      a.store,
				lgr:        a.lgr,
				itemParser: &CoreRadioParser{
					Lgr:    a.lgr,
					Client: http.DefaultClient,
				},
				siteLabel: link.Host,
			},
			lgr:       a.lgr,
			ch:        a.ch,
			store:     a.store,
			withDelay: false,
			name:      "coreRadio",
		},
		sch:  a.scheduler,
		name: "coreRadio",
		lgr:  a.lgr,
	}

	_, err = a.scheduler.Every(20).Minutes().Do(job.Do, ctx)
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

	if ProxyType(opt.ProxyType) == ProxyTypeSocks5 {
		dialer, err := proxy.SOCKS5("tcp", opt.ProxyURL, &proxy.Auth{
			User:     opt.User,
			Password: opt.Passwd,
		}, defaultDialer)
		if err != nil {
			return nil, err
		}

		return dialer, nil
	}

	return defaultDialer, nil

}

func newClient(dialer proxy.Dialer, opt *Options) (*http.Client, error) {
	httpProxy := http.ProxyFromEnvironment

	if ProxyType(opt.ProxyType) == ProxyTypeHttp {
		proxyURL, err := url.Parse(opt.ProxyURL)
		if err != nil {
			return nil, err
		}

		httpProxy = http.ProxyURL(proxyURL)
	}

	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Proxy: httpProxy,
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
				return dialer.Dial(network, addr)
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}, nil
}

func processUnpublished(ctx context.Context, app *App, lgr lgr.L) {
	items, err := app.store.GetUnpublished(ctx)
	if err != nil {
		lgr.Logf("[ERROR] GetUnpublished: %v", err)
	} else {
		arr := make([]News, 0, len(items))
		for _, v := range items {
			arr = append(arr, v)
		}
		app.ch <- arr
	}
}
