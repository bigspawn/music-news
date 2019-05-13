package main

import (
	botWrap "github.com/bigspawn/music-news/bot"
	"github.com/bigspawn/music-news/db"
	"github.com/bigspawn/music-news/parser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jasonlvhit/gocron"
	"log"
)

const (
	user            = "bigspawn"
	passwd          = "BxYxvmFrohdDwsRjuUjAKXjPRwUvEpmB"
	botID           = "417691036:AAHaptah5zlg5GpRvHkWka_680ZL3_7MKhw"
	chatID    int64 = -1001112604760
	chatIDDev int64 = -1001086871152
	feedUrl         = "https://kingdom-leaks.com/index.php?/rss/3-the-kingdom-leaks-homepage-feed.xml"
	address         = "test.bigspawn.com:1080"
	dbUrlDev        = "postgres://go-music:mysecretpassword@localhost:8532/postgres?sslmode=disable"
	dbUrl           = "postgres://postgres:xxQWV2EkIUjzcyeag27AqTiNBjiupob8PHU@test.bigspawn.com:15432/music?sslmode=disable"
)

func main() {
	bot, err := botWrap.Create(user, passwd, address, botID)
	if err != nil {
		log.Fatalf("[ERROR] Error %v", err)
	}
	gocron.Every(10).Minutes().Do(parse, bot)
	//gocron.RunAll()
	_, time := gocron.NextRun()
	log.Println(time)
	<-gocron.Start()
}

func parse(bot *tgbotapi.BotAPI) {
	con := db.Connection(dbUrl)
	defer con.Close()
	news, err := parser.Parse(feedUrl, con)
	if err != nil {
		log.Fatalf("[ERROR] Error %v", err)
	}
	for _, n := range news {
		botWrap.SendImage(chatID, n, bot)
		botWrap.SendNews(chatID, n, bot)
		log.Printf("[INFO] Item was send [%v]", n)
	}
}
