package main

import (
	botWrap "github.com/bigspawn/music-news/bot"
	"github.com/bigspawn/music-news/parser"
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
)

func handleError(err error) {
	if err != nil {
		log.Fatalf("[ERROR] Error %s", err)
	}
}

func main() {
	bot, err := botWrap.Create(user, passwd, address, botID)
	handleError(err)

	news, err := parser.Parse(feedUrl)
	handleError(err)

	for _, n := range news {
		botWrap.SendImage(chatIDDev, n, bot)
		botWrap.SendNews(chatIDDev, n, bot)
		log.Printf("[INFO] Item was send %s", n)
	}
}
