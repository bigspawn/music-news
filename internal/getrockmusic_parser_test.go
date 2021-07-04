package internal

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/require"
)

func Test_getRockMusicParser_Parse(t *testing.T) {
	t.Skip()

	feed, err := gofeed.NewParser().Parse(strings.NewReader(getRockMusicParserRss))
	require.NoError(t, err)

	p := NewGetRockMusicParser(lgr.New(), http.DefaultClient)
	ctx := context.Background()
	for _, i := range feed.Items {
		result, err := p.Parse(ctx, i)
		if err == ErrSkipItem {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("result %v\n", result)
	}
}

const getRockMusicParserRss = `
<?xml version="1.0" encoding="utf-8"?>
<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/">
    <channel>
        <title>GetRockMusic</title>
        <link>https://getrockmusic.net/</link>
        <language>ru</language>
        <description>GetRockMusic</description>
        <generator>DataLife Engine</generator>
        <item>
            <title>Break of Dusk - …and the Winter Came</title>
            <guid isPermaLink="true">https://getrockmusic.net/59292-break-of-dusk-and-the-winter-came.html</guid>
            <link>https://getrockmusic.net/59292-break-of-dusk-and-the-winter-came.html</link>
            <description>
                <![CDATA[<center><img title="" src="https://getrockmusic.net/uploads/posts/2021-01/1610016949_cover.jpg" alt="" /><br />Artist: <a href="https://getrockmusic.net/artist/break-of-dusk" title="Break of Dusk">Break of Dusk</a><br /> Album: …and the Winter Came<br /> Genre: <a href="https://getrockmusic.net/melodic-death-metal/">Melodic Death Metal</a><br /> Country: Finland<br /> Released: 2021</center>]]>
            </description>
            <category>
                <![CDATA[Melodic Death Metal]]>
            </category>
            <dc:creator>GetRock</dc:creator>
            <pubDate>Thu, 07 Jan 2021 12:55:39 +0200</pubDate>
        </item>
    </channel>
</rss>
`
