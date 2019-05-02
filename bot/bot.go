package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
)

func Create(user, password, address, token string) (*tgbotapi.BotAPI, error) {
	dialer, err := proxy.SOCKS5("tcp", address, &proxy.Auth{User: user, Password: password}, proxy.Direct)
	handleError(err)
	client := &http.Client{Transport: &http.Transport{Dial: dialer.Dial}}
	bot, err := tgbotapi.NewBotAPIWithClient(token, client)
	return bot, err
}

func SendNews(chatId int64, message string, news *gofeed.Item, bot *tgbotapi.BotAPI) bool {
	msg := tgbotapi.NewMessage(chatId, message)
	buttonURL := tgbotapi.NewInlineKeyboardButtonURL("Site", news.Link)
	markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttonURL))
	msg.ReplyMarkup = markup
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func SendImage(chatId int64, src string, bot *tgbotapi.BotAPI) bool {
	image := tgbotapi.NewPhotoShare(chatId, src)
	_, err := bot.Send(image)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("[ERROR] Error %s", err)
	}
}
