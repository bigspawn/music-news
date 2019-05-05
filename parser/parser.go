package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

const divSelector = "div.ipsType_normal.ipsType_richText.ipsContained[data-role=commentContent][data-controller=core.front.core.lightboxedImages]"

var excludeWords = []string{"Leaked\n"}
var excludeLastWords = []string{"Download\n", "Downloads\n", "Total length:"}
var excludeGenders = []string{"pop", "rap", "folk", "synthpop", "r&b", "thrash metal", "J-Core", "R&amp;B"}

type News struct {
	ID           int
	Text         string
	ImageLink    string
	DownloadLink []string
	PageLink     string
}

func Parse(feedUrl string) ([]News, error) {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(feedUrl)
	if err != nil {
		log.Fatalf("[ERROR] Parsing site : %s", err)
	}
	var news []News
	for _, item := range feed.Items {
		if item != nil {
			log.Printf("[INFO] %s : %s", item.Title, item.Link)
			var n News
			regExp, err := regexp.Compile("\n{2,}|\\s{2,}")
			if err != nil {
				log.Printf("[ERROR] Regexp error: %s", err)
				continue
			}
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
			description, err := url.QueryUnescape(regExp.ReplaceAllString(item.Description, "\n"))
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}
			if containText(description, excludeGenders) {
				log.Printf("[DEBUG] Exclude item %s : %s", item.Title, item.Link)
				continue
			}

			res, err := http.Get(item.Link)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}
			defer res.Body.Close()

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
				continue
			}

			articleDiv := doc.Find("div.ipsType_normal.ipsType_richText.ipsContained").First()
			link := articleDiv.Find("a[rel~='external']").First()
			if val, exists := link.Attr("href"); exists {
				log.Printf("[INFO] link %s", val)
				n.DownloadLink = append(n.DownloadLink, val)
			} else {
				continue
			}

			//if downloadLink, exists := first.Attr("href"); exists {
			//	n.DownloadLink = append(n.DownloadLink, downloadLink)
			//} else {
			//	continue
			//}

			imageLink := document.Find("img.ipsImage").First()
			if link, exist := imageLink.Attr("src"); exist {
				n.ImageLink = link
				n.Text = item.Title + "\n" + normalize(description)
				n.PageLink = item.Link
			} else {
				continue
			}
			news = append(news, n)
		}
	}
	return news, nil
}

func normalize(description string) string {
	for _, word := range excludeLastWords {
		index := strings.LastIndex(description, word)
		if index > 0 {
			description = description[:index]
		}
	}
	for _, word := range excludeWords {
		index := strings.Index(description, word)
		if index > 0 {
			description = description[index : index+utf8.RuneCountInString(word)]
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
