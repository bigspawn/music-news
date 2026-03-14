package internal

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strings"
	"testing"

	goOdesli "github.com/bigspawn/go-odesli"
	"github.com/go-pkgz/lgr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tb "gopkg.in/telebot.v3"
)

type mockBotSender struct {
	sendNewsFn        func(ctx context.Context, n News) (int, error)
	sendReleaseNewsFn func(ctx context.Context, n ReleaseNews) (int, error)
	deleteFn          func(ctx context.Context, id int) error
}

func (m *mockBotSender) SendNews(ctx context.Context, n News) (int, error) {
	return m.sendNewsFn(ctx, n)
}

func (m *mockBotSender) SendReleaseNews(ctx context.Context, n ReleaseNews) (int, error) {
	return m.sendReleaseNewsFn(ctx, n)
}

func (m *mockBotSender) Delete(ctx context.Context, id int) error {
	return m.deleteFn(ctx, id)
}

func TestRetryableSendReleaseNews_FloodError_DeletesImageBeforeRetry(t *testing.T) {
	callOrder := make([]string, 0)
	callCount := 0

	mock := &mockBotSender{
		sendReleaseNewsFn: func(_ context.Context, _ ReleaseNews) (int, error) {
			callCount++
			if callCount == 1 {
				callOrder = append(callOrder, "send_1_image_sent")
				// Image sent (id=42), text message failed with FloodError
				return 42, fmt.Errorf("%w", tb.FloodError{RetryAfter: 0})
			}
			callOrder = append(callOrder, "send_2_success")
			return 100, nil
		},
		deleteFn: func(_ context.Context, id int) error {
			callOrder = append(callOrder, fmt.Sprintf("delete_%d", id))
			return nil
		},
	}

	api, err := NewRetryableBotApi(RetryableBotApiParams{
		Lgr: lgr.Default(),
		Bot: mock,
	})
	require.NoError(t, err)

	n := ReleaseNews{
		News:          News{Title: "Test - Album"},
		ReleaseLink:   "https://song.link/test",
		PlatformLinks: map[goOdesli.Platform]string{goOdesli.PlatformSpotify: "https://spotify.com/test"},
	}

	err = api.SendReleaseNews(context.Background(), n)
	require.NoError(t, err)

	// delete must happen BEFORE the retry send
	require.Equal(t, []string{"send_1_image_sent", "delete_42", "send_2_success"}, callOrder)
}

func TestRetryableSendReleaseNews_NoFloodError_NoRetry(t *testing.T) {
	mock := &mockBotSender{
		sendReleaseNewsFn: func(_ context.Context, _ ReleaseNews) (int, error) {
			return 100, nil
		},
		deleteFn: func(_ context.Context, _ int) error {
			t.Fatal("delete should not be called on success")
			return nil
		},
	}

	api, err := NewRetryableBotApi(RetryableBotApiParams{Lgr: lgr.Default(), Bot: mock})
	require.NoError(t, err)

	err = api.SendReleaseNews(context.Background(), ReleaseNews{News: News{Title: "Test"}})
	require.NoError(t, err)
}

