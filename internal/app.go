package internal

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
	"github.com/pkg/errors"
	"net/http"
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
		createNotifier(ctx, opt, lgr, store, bot, scheduler)
	} else {
		if err := createScrapes(ctx, lgr, store, bot, scheduler); err != nil {
			return nil, err
		}
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

func createScrapes(ctx context.Context, lgr lgr.L, store *Store, bot *TelegramBot, scheduler *gocron.Scheduler) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	alterportalRssFeedParser := NewRssFeedParser(AlterportalRSSFeedURL, store, lgr, NewAlterportalParser(lgr, client))
	alterportalScr := NewMusicScraper(bot, alterportalRssFeedParser, lgr, store)
	alterportalJob := NewJob(alterportalScr, scheduler, "alterportal", lgr)
	_, err := scheduler.Every(10).Minutes().Do(alterportalJob.Do, ctx)
	if err != nil {
		return err
	}

	music4newgenRssFeedParser := NewRssFeedParser(Music4newgenRSSFeedURL, store, lgr, NewMusic4newgen(lgr, client))
	music4newgenScr := NewMusicScraper(bot, music4newgenRssFeedParser, lgr, store)
	music4newgenJob := NewJob(music4newgenScr, scheduler, "music4newgen", lgr)
	_, err = scheduler.Every(11).Minutes().Do(music4newgenJob.Do, ctx)
	if err != nil {
		return err
	}
	return nil
}

func createNotifier(ctx context.Context, opt *Options, lgr lgr.L, store *Store, bot *TelegramBot, scheduler *gocron.Scheduler) {
	linkApi := NewLinkApi(opt.SongAPIKey, lgr)
	n := NewNotifier(store, bot, linkApi, lgr)
	task := func() {
		if err := n.Notify(ctx); err != nil {
			lgr.Logf("[ERROR] notifier %v", err)
		}

		_, next := scheduler.NextRun()
		lgr.Logf("[INFO] job next start %s", next)
	}
	scheduler.Every(3).Hour().Do(task)
}

func NewJob(s MusicScraper, sch *gocron.Scheduler, name string, lgr lgr.L) *Job {
	return &Job{
		s:    s,
		sch:  sch,
		name: name,
		lgr:  lgr,
	}
}

type Job struct {
	s    MusicScraper
	sch  *gocron.Scheduler
	name string
	lgr  lgr.L
}

func (j Job) Do(ctx context.Context) {
	if err := j.s.Scrape(ctx); err != nil {
		j.lgr.Logf("[ERROR] %s scraper %v", j.name, err)
	}

	_, next := j.sch.NextRun()
	j.lgr.Logf("[INFO] %s job next start %s", j.name, next)
}
