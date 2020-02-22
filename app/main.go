package main

import (
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/jasonlvhit/gocron"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
)

const (
	period         = 30 // minutes
	timeoutBetween = 10 * time.Second
)

type params struct {
	User     string `long:"user" env:"USER"`
	Passwd   string `long:"passwd" env:"PASSWD"`
	ProxyURL string `long:"proxy_url" env:"PROXY_URL"`
	BotID    string `long:"bot_id" env:"BOT_ID" required:"true"`
	ChatID   int64  `long:"chat_id" env:"CHAT_ID" required:"true"`
	FeedURL  string `long:"feed_url" env:"FEED_URL" required:"true"`
	DbURL    string `long:"db_url" env:"DB_URL" required:"true"`
}

func main() {
	params := &params{}
	p := flags.NewParser(params, flags.Default)
	_, err := p.Parse()
	if err != nil {
		os.Exit(0)
	}

	log.Printf("[DEBUG] %v", params)

	bot, err := NewBotAPI(params)
	if err != nil {
		log.Fatalf("[ERROR] Error %e", err)
	}

	store, err := NewNewsStore(params.DbURL)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = store.conn.Close()
	}()

	parser := &SiteParser{
		Exclude: struct {
			Words     []string
			LastWords []string
			Genders   []string
		}{
			Words:     ExcludeWords,
			LastWords: ExcludeLastWords,
			Genders:   ExcludeGenders,
		},
		FeedParser: gofeed.NewParser(),
		Store:      store,
		URL:        params.FeedURL,
	}

	bot, err = NewBotAPI(params)
	if err != nil {
		log.Fatalf("[ERROR] Error %e", err)
	}

	gocron.Every(period).Minutes().Do(work, bot, parser)
	gocron.RunAll()

	_, time := gocron.NextRun()
	log.Printf("[INFO] Next start [%v]", time)

	<-gocron.Start()
}

func work(b *Bot, p *SiteParser) {
	news, err := p.Parse()
	if err != nil {
		log.Printf("[ERROR] Error %s. Waiting for next execution", err.Error())
		return
	}

	for _, n := range news {
		time.Sleep(timeoutBetween)

		err := b.SendImage(n)
		if err != nil {
			continue
		}
		_ = b.SendNews(n)
		log.Printf("[INFO] Item was send [%s]", n.Title)
	}
}
