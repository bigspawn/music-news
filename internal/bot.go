package internal

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	goOdesli "github.com/bigspawn/go-odesli"
	"github.com/disintegration/imaging"
	"github.com/go-pkgz/lgr"
	tb "gopkg.in/telebot.v3"
)

type botSender interface {
	SendNews(ctx context.Context, n News) (int, error)
	SendReleaseNews(ctx context.Context, n ReleaseNews) (int, error)
	Delete(ctx context.Context, id int) error
}

type RetryableBotApiParams struct {
	Lgr lgr.L
	Bot botSender
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
	Lgr        lgr.L
	Bot        *tb.Bot
	ChantID    tb.ChatID
	HTTPClient *http.Client
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
	if params.HTTPClient == nil {
		params.HTTPClient = http.DefaultClient
	}
	return &BotAPI{
		BotAPIParams: params,
	}, nil
}

// SafeInlineButton creates a URL-only inline button that's safe for channel use.
// This avoids the telebot.v3 bug where InlineQueryChat field is sent even when empty.
func SafeInlineButton(text, url string) tb.InlineButton {
	return tb.InlineButton{
		Text: text,
		URL:  url,
	}
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

	keyboard := [][]tb.InlineButton{{SafeInlineButton("Site Page", pageLinkURL)}}
	for i, s := range n.DownloadLink {
		l, err := EncodeQuery(s)
		if err != nil {
			return 0, fmt.Errorf("failed to encode download link: %w", err)
		}

		if s == "" {
			continue
		}

		keyboard[0] = append(keyboard[0], SafeInlineButton(fmt.Sprintf("Download_%d", i), l))
	}

	text := buildNewsText(n.Title, n.Text)

	msg, err := api.Bot.Send(api.ChantID, text, &tb.ReplyMarkup{InlineKeyboard: keyboard})
	if err != nil {
		return id, fmt.Errorf("failed to send news: %w", err)
	}

	return msg.ID, nil
}

const (
	maxImageSize              = 10 * 1024 * 1024 // 10MB Telegram limit for photos
	maxTelegramMessageLength  = 4096
)

func truncateText(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}
	// try to cut at last newline within limit
	truncated := text[:maxLen]
	if idx := strings.LastIndex(truncated, "\n"); idx > 0 {
		return truncated[:idx]
	}
	return truncated
}

func buildNewsText(title, body string) string {
	text := fmt.Sprintf("%s\n%s", title, body)
	return truncateText(text, maxTelegramMessageLength)
}

func buildReleaseNewsText(title, body, releaseLink string) string {
	suffix := fmt.Sprintf("\n<a href=\"%s\">Release album link</a>", releaseLink)
	maxBody := maxTelegramMessageLength - len(title) - 1 - len(suffix) // 1 for \n between title and body
	if len(body) > maxBody {
		body = truncateText(body, maxBody)
	}
	return fmt.Sprintf("%s\n%s%s", title, body, suffix)
}

func (api *BotAPI) SendImageByLink(ctx context.Context, imageLink string) (int, error) {
	link, err := EncodeQuery(imageLink)
	if err != nil {
		return 0, fmt.Errorf("failed to encode image link: %w", err)
	}

	msg, err := api.Bot.Send(api.ChantID, &tb.Photo{File: tb.FromURL(link)})
	if err == nil {
		return msg.ID, nil
	}

	if !isTelegramImageError(err) {
		return 0, fmt.Errorf("failed to send image: %w", err)
	}

	// Fallback: download image ourselves and upload as file
	api.Lgr.Logf("[INFO] image FromURL failed for %s, falling back to upload: %v", imageLink, err)
	return api.sendImageByUpload(ctx, imageLink)
}

func isTelegramImageError(err error) bool {
	s := err.Error()
	return strings.Contains(s, "wrong type of the web page content") ||
		strings.Contains(s, "failed to get HTTP URL content") ||
		strings.Contains(s, "wrong file identifier/HTTP URL specified")
}

const maxImageWidth = 1200

func (api *BotAPI) sendImageByUpload(ctx context.Context, imageLink string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageLink, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create image request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0")

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to download image: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to download image: status=%s", resp.Status)
	}

	data, err := io.ReadAll(io.LimitReader(resp.Body, maxImageSize+1))
	if err != nil {
		return 0, fmt.Errorf("failed to read image body: %w", err)
	}

	resized, err := resizeImage(data)
	if err != nil {
		return 0, fmt.Errorf("failed to resize image: %w", err)
	}

	msg, err := api.Bot.Send(api.ChantID, &tb.Photo{File: tb.FromReader(bytes.NewReader(resized))})
	if err != nil {
		return 0, fmt.Errorf("failed to upload image: %w", err)
	}

	return msg.ID, nil
}

func resizeImage(data []byte) ([]byte, error) {
	img, err := imaging.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	if img.Bounds().Dx() > maxImageWidth {
		img = imaging.Resize(img, maxImageWidth, 0, imaging.Lanczos)
	}

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 85}); err != nil {
		return nil, fmt.Errorf("encode jpeg: %w", err)
	}

	return buf.Bytes(), nil
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

	text := buildReleaseNewsText(n.Title, n.Text, n.ReleaseLink)

	rows := make([]tb.InlineButton, 0, len(n.PlatformLinks))
	for platform, link := range n.PlatformLinks {
		linkURL, eErr := EncodeQuery(link)
		if eErr != nil {
			return id, fmt.Errorf("failed to encode platform link: %w", eErr)
		}

		rows = append(rows, SafeInlineButton(string(platform), linkURL))
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
