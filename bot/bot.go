package bot

import (
	"github.com/bigspawn/music-news/parser"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
)

func Create(user, password, address, token string) (*tgbotapi.BotAPI, error) {
	dialer, err := proxy.SOCKS5("tcp", address, &proxy.Auth{User: user, Password: password}, proxy.Direct)
	if err != nil {
		log.Fatalf("[ERROR] Error %s", err)
	}
	client := &http.Client{Transport: &http.Transport{Dial: dialer.Dial}}
	bot, err := tgbotapi.NewBotAPIWithClient(token, client)
	return bot, err
}

func SendNews(chatId int64, news parser.News, bot *tgbotapi.BotAPI) bool {
	pageButton := tgbotapi.NewInlineKeyboardButtonURL("Site page", news.PageLink)
	downloadButton := tgbotapi.NewInlineKeyboardButtonURL("Download", news.DownloadLink[0])
	markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(pageButton, downloadButton))
	msg := tgbotapi.NewMessage(chatId, news.Text)
	msg.ReplyMarkup = markup
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func SendImage(chatId int64, news parser.News, bot *tgbotapi.BotAPI) bool {
	image := tgbotapi.NewPhotoShare(chatId, news.ImageLink)
	_, err := bot.Send(image)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
