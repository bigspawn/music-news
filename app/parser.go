package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

var (
	ExcludeWords     = []string{"Leaked\n"}
	ExcludeLastWords = []string{"Download\n", "Downloads\n", "Total length:"}
	ExcludeGenders   = []string{"pop", "rap", "folk", "synthpop", "r&b", "thrash metal", "J-Core", "R&amp;B"}

	imageCssSelectorPath      = "html.wf-roboto-n3-active.wf-roboto-n4-active.wf-roboto-i4-active.wf-roboto-n7-active.wf-roboto-i3-active.wf-roboto-i7-active.wf-active body.ipsApp.ipsApp_front.ipsJS_has.ipsClearfix.ipsApp_noTouch main#ipsLayout_body.ipsLayout_container.v-nav-wrap div#ipsLayout_contentArea div#ipsLayout_contentWrapper div#ipsLayout_mainArea div.cTopic.ipsClear.ipsSpacer_top div.ipsAreaBackground_light form article#elComment_212133.cPost.ipsBox.ipsComment.ipsComment_parent.ipsClearfix.ipsClear.ipsColumns.ipsColumns_noSpacing.ipsColumns_collapsePhone div.ipsColumn.ipsColumn_fluid div#comment-212133_wrap.ipsComment_content.ipsType_medium.ipsFaded_withHover div.cPost_contentWrap.ipsPad div.ipsType_normal.ipsType_richText.ipsContained p a img.ipsImage"
	imageCssSelectorPathThumb = imageCssSelectorPath + ".ipsImage_thumbnailed"
	imgSelector               = "img.ipsImage.ipsImage_thumbnailed"
	imgSelector2              = "img.ipsImage"
)

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
	var news []*News
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

			n.PageLink = item.Link
			n.Title = item.Title

			n.ImageLink, err = extractImage(document)
			if err != nil {
				log.Printf("[ERROR] Can't find image link for %v", n)
				continue
			}

			n.Text = p.normalize(description)

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

func extractImage(document *goquery.Document) (string, error) {
	imageLink := document.
		Find("img.ipsImage").
		Not("img[src~='https://kingdom-leaks.com/img/lastfm-logo.png']").
		First()

	if link, ok := imageLink.Attr("src"); ok && link != "" {
		return link, nil
	}

	return "", errors.New("image link not found")
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
