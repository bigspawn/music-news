package internal

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const (
	GetRockMusicHost         = "https://getrockmusic.net"
	GetRockMusicParserRssURL = GetRockMusicHost + "/rss.xml"
)

type GetRockMusicParser struct {
	Lgr    lgr.L
	Client *http.Client
}

func (p *GetRockMusicParser) Parse(ctx context.Context, item *gofeed.Item) (*News, error) {
	news := NewNewsFromItem(item)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, item.Link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,ru-RU;q=0.5,ru;q=0.3")

	res, err := p.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response not 200: code=%d; status=%s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var exists bool
	news.ImageLink, exists = doc.Find(".fscover").Find("img[src]").Attr("src")
	if !exists {
		return nil, errors.New("image src not exists")
	}

	if strings.HasPrefix(news.ImageLink, "/uploads/") {
		news.ImageLink = GetRockMusicHost + news.ImageLink
	}

	content := doc.Find("div.generalblock:nth-child(3)")
	content.Find("a[href]").Each(DownloadLinkSelector(news))

	if len(news.DownloadLink) == 0 {
		cntHtml, err := content.Html()
		if err == nil {
			p.Lgr.Logf("[ERROR] download Links not found: %s", cntHtml)
		}
		return nil, ErrSkipItem
	}

	builder := &strings.Builder{}
	for _, node := range content.Nodes {
		findText(node, builder)
	}
	news.Text = strings.TrimSpace(builder.String())

	if isSkippedGender(p.Lgr, news.Text) {
		return nil, ErrSkipItem
	}

	news.Text = regexpNL.ReplaceAllString(news.Text, "\n")
	news.Text = trimLast(news.Text)

	news.Text = news.Text[strings.Index(news.Text, "\n")+1:]

	return news, nil
}

func DownloadLinkSelector(news *News) func(_ int, s *goquery.Selection) {
	return func(_ int, s *goquery.Selection) {
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
	}
}
