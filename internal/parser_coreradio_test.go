package internal

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/require"
)

func TestCoreRadioParser_Parse(t *testing.T) {
	t.Skip("local integration test")

	is := require.New(t)

	f, err := os.Open("../tmp/rss_coreradio.xml")
	is.NoError(err)

	feed, err := gofeed.NewParser().Parse(f)
	is.NoError(err)

	p := &CoreRadioParser{
		Client: http.DefaultClient,
		Lgr:    lgr.Default(),
	}

	for _, item := range feed.Items {
		n, err := p.Parse(context.Background(), item)
		// is.NoError(err)
		if err != nil {
			continue
		}
		t.Logf("%v\n", n)

		is.NotEmpty(n.Title)
		is.NotEmpty(n.PageLink)
		is.NotEmpty(n.DateTime)
		is.NotEmpty(n.ImageLink)
		is.NotEmpty(n.DownloadLink)
		is.NotEmpty(n.Text)
	}
}

func TestExtractLink(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "1",
			s:    "https://coreradio.ru/engine/go.php?url=aHR0cHM6Ly9vdW8uaW8vcXMvUmZGT3FUQjk%2Fcz1odHRwczovL3ZrLmNvbS9kb2MxNjAxMTgwMDJfNjE4NDc0MDAzPzk4NEs%3D",
			want: "cz1odHRwczovL3ZrLmNvbS9kb2MxNjAxMTgwMDJfNjE4NDc0MDAzPzk4NEs",
		},
		{
			name: "2",
			s:    "https://coreradio.ru/engine/go.php?url=aHR0cHM6Ly9vdW8uaW8vcXMvUmZGT3FUQjk%2Fcz1odHRwczovL3ZrLmNvbS9kb2M1ODM1MTY3MThfNjE1MzMwNjQ0PzM0MkE%3D",
			want: "cz1odHRwczovL3ZrLmNvbS9kb2M1ODM1MTY3MThfNjE1MzMwNjQ0PzM0MkE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractLink(tt.s); got != tt.want {
				t.Errorf("ExtractLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    string
		wantErr bool
	}{
		{
			name:    "1",
			s:       "cz1odHRwczovL3ZrLmNvbS9kb2M1ODM1MTY3MThfNjE1MzMwNjQ0PzM0MkE",
			want:    "s=https://vk.com/doc583516718_615330644?342A",
			wantErr: false,
		},
		{
			name:    "2",
			s:       "cz1odHRwczovL3ZrLmNvbS9kb2MxNjAxMTgwMDJfNjE4NDc0MDAzPzk4NEs",
			want:    "s=https://vk.com/doc160118002_618474003?984K",
			wantErr: false,
		},
		{
			name:    "3",
			s:       "aHR0cHM6Ly9jbG91ZC5tYWlsLnJ1L3B1YmxpYy9LdTVmL2trR0dnM1VzWg",
			want:    "https://cloud.mail.ru/public/Ku5f/kkGGg3UsZ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBase64(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeBase64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractAfterDecode(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "1", s: "s=https://vk.com/doc583516718_615330644?342A", want: "https://vk.com/doc583516718_615330644"},
		{name: "2", s: "s=https://vk.com/doc160118002_618474003?984K", want: "https://vk.com/doc160118002_618474003"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractAfterDecode(tt.s); got != tt.want {
				t.Errorf("ExtractAfterDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
