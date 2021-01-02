package internal

import (
	"context"
	"fmt"
	"github.com/go-pkgz/lgr"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"sort"
)

type TelegramBot struct {
	BotAPI *tgbotapi.BotAPI
	ChatId int64

	lgr lgr.L
}

func NewTelegramBotAPI(p *Options, lgr lgr.L) (*TelegramBot, error) {
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
	return &TelegramBot{
		BotAPI: bot,
		ChatId: p.ChatID,
		lgr:    lgr,
	}, nil

}

func (b *TelegramBot) SendNews(_ context.Context, item *News) error {
	b.lgr.Logf("[INFO] send news %v", item)

	pUrl, err := encodeQuery(item.PageLink)
	if err != nil {
		return err
	}

	_, err = b.BotAPI.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: b.ChatId,
			ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Site page", pUrl),
				tgbotapi.NewInlineKeyboardButtonURL("Download", item.DownloadLink[0]),
			)),
		},
		Text:                  fmt.Sprintf("%s\n%s", item.Title, item.Text),
		DisableWebPagePreview: false,
	})
	if err != nil {
		return err
	}
	return nil
}

func (b *TelegramBot) SendImage(_ context.Context, n *News) (int, error) {
	b.lgr.Logf("[INFO] send image %v", n)

	msg, err := b.BotAPI.Send(tgbotapi.NewPhotoShare(b.ChatId, n.ImageLink))
	if err != nil {
		return 0, err
	}
	return msg.MessageID, nil
}

func (b *TelegramBot) SendRelease(item *News, releaseLink string) error {
	b.lgr.Logf("[INFO] send release link %s, %v", releaseLink, item)

	id, err := b.SendImage(nil, item)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(b.ChatId, item.Title+"\n"+item.Text+"\nRelease album link: "+releaseLink)
	if _, err := b.BotAPI.Send(msg); err != nil {
		_, _ = b.BotAPI.DeleteMessage(tgbotapi.NewDeleteMessage(b.ChatId, id))
		return err
	}
	return nil
}

func (b *TelegramBot) SendReleaseWithButtons(item *News, releaseLink string, links map[Platform]string) error {
	b.lgr.Logf("[INFO] send release with links %v, %v", links, item)

	id, err := b.SendImage(nil, item)
	if err != nil {
		return err
	}

	var rows []tgbotapi.InlineKeyboardButton
	for p, l := range links {
		rows = append(rows, tgbotapi.NewInlineKeyboardButtonURL(string(p), l))
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Text > rows[j].Text
	})

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:      b.ChatId,
			ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(rows),
		},
		Text:                  fmt.Sprintf("%s\n%s\n<a href=\"%s\">Release album link</a>", item.Title, item.Text, releaseLink),
		DisableWebPagePreview: false,
		ParseMode:             tgbotapi.ModeHTML,
	}

	if _, err := b.BotAPI.Send(msg); err != nil {
		_, _ = b.BotAPI.DeleteMessage(tgbotapi.NewDeleteMessage(b.ChatId, id))
		return err
	}
	return nil
}

func encodeQuery(u string) (string, error) {
	uu, e := url.Parse(u)
	if e != nil {
		return "", e
	}
	uu.RawQuery = uu.Query().Encode()
	return uu.String(), nil
}
