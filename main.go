package main

import (
	"log"
	"os"

	botWrap "github.com/bigspawn/music-news/bot"
	"github.com/bigspawn/music-news/db"
	"github.com/bigspawn/music-news/parser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jasonlvhit/gocron"
	"github.com/jessevdk/go-flags"
)

type EnvParams struct {
	User     string `long:"user" env:"USER" required:"true"`
	Passwd   string `long:"passwd" env:"PASSWD" required:"true"`
	BotID    string `long:"bot_id" env:"BOT_ID" required:"true"`
	ChatID   int64  `long:"chat_id" env:"CHAT_ID" required:"true"`
	FeedURL  string `long:"feed_url" env:"FEED_URL" required:"true"`
	ProxyURL string `long:"proxy_url" env:"PROXY_URL" required:"true"`
	DbURL    string `long:"db_url" env:"DB_URL" required:"true"`
}

func main() {

	var params EnvParams
	p := flags.NewParser(&params, flags.Default)
	if _, err := p.Parse(); err != nil {
		os.Exit(0)
	}

	log.Printf("[DEBUG] %v", params)

	bot, err := botWrap.Create(params.User, params.Passwd, params.ProxyURL, params.BotID)
	if err != nil {
		log.Fatalf("[ERROR] Error %e", err)
	}
	gocron.Every(10).Minutes().Do(parse, bot, params)
	gocron.RunAll()
	_, time := gocron.NextRun()
	log.Printf("[INFO] Next start [%v]", time)
	<-gocron.Start()
}

func parse(bot *tgbotapi.BotAPI, params EnvParams) {
	con := db.Connection(params.DbURL)
	defer con.Close()
	news, err := parser.Parse(params.FeedURL, con)
	if err != nil {
		log.Printf("[ERROR] Error %v. Waiting for next execution", err)
		return
	}
	for _, n := range news {
		if wasSend := botWrap.SendImage(params.ChatID, n, bot); wasSend {
			botWrap.SendNews(params.ChatID, n, bot)
			log.Printf("[INFO] Item was send [%v]", n.Title)
		}
	}
}