func TestRetryableSendReleaseNews_NonFloodError_NoRetry(t *testing.T) {
	mock := &mockBotSender{
		sendReleaseNewsFn: func(_ context.Context, _ ReleaseNews) (int, error) {
			return 0, fmt.Errorf("some other error")
		},
		deleteFn: func(_ context.Context, _ int) error {
			t.Fatal("delete should not be called when id=0")
			return nil
		},
	}

	api, err := NewRetryableBotApi(RetryableBotApiParams{Lgr: lgr.Default(), Bot: mock})
	require.NoError(t, err)

	err = api.SendReleaseNews(context.Background(), ReleaseNews{News: News{Title: "Test"}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "some other error")
}

func TestRetryableSendNews_FloodError_DeletesImageBeforeRetry(t *testing.T) {
	callOrder := make([]string, 0)
	callCount := 0

	mock := &mockBotSender{
		sendNewsFn: func(_ context.Context, _ News) (int, error) {
			callCount++
			if callCount == 1 {
				callOrder = append(callOrder, "send_1_image_sent")
				return 42, fmt.Errorf("%w", tb.FloodError{RetryAfter: 0})
			}
			callOrder = append(callOrder, "send_2_success")
			return 100, nil
		},
		deleteFn: func(_ context.Context, id int) error {
			callOrder = append(callOrder, fmt.Sprintf("delete_%d", id))
			return nil
		},
	}

	api, err := NewRetryableBotApi(RetryableBotApiParams{Lgr: lgr.Default(), Bot: mock})
	require.NoError(t, err)

	err = api.SendNews(context.Background(), News{Title: "Test"})
	require.NoError(t, err)
	require.Equal(t, []string{"send_1_image_sent", "delete_42", "send_2_success"}, callOrder)
}

func Test_truncateText(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		maxLen int
		want   string
	}{
		{
			name:   "short text unchanged",
			text:   "Hello world",
			maxLen: 100,
			want:   "Hello world",
		},
		{
			name:   "exact limit unchanged",
			text:   "12345",
			maxLen: 5,
			want:   "12345",
		},
		{
			name:   "long text truncated at last newline",
			text:   "Line 1\nLine 2\nLine 3\nLine 4\nLine 5",
			maxLen: 20,
			want:   "Line 1\nLine 2",
		},
		{
			name:   "no newline in text truncated hard",
			text:   "A very long line without any newlines at all",
			maxLen: 20,
			want:   "A very long line wit",
		},
		{
			name:   "telegram message limit with tracklist",
			text:   "Title\n" + strings.Repeat("1. Track Name Here\n", 250),
			maxLen: 4096,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncateText(tt.text, tt.maxLen)
			require.LessOrEqual(t, len(got), tt.maxLen, "result exceeds max length")
			if tt.want != "" {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_buildNewsText(t *testing.T) {
	longText := strings.Repeat("1. Some Track Title By Artist Name\n", 200)

	text := buildNewsText("Band - Album (2026)", longText)
	require.LessOrEqual(t, len(text), maxTelegramMessageLength,
		"news text must not exceed Telegram limit")
	require.True(t, strings.HasPrefix(text, "Band - Album (2026)\n"),
		"title must be preserved")
}

func Test_buildReleaseNewsText(t *testing.T) {
	longText := strings.Repeat("1. Some Track Title By Artist Name\n", 200)

	text := buildReleaseNewsText("Band - Album (2026)", longText, "https://song.link/test")
	require.LessOrEqual(t, len(text), maxTelegramMessageLength,
		"release news text must not exceed Telegram limit")
	require.True(t, strings.HasPrefix(text, "Band - Album (2026)\n"),
		"title must be preserved")
	require.True(t, strings.HasSuffix(text, "\n<a href=\"https://song.link/test\">Release album link</a>"),
		"release link must be preserved")
}

func createTestPNG(t *testing.T, width, height int) []byte {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			img.Set(x, y, color.RGBA{R: uint8(x % 256), G: uint8(y % 256), B: 128, A: 255})
		}
	}
	var buf bytes.Buffer
	require.NoError(t, png.Encode(&buf, img))
	return buf.Bytes()
}

func Test_resizeImage_LargeImage(t *testing.T) {
	data := createTestPNG(t, 2400, 2400)

	resized, err := resizeImage(data)
	require.NoError(t, err)
	require.NotEmpty(t, resized)

	// verify output is a valid image with correct width
	img, _, err := image.Decode(bytes.NewReader(resized))
	require.NoError(t, err)
	require.Equal(t, maxImageWidth, img.Bounds().Dx(), "width should be resized to maxImageWidth")
}

func Test_resizeImage_SmallImage(t *testing.T) {
	data := createTestPNG(t, 600, 600)

	resized, err := resizeImage(data)
	require.NoError(t, err)
	require.NotEmpty(t, resized)
}

func Test_resizeImage_ExactWidth(t *testing.T) {
	data := createTestPNG(t, maxImageWidth, 800)

	resized, err := resizeImage(data)
	require.NoError(t, err)
	require.NotEmpty(t, resized)
}
