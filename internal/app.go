package internal

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"time"

	itunes "github.com/bigspawn/go-itunes-api"
	"github.com/bigspawn/go-odesli"
	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	tb "gopkg.in/telebot.v3"
)

type App struct {
	lgr           lgr.L
	store         *Store
	scheduler     *gocron.Scheduler
	ch            chan []News
	reNotifiedJob *ReNotifiedJob
}

func NewApp(ctx context.Context, opt *Options, lgr lgr.L) (*App, error) {
	db, err := sql.Open(driver, opt.DBUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	itunesAPI, err := itunes.NewClient(itunes.ClientOption{})
	if err != nil {
		return nil, fmt.Errorf("failed to create itunes api client: %w", err)
	}

	odesliAPI, err := odesli.NewClient(odesli.ClientOption{
		APIToken: opt.SongApiKey,
		Debug:    false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create odesli api client: %w", err)
	}

	links, err := NewLinksApi(LinksApiParams{
		Lgr:          lgr,
		ITunesClient: itunesAPI,
		OdesliClient: odesliAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create links api: %w", err)
	}

	store, err := NewStore(StoreParams{
		Lgr: lgr,
		DB:  db,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:     opt.BotID,
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tb.ModeHTML,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	scheduler := gocron.NewScheduler(time.UTC)

	_, err = createNotifier(ctx, lgr, opt, bot, store, links, scheduler)
	if err != nil {
		return nil, fmt.Errorf("failed to create notifier: %w", err)
	}

	var (
		ch            = make(chan []News)
		httpClient    = NewHttpClient(NewDialer())
		reNotifiedJob *ReNotifiedJob
	)

	if !opt.OnlyNotifier {
		err = runCoreRadio(ctx, lgr, store, httpClient, ch, scheduler)
		if err != nil {
			return nil, fmt.Errorf("failed to run core radio scraper: %w", err)
		}

		err = runAlterPortal(ctx, lgr, store, httpClient, ch, scheduler)
		if err != nil {
			return nil, fmt.Errorf("failed to run alter portal scraper: %w", err)
		}

		// cloudflare: 403 Forbidden
		// err = runGetRockMusic(ctx, lgr, store, httpClient, ch, scheduler)
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to run get rock music scraper: %w", err)
		// }

		_, err = createPublisher(ctx, lgr, store, bot, opt, ch)
		if err != nil {
			return nil, fmt.Errorf("failed to create publisher: %w", err)
		}

		reNotifiedJob = NewReNotifiedJob(ReNotifiedJobParams{
			lgr:   lgr,
			store: store,
			ch:    ch,
		})
	}

	return &App{
		lgr:           lgr,
		store:         store,
		scheduler:     scheduler,
		ch:            ch,
		reNotifiedJob: reNotifiedJob,
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	go a.scheduler.StartAsync()
	if a.reNotifiedJob != nil {
		go a.reNotifiedJob.Run(ctx)
	}
	return nil
}

func (a *App) Stop() {
	if a.reNotifiedJob != nil {
		a.reNotifiedJob.Stop()
	}
	a.scheduler.Stop()
	a.scheduler.Clear()
	a.store.Stop()
	close(a.ch)
}

func createNotifier(
	ctx context.Context,
	lgr lgr.L,
	opt *Options,
	bot *tb.Bot,
	store *Store,
	links *LinksApi,
	scheduler *gocron.Scheduler,
) (*Notifier, error) {
	notifyBot, err := NewBotAPI(BotAPIParams{
		Bot:     bot,
		ChantID: tb.ChatID(opt.NotifierChatID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create bot api: %w", err)
	}

	notifyRetryBot, err := NewRetryableBotApi(RetryableBotApiParams{
		Lgr: lgr,
		Bot: notifyBot,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create retryable bot: %w", err)
	}

	notifier, err := NewNotifier(NotifierParams{
		Lgr:    lgr,
		Store:  store,
		BotAPI: notifyRetryBot,
		Links:  links,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create notifier: %w", err)
	}

	_, err = scheduler.
		Every(1).
		Hour().
		DoWithJobDetails(func(job gocron.Job) {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						lgr.Logf("[ERROR] PANIC recovered in notifier: %v", r)
					}
				}()
				if nErr := notifier.Notify(ctx); nErr != nil {
					lgr.Logf("[ERROR] notifier %v", nErr)
				}
				lgr.Logf("[INFO] notifier done, next run %v", job.NextRun())
			}()
		})
	if err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}

	return notifier, nil
}

func runCoreRadio(
	ctx context.Context,
	lgr lgr.L,
	store *Store,
	httpClient *http.Client,
	ch chan []News,
	scheduler *gocron.Scheduler,
) error {
	link, err := url.Parse(CoreRadioParserRssURL)
	if err != nil {
		return err
	}
	s := &Scraper{
		parser: &Parser{
			url:        CoreRadioParserRssURL,
			feedParser: gofeed.NewParser(),
			store:      store,
			lgr:        lgr,
			itemParser: &CoreRadioParser{
				Lgr:    lgr,
				Client: httpClient,
			},
			siteLabel: link.Host,
			withDelay: false,
		},
		lgr:       lgr,
		ch:        ch,
		store:     store,
		withDelay: false,
		name:      "coreRadio",
	}
	_, err = scheduler.
		Every(16).
		Minutes().
		DoWithJobDetails(func(job gocron.Job) {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						lgr.Logf("[ERROR] PANIC recovered in CoreRadio scraper: %v", r)
					}
				}()
				sErr := s.Scrape(ctx)
				if sErr != nil {
					lgr.Logf("[ERROR] failed to scrape %s: %v", s.name, sErr)
				}
				lgr.Logf("[INFO] CoreRadioParser done, next run %v", job.NextRun())
			}()
		})
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}
	return nil
}

func runGetRockMusic(
	ctx context.Context,
	lgr lgr.L,
	store *Store,
	httpClient *http.Client,
	ch chan []News,
	scheduler *gocron.Scheduler,
) error {
	link, err := url.Parse(GetRockMusicParserRssURL)
	if err != nil {
		return err
	}
	parser := gofeed.NewParser()
	parser.Client = httpClient
	s := &Scraper{
		parser: &Parser{
			url:        GetRockMusicParserRssURL,
			feedParser: parser,
			store:      store,
			lgr:        lgr,
			itemParser: &GetRockMusicParser{
				Lgr:    lgr,
				Client: httpClient,
			},
			siteLabel: link.Host,
			withDelay: false,
		},
		lgr:       lgr,
		ch:        ch,
		store:     store,
		withDelay: false,
		name:      "getRockMusic",
	}
	_, err = scheduler.
		Every(17).
		Minutes().
		DoWithJobDetails(func(job gocron.Job) {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						lgr.Logf("[ERROR] PANIC recovered in GetRockMusic scraper: %v", r)
					}
				}()
				sErr := s.Scrape(ctx)
				if sErr != nil {
					lgr.Logf("[ERROR] failed to scrape %s: %v", s.name, sErr)
				}
				lgr.Logf("[INFO] GetRockMusicParser done, next run %v", job.NextRun())
			}()
		})
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}
	return nil
}

