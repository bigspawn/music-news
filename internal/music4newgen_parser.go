package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const Music4newgenRSSFeedURL = "https://music4newgen.org/rss.xml"

func NewMusic4newgen(lgr lgr.L, client *http.Client) ItemParser {
	return &m4ngParser{
		lgr:    lgr,
		client: client,
	}
}

type m4ngParser struct {
	lgr    lgr.L
	client *http.Client
}

func (m m4ngParser) Parse(ctx context.Context, item *gofeed.Item) (*News, error) {
	news := &News{
		Title:    strings.TrimSpace(item.Title),
		PageLink: item.Link,
		DateTime: item.PublishedParsed,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, item.Link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,ru-RU;q=0.5,ru;q=0.3")

	res, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response not 200: code=%d; status=%s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	content := doc.Find(".full-story > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(1) > div:nth-child(2)")

	imageLink, exists := content.Find("img[src]").Attr("src")
	if !exists {
		return nil, errors.New("image src not exists")
	}

	news.ImageLink = "https://music4newgen.org" + imageLink

	content.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		if href == "" {
			return
		}

		prefix := "https://music4newgen.org/index.php?do=go&url="
		if strings.HasPrefix(href, prefix) {
			href = strings.TrimPrefix(href, prefix)
			index := strings.Index(href, "%")
			if index != -1 {
				href = href[:index]
			}
			bytes, err := base64.StdEncoding.DecodeString(href)
			if err != nil {
				m.lgr.Logf("[ERROR] decode base64 %w", err)
			}
			href = string(bytes)
			href = strings.TrimSuffix(href, ".ht") + ".html"
		}

		if isAllowedFileHost(href) {
			news.DownloadLink = append(news.DownloadLink, href)
		}
	})

	builder := &strings.Builder{}
	for _, node := range content.Nodes {
		findText(node, builder)
	}
	news.Text = strings.TrimSpace(builder.String())

	if isSkippedGender(news.Text) {
		return nil, errSkipItem
	}

	news.Text = newLinesRE.ReplaceAllString(news.Text, "\n")
	news.Text = trimLast(news.Text)

	return news, nil
}