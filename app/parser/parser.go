package parser

import (
	"database/sql"
	"github.com/PuerkitoBio/goquery"
	"github.com/bigspawn/music-news/db"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var excludeWords = []string{"Leaked\n"}
var excludeLastWords = []string{"Download\n", "Downloads\n", "Total length:"}
var excludeGenders = []string{"pop", "rap", "folk", "synthpop", "r&b", "thrash metal", "J-Core", "R&amp;B"}

func Parse(feedUrl string, connection *sql.DB) ([]db.News, error) {
	var news []db.News
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(feedUrl)
	if err != nil {
		return news, err
	}
	for _, item := range feed.Items {
		if item != nil {
			row, err := db.Select(connection, item.Title)
			if err != nil {
				log.Printf("[ERROR] Error %v", err)
				continue
			}
			if row != nil && row.Next() {
				log.Printf("[INFO] Skip news [%s], it contains in DB", item.Title)
				continue
			}

			log.Printf("[INFO] Parse news [%s: %s]", item.Title, item.Link)
			regExp, err := regexp.Compile("\\n{2,}|\\s{2,}")
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
				log.Printf("[DEBUG] Exclude item [%s: %s]", item.Title, item.Link)
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

			var n db.News
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
				n.Text = normalize(description)
				n.PageLink = item.Link
				n.Title = item.Title
			} else {
				continue
			}

			err = db.Insert(connection, n)
			if err != nil {
				log.Printf("[ERROR] Error %s", err)
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
