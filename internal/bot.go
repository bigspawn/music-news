package internal

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	goOdesli "github.com/bigspawn/go-odesli"
	"github.com/go-pkgz/lgr"
	tb "gopkg.in/telebot.v3"
)

type RetryableBotApiParams struct {
	Lgr lgr.L
	Bot *BotAPI
}

func (p *RetryableBotApiParams) Validate() error {
	if p.Lgr == nil {
		return errors.New("lgr is required")
	}

	if p.Bot == nil {
		return errors.New("bot is required")
	}
	return nil
}

type RetryableBotApi struct {
	RetryableBotApiParams
}

func NewRetryableBotApi(params RetryableBotApiParams) (*RetryableBotApi, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	return &RetryableBotApi{
		RetryableBotApiParams: params,
	}, nil
}

func (api RetryableBotApi) SendNews(ctx context.Context, n News) error {
	id, err := api.Bot.SendNews(ctx, n)
	if err == nil {
		return nil
	}

	if id > 0 {
		_ = api.Delete(ctx, id)
	}

	// Return image error without fallback
	var bErr *tb.Error
	if errors.As(err, &bErr) && bErr.Code == 400 &&
		(bErr.Message == "wrong type of the web page content" ||
		 bErr.Message == "failed to get HTTP URL content") {
		api.Lgr.Logf("[WARN] image failed for [%s]: %v", n.Title, err)
		return err
	}

	retryErr := api.retry(ctx, n.Title, err, func() error {
		return api.SendNews(ctx, n)
	})
	if retryErr == nil {
		return nil
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

	if id > 0 {
		_ = api.Delete(ctx, id)
	}

	// Return image error without fallback
	var bErr *tb.Error
	if errors.As(err, &bErr) && bErr.Code == 400 &&
		(bErr.Message == "wrong type of the web page content" ||
		 bErr.Message == "failed to get HTTP URL content") {
		api.Lgr.Logf("[WARN] release image failed for [%s]: %v", n.Title, err)
		return err
	}

	retryErr := api.retry(ctx, n.Title, err, func() error { return api.SendReleaseNews(ctx, n) })
	if retryErr == nil {
		return nil
	}

	return retryErr
}

func (api RetryableBotApi) retry(ctx context.Context, info interface{}, err error, f func() error) error {
	var floodErr tb.FloodError
	if !errors.As(err, &floodErr) {
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

type BotAPIParams struct {
	Bot     *tb.Bot
	ChantID tb.ChatID
}

func (p *BotAPIParams) Validate() error {
	if p.Bot == nil {
		return errors.New("bot is required")
	}
	if p.ChantID == 0 {
		return errors.New("chant_id is required")
	}
	return nil
}

type BotAPI struct {
	BotAPIParams
}

func NewBotAPI(params BotAPIParams) (*BotAPI, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &BotAPI{
		BotAPIParams: params,
	}, nil
}

func (api *BotAPI) SendNews(ctx context.Context, n News) (int, error) {
	id, err := api.SendImageByLink(ctx, n.ImageLink)
	if err != nil {
		return 0, fmt.Errorf("failed to send image: %w", err)
	}

	pageLinkURL, err := EncodeQuery(n.PageLink)
	if err != nil {
		return 0, fmt.Errorf("failed to encode page link: %w", err)
	}

	keyboard := [][]tb.InlineButton{{{Text: "Site Page", URL: pageLinkURL}}}
	for i, s := range n.DownloadLink {
		l, err := EncodeQuery(s)
		if err != nil {
			return 0, fmt.Errorf("failed to encode download link: %w", err)
		}

		if s == "" {
			continue
		}

		keyboard[0] = append(keyboard[0], tb.InlineButton{
			Text: fmt.Sprintf("Download_%d", i),
			URL:  l,
		})
	}

	text := fmt.Sprintf("%s\n%s", n.Title, n.Text)

	msg, err := api.Bot.Send(api.ChantID, text, &tb.ReplyMarkup{InlineKeyboard: keyboard})
	if err != nil {
		return id, fmt.Errorf("failed to send news: %w", err)
	}

	return msg.ID, nil
}

func (api *BotAPI) SendImageByLink(ctx context.Context, imageLink string) (int, error) {
	link, err := EncodeQuery(imageLink)
	if err != nil {
		return 0, fmt.Errorf("failed to encode image link: %w", err)
	}

	msg, err := api.Bot.Send(api.ChantID, &tb.Photo{File: tb.FromURL(link)})
	if err != nil {
		return 0, fmt.Errorf("failed to send image: %w", err)
	}

	return msg.ID, nil
}

func (api *BotAPI) Delete(ctx context.Context, id int) error {
	return api.Bot.Delete(&tb.Message{ID: id, Chat: &tb.Chat{ID: int64(api.ChantID)}})
}

type ReleaseNews struct {
	News

	ReleaseLink   string
	PlatformLinks map[goOdesli.Platform]string
}

func (api *BotAPI) SendReleaseNews(ctx context.Context, n ReleaseNews) (int, error) {
	id, err := api.SendImageByLink(ctx, n.ImageLink)
	if err != nil {
		return 0, fmt.Errorf("failed to send image: %w", err)
	}

	text := fmt.Sprintf("%s\n%s\n<a href=\"%s\">Release album link</a>", n.Title, n.Text, n.ReleaseLink)

	rows := make([]tb.InlineButton, 0, len(n.PlatformLinks))
	for platform, link := range n.PlatformLinks {
		linkURL, eErr := EncodeQuery(link)
		if eErr != nil {
			return id, fmt.Errorf("failed to encode platform link: %w", eErr)
		}

		rows = append(rows, tb.InlineButton{Text: string(platform), URL: linkURL})
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Text > rows[j].Text
	})

	msg, err := api.Bot.Send(api.ChantID, text, &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{rows},
	})
	if err != nil {
		return id, fmt.Errorf("failed to send message: %w", err)
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
