package internal

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
)

const CoreRadioParserRssURL = "https://coreradio.ru/rss.xml"

type CoreRadioParser struct {
	Client *http.Client
	Lgr    lgr.L
}

func (p *CoreRadioParser) Parse(ctx context.Context, item *gofeed.Item) (*News, error) {
	resp, err := GetPage(ctx, p.Client, item.Link)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("coreradio: response is not 200 OK: status=%s, link=%s", resp.Status, item.Link)
	}

	news := NewNewsFromItem(item)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("coreradio: NewDocumentFromReader: %w", err)
	}

	// image link
	var ok bool
	news.ImageLink, ok = doc.
		Find("#dle-content > div.full-news > div.full-news-top > div.full-news-left > center > a > img").
		Attr("src")
	if !ok {
		itemDoc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(item.Description))
		if err != nil {
			return nil, fmt.Errorf("coreradio: NewDocumentFromReader: Description: %w", err)
		}

		news.ImageLink, ok = itemDoc.Find("img[src]").Attr("src")
		if !ok {
			return nil, fmt.Errorf("coreradio: Find: image: not found: link=%s", item.Link)
		}
	}
	news.ImageLink = WebpToPng(news.ImageLink)

	// download link
	doc.
		Find("#dle-content > div.full-news > div.full-news-top > div.full-news-right > center > div").
		Find("a[href]").
		Each(DownloadLinkSelector(news))

	if len(news.DownloadLink) == 0 {
		return nil, ErrSkipItem
	}

	for i := range news.DownloadLink {
		link, err := DecodeBase64(ExtractLink(news.DownloadLink[i]))
		if err != nil {
			return nil, fmt.Errorf("coreradio: DecodeBase64: link=%s", item.Link)
		}
		news.DownloadLink[i] = ExtractAfterDecode(link)
	}

	// text
	content := doc.Find("#dle-content > div.full-news > div.full-news-top > div.full-news-right > div.full-news-info")

	b := &strings.Builder{}
	for _, n := range content.Nodes {
		findText(n, b)
	}

	news.Text = strings.TrimSpace(b.String())

	var last int
	for i, r := range news.Text {
		if r == '\n' && i < len(news.Text)-1 {
			last = i
		}
	}

	if last > 0 {
		news.Text = news.Text[:last]
	}

	if isSkippedGenre(p.Lgr, news.Text) {
		return nil, ErrSkipItem
	}

	news.Text = moreThan2NewLinesRegexp.ReplaceAllString(news.Text, "\n")

	b = &strings.Builder{}
	for _, s := range strings.Split(news.Text, "\n") {
		if s[0] != '.' {
			b.WriteString(strings.TrimSpace(s))
			b.WriteRune('\n')
		}
	}
	news.Text = strings.TrimSpace(b.String())

	return news, nil
}

func NewNewsFromItem(item *gofeed.Item) *News {
	return &News{
		Title:    strings.TrimSpace(item.Title),
		PageLink: item.Link,
		DateTime: *item.PublishedParsed,
	}
}

func GetPage(ctx context.Context, client *http.Client, link string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/png,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,ru-RU;q=0.5,ru;q=0.3")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DecodeBase64(s string) (string, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ExtractLink(s string) string {
	const (
		engineURL   = "https://coreradio.ru/engine/go.php?url="
		equalSymbol = "%3D"
		slash       = "%2F"
		slashLen    = len(slash)
	)
	s = strings.TrimPrefix(s, engineURL)
	s = strings.TrimSuffix(s, equalSymbol)
	s = s[strings.Index(s, slash)+slashLen:]
	return s
}

func ExtractAfterDecode(s string) string {
	const (
		prefixS      = "s="
		questionRune = '?'
	)
	s = strings.TrimLeft(s, prefixS)
	s = s[:strings.IndexRune(s, questionRune)]
	return s
}
