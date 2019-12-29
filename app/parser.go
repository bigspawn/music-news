package main

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

var ExcludeWords = []string{"Leaked\n"}
var ExcludeLastWords = []string{"Download\n", "Downloads\n", "Total length:"}
var ExcludeGenders = []string{"pop", "rap", "folk", "synthpop", "r&b", "thrash metal", "J-Core", "R&amp;B"}

type SiteParser struct {
	Exclude struct {
		Words     []string
		LastWords []string
		Genders   []string
	}
	FeedParser *gofeed.Parser
	Store      *NewsStore
	URL        string
}

func (p *SiteParser) Parse() ([]*News, error) {
	feed, err := p.FeedParser.ParseURL(p.URL)
	if err != nil {
		return nil, err
	}
	news := make([]*News, 10)
	for _, item := range feed.Items {
		if item != nil {
			exist, err := p.Store.Exist(item.Title)
			if err != nil {
				return nil, err
			}
			if exist {
				continue
			}

			log.Printf("[INFO] Parse news [%s: %s]", item.Title, item.Link)

			response, err := http.Get(item.Link)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}
			document, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}
			re := regexp.MustCompile("[\\n\\t\\s]{2,}")
			desc := re.ReplaceAllString(item.Description, "\n")
			description, err := url.QueryUnescape(desc)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}
			if containText(description, p.Exclude.Genders) {
				log.Printf("[DEBUG] Exclude item [%s: %s]", item.Title, item.Link)
				continue
			}

			res, err := http.Get(item.Link)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}
			_ = res.Body.Close()

			n := &News{}
			n.DateTime = item.PublishedParsed

			articleDiv := doc.Find("div.ipsType_normal.ipsType_richText.ipsContained").First()
			link := articleDiv.Find("a[rel~='external']").First()
			if val, exists := link.Attr("href"); exists {
				log.Printf("[INFO] Download link [%s]", val)
				n.DownloadLink = append(n.DownloadLink, val)
			} else {
				continue
			}

			imageLink := document.Find("img.ipsImage").First()
			if link, exist := imageLink.Attr("src"); exist {
				n.ImageLink = link
				n.Text = p.normalize(description)
				n.PageLink = item.Link
				n.Title = item.Title
			} else {
				continue
			}

			err = p.Store.Insert(n)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}
			news = append(news, n)
		}
	}
	return news, nil
}

func (p *SiteParser) normalize(description string) string {
	for _, word := range p.Exclude.LastWords {
		index := strings.LastIndex(description, word)
		if index > 0 {
			description = description[:index]
		}
	}
	for _, word := range p.Exclude.Words {
		index := strings.Index(description, word)
		if index > -1 {
			description = description[:index] + description[index+utf8.RuneCountInString(word):]
		}
	}
	return description
}

func containText(text string, words []string) bool {
	for _, word := range words {
		if strings.Contains(strings.ToLower(text), strings.ToLower(word)) {
			return true
		}
	}
	return false
}