func runAlterPortal(
	ctx context.Context,
	lgr lgr.L,
	store *Store,
	httpClient *http.Client,
	ch chan []News,
	scheduler *gocron.Scheduler,
) error {
	link, err := url.Parse(AlterPortalParserRssURL)
	if err != nil {
		return err
	}
	parser := gofeed.NewParser()
	parser.Client = httpClient
	s := &Scraper{
		parser: &Parser{
			url:        AlterPortalParserRssURL,
			feedParser: parser,
			store:      store,
			lgr:        lgr,
			itemParser: &AlterPortalParser{
				Lgr:    lgr,
				Client: httpClient,
			},
			siteLabel: link.Host,
			withDelay: false,
		},
		lgr:       lgr,
		ch:        ch,
		store:     store,
		withDelay: false,
		name:      "alterPortal",
	}
	_, err = scheduler.
		Every(15).
		Minutes().
		DoWithJobDetails(func(job gocron.Job) {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						lgr.Logf("[ERROR] PANIC recovered in AlterPortal scraper: %v", r)
					}
				}()
				sErr := s.Scrape(ctx)
				if sErr != nil {
					lgr.Logf("[ERROR] failed to scrape %s: %v", s.name, sErr)
				}
				lgr.Logf("[INFO] AlterPortalParser done, next run %v", job.NextRun())
			}()
		})
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}
	return nil
}

func createPublisher(
	ctx context.Context,
	lgr lgr.L,
	store *Store,
	bot *tb.Bot,
	opt *Options,
	ch chan []News,
) (*Publisher, error) {
	notifyBot, err := NewBotAPI(BotAPIParams{
		Bot:     bot,
		ChantID: tb.ChatID(opt.NewsChatID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create bot api: %w", err)
	}

	notifyRetryBot, err := NewRetryableBotApi(RetryableBotApiParams{
		Lgr: lgr,
		Bot: notifyBot,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create retryable bot api: %w", err)
	}

	publisher, err := NewPublisher(PublisherParams{
		Lgr:    lgr,
		NewsCh: ch,
		BotAPI: notifyRetryBot,
		Store:  store,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create publisher: %w", err)
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				lgr.Logf("[ERROR] PANIC recovered in publisher: %v", r)
			}
		}()
		sErr := publisher.Start(ctx)
		if sErr != nil {
			lgr.Logf("[ERROR] failed to start publisher: %v", sErr)
		}
	}()

	return publisher, nil
}
