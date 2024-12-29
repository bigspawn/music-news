package internal

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
)

const CoreRadioParserRssURL = "https://coreradio.online/rss.xml"

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

	news, err := NewNewsFromItem(item)
	if err != nil {
		return nil, fmt.Errorf("coreradio: failed to create news from item: %w", err)
	}

	return ParseHtml(ctx, p.Lgr, news, resp.Body)
}

func ParseHtml(ctx context.Context, l lgr.L, news *News, r io.Reader) (*News, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("coreradio: NewDocumentFromReader: %w", err)
	}

	// image link
	imageLink, ok := doc.
		Find("#dle-content > div.full-news > div.full-news-top > div.full-news-left > center > a > img").
		Attr("src")
	if ok {
		il, err := url.Parse(imageLink)
		if err != nil {
			return nil, fmt.Errorf("coreradio: parse image link: %w", err)
		}

		if il.Query().Get("url") != "" {
			imageLink, err = DecodeBase64(il.Query().Get("url"))
			if err != nil {
				return nil, fmt.Errorf("coreradio: decode image base64: %w", err)
			}

			imageLink = ExtractAfterDecode(imageLink)
		}

		news.ImageLink = WebpToPng(imageLink)
	}

	// download link
	doc.
		Find(".quotel").
		Find("a[href]").
		Each(DownloadLinkSelector(news))

	if len(news.DownloadLink) == 0 {
		return nil, fmt.Errorf("coreradio: download link not found: %w", ErrSkipItem)
	}

	var links []string
	for i := range news.DownloadLink {
		if !strings.Contains(news.DownloadLink[i], engineSuffix) {
			l.Logf("[INFO] skip wrong link for parser: %s", news.DownloadLink[i])
			continue
		}

		l.Logf("[DEBUG] link: %s\n", news.DownloadLink[i])

		link, err := DecodeBase64(ExtractLink(news.DownloadLink[i]))
		if err != nil {
			return nil, fmt.Errorf("coreradio: DecodeBase64: link=%s: %w", news.DownloadLink[i], err)
		}

		l.Logf("[DEBUG] decoded link: %s\n", link)

		purl, err := url.ParseQuery(link)
		if err != nil {
			return nil, fmt.Errorf("coreradio: ParseQuery: link=%s: %w", link, err)
		}

		l.Logf("[DEBUG] parsed link: %s\n", purl)

		if purl.Get("url") == "" {
			l.Logf("[INFO] skip wrong link for parser: %s", news.DownloadLink[i])
			continue
		}

		ll := ExtractAfterDecode(purl.Get("url"))

		l.Logf("[DEBUG] extracted link: %s\n", ll)

		links = append(links, ll)
	}
	news.DownloadLink = links

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

	news.Text = moreThan2NewLinesRegexp.ReplaceAllString(news.Text, "\n")

	b = &strings.Builder{}
	for _, s := range strings.Split(news.Text, "\n") {
		if s[0] != '.' {
			b.WriteString(strings.TrimSpace(s))
			b.WriteRune('\n')
		}
	}
	news.Text = strings.TrimSpace(b.String())

	if isSkippedGenre(l, news.Text) {
		return nil, fmt.Errorf("coreradio: genre must be skipped: %w", ErrSkipItem)
	}

	return news, nil
}

func NewNewsFromItem(item *gofeed.Item) (*News, error) {
	itemDoc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(item.Description))
	if err != nil {
		return nil, fmt.Errorf("coreradio: NewDocumentFromReader: Description: %w", err)
	}

	imageLink, ok := itemDoc.Find("img[src]").Attr("src")
	if !ok {
		return nil, fmt.Errorf("coreradio: Find: image: not found: link=%s", item.Link)
	}

	return &News{
		Title:     strings.TrimSpace(item.Title),
		PageLink:  item.Link,
		DateTime:  *item.PublishedParsed,
		ImageLink: WebpToPng(imageLink),
	}, nil
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

const engineSuffix = "/engine/go.php?url="

func ExtractLink(s string) string {
	const (
		engineURL   = "https://" + coreradioHost + engineSuffix
		equalSymbol = "%3D"
		slash       = "%2F"
		slashLen    = len(slash)
	)
	s = strings.TrimPrefix(s, engineURL)
	s = strings.TrimSuffix(s, equalSymbol)
	idx := strings.Index(s, slash)
	if idx == -1 {
		return s
	}
	return s[idx+slashLen:]
}

func ExtractAfterDecode(s string) string {
	const (
		prefixS      = "s="
		questionRune = '?'
	)
	s = strings.TrimLeft(s, prefixS)
	idx := strings.IndexRune(s, questionRune)
	if idx == -1 {
		return s
	}
	return s[:idx]
}
