package internal

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
	"strings"
)

const AlterportalRSSFeedURL = "https://alterportal.net/rss.xml"

var newLinesRE = regexp.MustCompile("\n{2,}")

func NewAlterportalParser(lgr lgr.L, client *http.Client) ItemParser {
	return &alterportalParser{
		lgr:    lgr,
		client: client,
	}
}

type alterportalParser struct {
	lgr    lgr.L
	client *http.Client
}

func (p alterportalParser) Parse(ctx context.Context, item *gofeed.Item) (*News, error) {
	if strings.Contains(item.Link, "raznoe") ||
		strings.Contains(item.Link, "video") ||
		strings.Contains(item.Link, "neformat") {
		return nil, errSkipItem
	}

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

	title := doc.Find("title").Text()
	if strings.Contains(title, "502: Bad gateway") {
		return nil, errSkipItem
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
		return nil, errSkipItem
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

	news.Text = translate(news.Text)
	news.Text = newLinesRE.ReplaceAllString(news.Text, "\n")
	news.Text = trimLast(news.Text)

	return news, nil
}

func trimLast(text string) string {
	re := regexp.MustCompile("^\\d{1,3}\\.")
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
		if isSkippedWord(data) {
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
}

func isSkippedWord(data string) bool {
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

func isSkippedGender(data string) bool {
	for _, s := range skipGenders {
		if strings.Contains(data, s) || strings.Contains(data, s) {
			return true
		}
	}
	return false
}

var skipGenders = []string{
	"Retro Pop", "R&B", "Pop Music", "Pop Rock", "City Pop", "Disco", "Eurodance", "Hip-Hop",
	"retro pop", "r&b", "pop music", "pop rock", "city pop", "disco", "eurodance", "hip-hop",
	"Pop", "pop", "Rap / Hip Hop", "R'n'B",
}

func isAllowedFileHost(host string) bool {
	for _, s := range fileHosts {
		if strings.Contains(host, s) {
			return true
		}
	}
	return false
}

var fileHosts = []string{
	"mediafire.com",
	"zippyshare.com",
	"mega.nz",
	"solidfiles.com",
	"drive.google.com",
	"files.mail.ru",
	"disk.yandex.ru",
	"yadi.sk",
	"files.fm",
	"uppit.com",
	"filecrypt.cc",
	"turb.cc",
	"turbobit.net",
}
