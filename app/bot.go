package main

import (
	"context"
	"net"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/proxy"
)

var (
	newsNotified = promauto.NewCounter(prometheus.CounterOpts{
		Name: "notified_news_total",
		Help: "The total number of notified news",
	})
	newsSended = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sended_news_total",
		Help: "The total number of sended news",
	})
)

type TelegramBot struct {
	BotAPI *tgbotapi.BotAPI
	ChatId int64
}

func NewTelegramBotAPI(p *Options) (*TelegramBot, error) {
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
	return &TelegramBot{BotAPI: bot, ChatId: p.ChatID}, nil

}

func (b *TelegramBot) SendNews(item *News) error {
	Lgr.Logf("[INFO] send news %v", item)

	msg := tgbotapi.NewMessage(b.ChatId, item.Title+"\n"+item.Text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Site page", item.PageLink),
		tgbotapi.NewInlineKeyboardButtonURL("Download", item.DownloadLink[0]),
	))
	if _, err := b.BotAPI.Send(msg); err != nil {
		return err
	}

	newsSended.Inc()

	return nil
}

func (b *TelegramBot) SendImage(n *News) error {
	Lgr.Logf("[INFO] send image%v", n)

	photo := tgbotapi.NewPhotoShare(b.ChatId, n.ImageLink)
	_, err := b.BotAPI.Send(photo)
	if err != nil {
		return err
	}
	return nil
}

func (b *TelegramBot) SendRelease(item *News, releaseLink string) error {
	Lgr.Logf("[INFO] send release link %s, %v", releaseLink, item)

	if err := b.SendImage(item); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(b.ChatId, item.Title+"\n"+item.Text+"\nRelease album link: "+releaseLink)
	if _, err := b.BotAPI.Send(msg); err != nil {
		return err
	}

	newsNotified.Inc()

	return nil
}
