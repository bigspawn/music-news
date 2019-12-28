package main

import (
	"context"
	"log"
	"net"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
)

type Bot struct {
	BotAPI *tgbotapi.BotAPI
	ChatId int64
}

func NewBotAPI(p *params) (*Bot, error) {
	var err error
	var bot *tgbotapi.BotAPI
	if p.ProxyURL != "" {
		dialer, err := proxy.SOCKS5("tcp", p.ProxyURL, &proxy.Auth{User: p.User, Password: p.Passwd}, proxy.Direct)
		if err != nil {
			return nil, err
		}
		bot, err = tgbotapi.NewBotAPIWithClient(p.BotID, &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
					return dialer.Dial(network, addr)
				},
			},
		})
		if err != nil {
			return nil, err
		}
	} else {
		bot, err = tgbotapi.NewBotAPI(p.BotID)
		if err != nil {
			return nil, err
		}
	}
	return &Bot{BotAPI: bot, ChatId: p.ChatID}, nil

}

func (b *Bot) SendNews(news News) error {
	log.Printf("[INFO] send news in chat %d, %v", b.ChatId, news)

	msg := tgbotapi.NewMessage(b.ChatId, news.Title+"\n"+news.Text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Site page", news.PageLink),
			tgbotapi.NewInlineKeyboardButtonURL("Download", news.DownloadLink[0]),
		))

	_, err := b.BotAPI.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) SendImage(n News) error {
	log.Printf("[INFO] send image in chat %d, %v", b.ChatId, n)

	_, err := b.BotAPI.Send(tgbotapi.NewPhotoShare(b.ChatId, n.ImageLink))
	if err != nil {
		return err
	}
	return nil
}
