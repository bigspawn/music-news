package internal

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const GetRockMusicRss = "https://getrockmusic.net/rss.xml"

func NewGetRockMusicParser(lgr lgr.L, client *http.Client) ItemParser {
	return &getRockMusicParser{
		lgr:    lgr,
		client: client,
	}
}

type getRockMusicParser struct {
	lgr    lgr.L
	client *http.Client
}

func (p getRockMusicParser) Parse(ctx context.Context, item *gofeed.Item) (*News, error) {
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

	res, err := p.client.Do(req)
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
	content := doc.Find("#dle-content > div:nth-child(2)")

	var exists bool
	news.ImageLink, exists = content.Find("img[src]").Attr("src")
	if !exists {
		return nil, errors.New("image src not exists")
	}

	content.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		if href == "" {
			return
		}

		if isAllowedFileHost(href) {
			news.DownloadLink = append(news.DownloadLink, href)
		}
	})

	if len(news.DownloadLink) == 0 {
		cntHtml, err := content.Html()
		if err == nil {
			p.lgr.Logf("[ERROR] download links not found: %s", cntHtml)
		}
		return nil, errSkipItem
	}

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

	news.Text = news.Text[strings.Index(news.Text, "\n")+1:]

	return news, nil
}
