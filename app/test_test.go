package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PuerkitoBio/goquery"
)

func Test_One(t *testing.T) {

	//url := "https://kingdom-leaks.com/index.php?/forums/topic/34011-nordic-giants-amplify-human-vibration-2017/&tab=comments#comment-212291"
	url := "https://kingdom-leaks.com/index.php?/forums/topic/33986-chuggaboom-bohemian-rhapsody-single-2020/&tab=comments#comment-212139"
	response, err := http.Get(url)
	assert.NoError(t, err)

	document, err := goquery.NewDocumentFromReader(response.Body)
	assert.NoError(t, err)

	s := "img.ipsImage"
	document.Find(s).
		Not("img[src~='https://kingdom-leaks.com/img/lastfm-logo.png']").
		Each(
			func(i int, selection *goquery.Selection) {
				if v, ok := selection.Attr("src"); ok {
					t.Logf("-> [%s]", v)
				}
			},
		)
}
