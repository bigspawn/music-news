package main

import (
	botWrap "github.com/bigspawn/music-news/bot"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

const (
	user            = "bigspawn"
	passwd          = "BxYxvmFrohdDwsRjuUjAKXjPRwUvEpmB"
	botID           = "417691036:AAHaptah5zlg5GpRvHkWka_680ZL3_7MKhw"
	chatID    int64 = -1001112604760
	chatIDDev int64 = -1001086871152
	siteURL         = "https://kingdom-leaks.com/index.php?/rss/3-the-kingdom-leaks-homepage-feed.xml"
	address         = "test.bigspawn.com:1080"
)

/**
todo:
	1. правильная обработка ошибок
	2. фильтрация новостей по жанру
	3. если не отправилась картинка - не отправлять текст
	4. ссылка на загрузку файла
	5. кеширование отправленных сообщений для последующей фильтрации
	6. вынести константы в переменные окружения
*/

type News struct {
	ID   int
	Text string
	Link string
}

var excludeWords = []string{"Leaked\n"}
var excludeLastWords = []string{"Download\n", "Downloads\n"}
var excludeGenders = []string{"pop", "rap", "folk", "synthpop", "r&b", "thrash metal", "J-Core", "R&amp;B"}

func containText(text string, words []string) bool {
	for _, word := range words {
		if strings.Contains(strings.ToLower(text), strings.ToLower(word)) {
			return true
		}
	}
	return false
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("[ERROR] Error %s", err)
	}
}

func main() {
	bot, err := botWrap.Create(user, passwd, address, botID)
	handleError(err)
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(siteURL)
	handleError(err)

	for _, item := range feed.Items {
		if item != nil {
			log.Printf("[INFO] %s : %s", item.Title, item.Link)

			regExp, err := regexp.Compile("\n{2,}|\\s{2,}")
			if err != nil {
				log.Printf("[ERROR] Regexp error: %s", err)
				continue
			}

			description := regExp.ReplaceAllString(item.Description, "\n")
			response, err := http.Get(item.Link)
			handleError(err)

			document, err := goquery.NewDocumentFromReader(response.Body)
			handleError(err)

			description, err = url.QueryUnescape(description)
			handleError(err)

			if containText(description, excludeGenders) {
				log.Printf("[DEBUG] Exclude item %s : %s", item.Title, item.Link)
				continue
			}

			imageLink := document.Find("img.ipsImage").First()
			if imageLink != nil {
				link, exist := imageLink.Attr("src")
				if exist {
					if !botWrap.SendImage(chatIDDev, link, bot) {
						continue
					}
				}
			}

			for _, word := range excludeLastWords {
				i := strings.LastIndex(description, word)
				if i > 0 {
					description = description[:i]
					break
				}
			}

			message := item.Title + "\n" + description
			botWrap.SendNews(chatIDDev, message, item, bot)
			log.Printf("[INFO] Item was send %s", message)
		}
	}

}
