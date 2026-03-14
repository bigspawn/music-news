package internal

import (
	"image"
	"image/color"
	"image/png"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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
