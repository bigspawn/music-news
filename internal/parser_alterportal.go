package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

const AlterPortalParserRssURL = "https://alterportal.net/rss.xml"

var moreThan2NewLinesRegexp = regexp.MustCompile("\n{2,}")

type AlterPortalParser struct {
	Lgr    lgr.L
	Client *http.Client
}

func (p *AlterPortalParser) Parse(ctx context.Context, item *gofeed.Item) (*News, error) {
	if strings.Contains(item.Link, "raznoe") ||
		strings.Contains(item.Link, "video") ||
		strings.Contains(item.Link, "news") ||
		strings.Contains(item.Link, "neformat") {
		return nil, ErrSkipItem
	}

	news := &News{
		Title:    strings.TrimSpace(item.Title),
		PageLink: item.Link,
		DateTime: *item.PublishedParsed,
	}

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
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response not 200: code=%d; status=%s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	if strings.Contains(title, "502: Bad gateway") {
		return nil, ErrSkipItem
	}

	isVideo := false
	doc.Find(".full_title").
		Find("a[href]").
		Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if !exists {
				return
			}

			if href == "" {
				return
			}

			if strings.HasPrefix(href, "https://alterportal.net/video") {
				isVideo = true
			}
		})
	if isVideo {
		return nil, ErrSkipItem
	}

	content := doc.Find(".ftwo")

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

		if !isAllowedFileHost(href) {
			return
		}

		if strings.Contains(href, alterportalHost) {
			s, err := ExtractLinkFromParamURL(href)
			if err != nil {
				return
			}
			href, err = DecodeBase64StdPadding(s)
			if err != nil {
				return
			}
		}

		news.DownloadLink = append(news.DownloadLink, href)
	})

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
	news.Text = builder.String()

	if isSkippedGender(p.Lgr, news.Text) {
		return nil, ErrSkipItem
	}

	news.Text = translate(news.Text)
	news.Text = moreThan2NewLinesRegexp.ReplaceAllString(news.Text, "\n")
	news.Text = strings.ReplaceAll(news.Text, "[ ] \n", "")
	news.Text = strings.ReplaceAll(news.Text, ":: ", "")
	news.Text = trimLast(news.Text)
	news.Text = strings.TrimSpace(news.Text)

	return news, nil
}

func ExtractLinkFromParamURL(s string) (string, error) {
	u, err := url.Parse(html.UnescapeString(s))
	if err != nil {
		return "", err
	}

	return u.Query().Get("url"), nil
}

func DecodeBase64StdPadding(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

var re = regexp.MustCompile("^\\d{1,3}[.\\s]*[\\s-–]")

func trimLast(text string) string {
	lines := strings.Split(text, "\n")
	var j int
	for i := len(lines) - 1; i >= 0; i-- {
		if re.MatchString(lines[i]) {
			j = i
			break
		}
	}
	text = ""
	for i := 0; i <= j; i++ {
		text += lines[i]
		text += "\n"
	}
	return text
}

func findText(node *html.Node, builder *strings.Builder) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode {
		data := strings.TrimSpace(node.Data)
		if needAddWord(data) {
			builder.WriteString(data)
			builder.WriteString(" ")
		}
	}
	if node.Type == html.ElementNode {
		if node.Data == "br" {
			builder.WriteString("\n")
		}
	}
	if node.FirstChild != nil {
		findText(node.FirstChild, builder)
	}
	if node.NextSibling != nil {
		findText(node.NextSibling, builder)
	}
}

func translate(data string) string {
	for k, v := range translateMap {
		data = strings.ReplaceAll(data, k, v)
	}
	return data
}

var translateMap = map[string]string{
	"Стиль":       "Genre",
	"Cтиль":       "Genre",
	"Страна":      "Country",
	"Дата релиза": "Release",
	"Год выпуска": "Release",
	"Формат":      "Quality",
	"Размер":      "Size",
	"Треклист":    "Tracklist",
	"Лейбл":       "Label",
	"Качество":    "Quality",
	"Исполнитель": "Artist",
	"Альбом":      "Album",
	"Дата Релиза": "Release",
	"Дата выхода": "Release",
	"Дата Выхода": "Release",
}

func needAddWord(data string) bool {
	if data == "" {
		return false
	}
	if _, ok := skipWords[strings.ToLower(data)]; ok {
		return false
	}
	return true
}

var skipWords = map[string]struct{}{
	"official website": {}, "facebook": {}, "download": {}, "zippyshare": {}, "yadisk": {}, "i": {}, "скачать!": {},
	"вк!": {}, "instagram": {}, "twitter": {}, "spotify": {}, "|": {}, ":": {}, "прослушка!": {}, "cкачать": {},
	"официальный сайт": {}, "apple music": {}, "mediafire": {}, "прослушка": {},
}
