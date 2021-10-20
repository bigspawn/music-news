package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pkgz/lgr"
	tb "gopkg.in/tucnak/telebot.v2"
)

type RetryableBotApi struct {
	Bot BotAPI
	Lgr lgr.L
}

func (api RetryableBotApi) SendNews(ctx context.Context, n News) error {
	id, err := api.Bot.SendNews(ctx, n)
	if err == nil {
		return nil
	}

	retryErr := api.retry(ctx, n.Title, err, func() error { return api.SendNews(ctx, n) })
	if retryErr == nil {
		return nil
	}

	if id > 0 {
		_ = api.Delete(ctx, id)
	}

	return retryErr
}

func (api RetryableBotApi) Delete(ctx context.Context, id int) error {
	err := api.Bot.Delete(ctx, id)
	if err == nil {
		return nil
	}

	return api.retry(ctx, id, err, func() error { return api.Delete(ctx, id) })
}

func (api *RetryableBotApi) SendReleaseNews(ctx context.Context, n ReleaseNews) error {
	id, err := api.Bot.SendReleaseNews(ctx, n)
	if err == nil {
		return nil
	}

	retryErr := api.retry(ctx, n.Title, err, func() error { return api.SendReleaseNews(ctx, n) })
	if retryErr == nil {
		return nil
	}

	if id > 0 {
		_ = api.Delete(ctx, id)
	}

	return retryErr
}

func (api RetryableBotApi) retry(ctx context.Context, info interface{}, err error, f func() error) error {
	floodErr, ok := err.(tb.FloodError)
	if !ok {
		return err
	}

	duration := time.Duration(floodErr.RetryAfter) * time.Second

	api.Lgr.Logf("[DEBUG] sleep %s, title [%s]", duration, info)

	WaitUntil(ctx, duration)

	if err = f(); err != nil {
		return api.retry(ctx, info, err, f)
	}

	return nil
}

type BotAPI struct {
	Bot     *tb.Bot
	ChantID tb.ChatID
}

func (api *BotAPI) SendNews(ctx context.Context, n News) (int, error) {
	id, err := api.SendImageByLink(ctx, n.ImageLink)
	if err != nil {
		return 0, err
	}

	pageLinkURL, err := EncodeQuery(n.PageLink)
	if err != nil {
		return 0, err
	}

	downloadLinkURL, err := EncodeQuery(n.DownloadLink[0])
	if err != nil {
		return 0, err
	}

	text := fmt.Sprintf("%s\n%s", n.Title, n.Text)

	msg, err := api.Bot.Send(api.ChantID, text, &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{
				{Text: "Site Page", URL: pageLinkURL},
				{Text: "Download", URL: downloadLinkURL},
			},
		},
	})
	if err != nil {
		return id, err
	}

	return msg.ID, nil
}

func (api *BotAPI) SendImageByLink(ctx context.Context, imageLink string) (int, error) {
	link, err := EncodeQuery(imageLink)
	if err != nil {
		return 0, err
	}

	msg, err := api.Bot.Send(api.ChantID, &tb.Photo{File: tb.FromURL(link)})
	if err != nil {
		return 0, err
	}

	return msg.ID, nil
}

func (api *BotAPI) Delete(ctx context.Context, id int) error {
	return api.Bot.Delete(&tb.Message{ID: id, Chat: &tb.Chat{ID: int64(api.ChantID)}})
}

type ReleaseNews struct {
	News

	ReleaseLink   string
	PlatformLinks map[Platform]string
}

func (api *BotAPI) SendReleaseNews(ctx context.Context, n ReleaseNews) (int, error) {
	id, err := api.SendImageByLink(ctx, n.ImageLink)
	if err != nil {
		return 0, err
	}

	text := fmt.Sprintf("%s\n%s\n<a href=\"%s\">Release album link</a>", n.Title, n.Text, n.ReleaseLink)

	rows := make([]tb.InlineButton, 0, len(n.DownloadLink))
	for platform, link := range n.PlatformLinks {
		linkURL, err := EncodeQuery(link)
		if err != nil {
			return id, err
		}

		rows = append(rows, tb.InlineButton{Text: string(platform), URL: linkURL})
	}

	msg, err := api.Bot.Send(api.ChantID, text, &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{rows},
	})
	if err != nil {
		return id, err
	}

	return msg.ID, nil
}

func WaitUntil(ctx context.Context, duration time.Duration) {
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return
	case <-timer.C:
		return
	}
}
