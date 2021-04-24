package internal

import (
	"context"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func Test_m4ngParser_Parse(t *testing.T) {
	t.Skip()

	feed, err := gofeed.NewParser().Parse(strings.NewReader(m4ngRss))
	require.NoError(t, err)

	p := NewMusic4newgen(lgr.New(), http.DefaultClient)
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

const m4ngRss = `
<?xml version="1.0" encoding="utf-8"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:media="http://search.yahoo.com/mrss/" xmlns:turbo="http://turbo.yandex.ru" version="2.0">
<channel>
<title>Music4newgen (M4NG) - All Free New Music</title>
<link>https://music4newgen.org/</link>
<language>en</language>
<description>Music4newgen (M4NG) - All Free New Music</description>
<generator>DataLife Engine</generator><item>
<title>Madvillain - Madvillainy (2004)</title>
<guid isPermaLink="true">https://music4newgen.org/8244-madvillain-madvillainy-2004.html</guid>
<link>https://music4newgen.org/8244-madvillain-madvillainy-2004.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/335db41564cb33510b3ae0677892cf38.jpg" style="max-width:100%;" alt="Madvillain - Madvillainy (2004)"><br><b><span style="color:#51ac00">Artist:</span></b> Madvillain<br><b><span style="color:#51ac00">Album:</span></b> Madvillainy<br><b><span style="color:#51ac00">Release:</span></b> 2004<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Hip-Hop<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 117 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>1. The Illest Villains<br>2. Accordion<br>3. Meat Grinder<br>4. Bistro<br>5. Raid<br>6. America's Most Blunted<br>7. Sickfit<br>8. Rainbows<br>9. Curls<br>10. Do Not Fire!<br>11. Money Folder<br>12. Shadows of Tomorrow<br>13. Operation Lifesaver aka Mint Test<br>14. Figaro<br>15. Hardcore Hustle<br>16. Strange Ways<br>17. Fancy Clown<br>18. Eye<br>19. Supervillain Theme<br>20. All Caps<br>21. Great Day<br>22. Rhinestone Cowboy<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album Madvillain - Madvillainy (2004) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci8zMjM4ODczQkQ0Lmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Rap / Hip-Hop]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Sat, 02 Jan 2021 09:58:52 +0100</pubDate>
</item><item>
<title>The Black Keys - Brothers (Deluxe Remastered Anniversary Edition) (2020)</title>
<guid isPermaLink="true">https://music4newgen.org/8243-the-black-keys-brothers-deluxe-remastered-anniversary-edition-2020.html</guid>
<link>https://music4newgen.org/8243-the-black-keys-brothers-deluxe-remastered-anniversary-edition-2020.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/509e0e8b60ef5e98e7d8b76643e686ce.jpg" style="max-width:100%;" alt="The Black Keys - Brothers (Deluxe Remastered Anniversary Edition) (2020)"><br><b><span style="color:#51ac00">Artist:</span></b> The Black Keys<br><b><span style="color:#51ac00">Album:</span></b> Brothers (Deluxe Remastered Anniversary Edition)<br><b><span style="color:#51ac00">Release:</span></b> 2020<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Alternative / Blues Rock<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 158 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. Everlasting Light<br>02. Next Girl<br>03. Tighten Up<br>04. Howlin' for You<br>05. She's Long Gone<br>06. Black Mud<br>07. The Only One<br>08. Too Afraid to Love You<br>09. Ten Cent Pistol<br>10. Sinister Kid<br>11. The Go Getter<br>12. I'm Not the One<br>13. Unknown Brother<br>14. Never Gonna Give You Up<br>15. These Days<br>16. Chop and Change<br>17. Keep My Name Outta Your Mouth<br>18. Black Mud Part II<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album The Black Keys - Brothers (Deluxe Remastered Anniversary Edition) (2020) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci9FQUNDMEU3N0M0Lmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Alternative / Rock]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 09:07:48 +0100</pubDate>
</item><item>
<title>Jean-Michel Jarre - Welcome To The Other Side (Concert From Virtual Notre-Dame) (2020)</title>
<guid isPermaLink="true">https://music4newgen.org/8242-jean-michel-jarre-welcome-to-the-other-side-concert-from-virtual-notre-dame-2020.html</guid>
<link>https://music4newgen.org/8242-jean-michel-jarre-welcome-to-the-other-side-concert-from-virtual-notre-dame-2020.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/b57f5010bc3cd8933181bab798d3c468.jpg" style="max-width:100%;" alt="Jean-Michel Jarre - Welcome To The Other Side (Concert From Virtual Notre-Dame) (2020)"><br><b><span style="color:#51ac00">Artist:</span></b> Jean-Michel Jarre<br><b><span style="color:#51ac00">Album:</span></b> Welcome To The Other Side (Concert From Virtual Notre-Dame)<br><b><span style="color:#51ac00">Release:</span></b> 2020<br><b><span style="color:#51ac00">Country:</span></b> France<br><b><span style="color:#51ac00">Genre:</span></b> Electronic<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 112 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. The Opening (VR Live)<br>02. Oxygene 2 (JMJ Rework of Kosinski Remix)<br>03. The Architect (VR Live)<br>04. Oxygene 19 (VR Live)<br>05. Oxygene 8 (VR Live)<br>06. Jean-michel Jarre &amp; Tangerine Dream - Zero Gravity (VR Live)<br>07. Exit (VR Live)<br>08. Equinoxe 4 (VR Live)<br>09. Stardust (VR Live)<br>10. Herbalizer (VR Live)<br>11. Oxygene 4 (JMJ Rework of Astral Projection Remix)<br>12. Jean-michel Jarre &amp; Boys Noize - The Time Machine (VR Live)<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album Jean-Michel Jarre - Welcome To The Other Side (Concert From Virtual Notre-Dame) (2020) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci9DNUZGQjA4MTRDLmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Electronic / Dance / House]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 09:05:45 +0100</pubDate>
</item><item>
<title>Uncle Murda - Don&#039;t Come Outside, Vol. 3 (2021)</title>
<guid isPermaLink="true">https://music4newgen.org/8241-uncle-murda-dont-come-outside-vol-3-2021.html</guid>
<link>https://music4newgen.org/8241-uncle-murda-dont-come-outside-vol-3-2021.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/d6124dd316577556d5ca3cce588ffda7.jpg" style="max-width:100%;" alt="Uncle Murda - Don't Come Outside, Vol. 3 (2021)"><br><b><span style="color:#51ac00">Artist:</span></b> Uncle Murda<br><b><span style="color:#51ac00">Album:</span></b> Don't Come Outside, Vol. 3<br><b><span style="color:#51ac00">Release:</span></b> 2021<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Hip-Hop<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 85 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. Intro<br>02. Change Gone Come (feat. Mysonne, Tamika Mallory)<br>03. They Said (feat. Lil Tjay, Que Banz)<br>04. Bro Shit<br>05. Life's A Bitch<br>06. Party Full Of Demons (feat. Que Banz)<br>07. Russian Roulette<br>08. Whole Lotta Money (feat. Benny The Butcher, Que Banz)<br>09. Nothing Like Me (feat. Conway the Machine, Dios Moreno)<br>10. Part Of The Plan (feat. Jase)<br>11. Down Bad<br>12. Montana (feat. Rich Starz)<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album Uncle Murda - Don't Come Outside, Vol. 3 (2021) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci9GNkEwMUMwOTY0Lmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Rap / Hip-Hop]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 08:38:13 +0100</pubDate>
</item><item>
<title>Yung Pinch - WASHED ASHORE (2021)</title>
<guid isPermaLink="true">https://music4newgen.org/8240-yung-pinch-washed-ashore-2021.html</guid>
<link>https://music4newgen.org/8240-yung-pinch-washed-ashore-2021.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/a143b48f110f6a941c04ff9f92310e30.jpg" style="max-width:100%;" alt="Yung Pinch - WASHED ASHORE (2021)"><br><b><span style="color:#51ac00">Artist:</span></b> Yung Pinch<br><b><span style="color:#51ac00">Album:</span></b> WASHED ASHORE<br><b><span style="color:#51ac00">Release:</span></b> 2021<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Hip-Hop<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 77 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. NO ONES THERE<br>02. MIXED MESSAGES<br>03. CLOSE FRIENDS<br>04. FAR CRY<br>05. HARD TIMES<br>06. DEVILS WAY<br>07. HEART SHAPED NECKLACE<br>08. I DON’T EVEN CARE<br>09. PRECIOUS CARGO<br>10. OTHERSIDE<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album Yung Pinch - WASHED ASHORE (2021) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci8wNDk0QkNDOTNBLmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Rap / Hip-Hop]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 08:36:58 +0100</pubDate>
</item><item>
<title>Westside Tut - Don&#039;t Let Go 2 (Deluxe) (2021)</title>
<guid isPermaLink="true">https://music4newgen.org/8239-westside-tut-dont-let-go-2-deluxe-2021.html</guid>
<link>https://music4newgen.org/8239-westside-tut-dont-let-go-2-deluxe-2021.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/469c0e1c134956f8dffd54479acf9480.jpg" style="max-width:100%;" alt="Westside Tut - Don't Let Go 2 (Deluxe) (2021)"><br><b><span style="color:#51ac00">Artist:</span></b> Westside Tut<br><b><span style="color:#51ac00">Album:</span></b> Don't Let Go 2 (Deluxe)<br><b><span style="color:#51ac00">Release:</span></b> 2021<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Hip-Hop<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 106 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. Heartbroken Heartbreaker<br>02. On My Line<br>03. Get Back (feat. Youngboy Never Broke Again)<br>04. Stressed Out<br>05. Big Boy Money (feat. Tee Grizzley)<br>06. Last Night<br>07. On the Opps (feat. NLE Choppa)<br>08. I Don't Think You Have (feat. Bla$ta)<br>09. Hot Shit (feat. Bdm Drewski)<br>10. Feelings in a Jar<br>11. Play Cousin (feat. Sada Baby)<br>12. 43 &amp; 0 (feat. Tee Grizzley)<br>13. No Disses<br>14. On Me (feat. Lil Tjay)<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album Westside Tut - Don't Let Go 2 (Deluxe) (2021) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci9CNTVDRTVDNzc4Lmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Rap / Hip-Hop]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 08:35:20 +0100</pubDate>
</item><item>
<title>TheHxliday - Batbxy (2021)</title>
<guid isPermaLink="true">https://music4newgen.org/8238-thehxliday-batbxy-2021.html</guid>
<link>https://music4newgen.org/8238-thehxliday-batbxy-2021.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/ab5c87a6a4f34d730acbdf41bc7fb020.jpg" style="max-width:100%;" alt="TheHxliday - Batbxy (2021)"><br><b><span style="color:#51ac00">Artist:</span></b> TheHxliday<br><b><span style="color:#51ac00">Album:</span></b> Batbxy<br><b><span style="color:#51ac00">Release:</span></b> 2021<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Hip-Hop<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 43 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. NxBody<br>02. Laugh A Little<br>03. Thank U<br>04. LxneChild<br>05. Batgirl<br>06. Bad<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album TheHxliday - Batbxy (2021) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci83RjFCRTQ5NEFDLmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Rap / Hip-Hop]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 08:33:44 +0100</pubDate>
</item><item>
<title>The Dirty Nil - Fuck Art (2021)</title>
<guid isPermaLink="true">https://music4newgen.org/8237-the-dirty-nil-fuck-art-2021.html</guid>
<link>https://music4newgen.org/8237-the-dirty-nil-fuck-art-2021.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/edc5344597112e973e3c5224933585cd.jpg" style="max-width:100%;" alt="The Dirty Nil - Fuck Art (2021)"><br><b><span style="color:#51ac00">Artist:</span></b> The Dirty Nil<br><b><span style="color:#51ac00">Album:</span></b> Fuck Art<br><b><span style="color:#51ac00">Release:</span></b> 2021<br><b><span style="color:#51ac00">Country:</span></b> Canada<br><b><span style="color:#51ac00">Genre:</span></b> Alt. Rock<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 88 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. Doom Boy<br>02. Blunt Force Concussion<br>03. Elvis '77<br>04. Done with Drugs<br>05. Ride or Die<br>06. Hang Yer Moon<br>07. Damage Control<br>08. Hello Jealousy<br>09. Possession<br>10. To the Guy Who Stole My Bike<br>11. One More and the Bill<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album The Dirty Nil - Fuck Art (2021) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci9DM0EwOUJCRDA4Lmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Alternative / Rock]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 08:32:20 +0100</pubDate>
</item><item>
<title>R.A.P. Ferreira - Bob&#039;s Son (2021)</title>
<guid isPermaLink="true">https://music4newgen.org/8236-rap-ferreira-bobs-son-2021.html</guid>
<link>https://music4newgen.org/8236-rap-ferreira-bobs-son-2021.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/1ec3b976e47f8025a16aebaeeac86f61.jpg" style="max-width:100%;" alt="R.A.P. Ferreira - Bob's Son (2021)"><br><b><span style="color:#51ac00">Artist:</span></b> R.A.P. Ferreira &amp; Scallops Hotel<br><b><span style="color:#51ac00">Album:</span></b> Bob's Son<br><b><span style="color:#51ac00">Release:</span></b> 2021<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Hip-Hop<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 86 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. battle report (feat. Pink Navel)<br>02. the cough bomber's return<br>03. yamships, flaxseed<br>04. diogenes on the auction block<br>05. redguard snipers (feat. SB the Moor)<br>06. sips of ripple wine<br>07. skrenth<br>08. bobby digital's little wings<br>09. listening<br>10. high rise in newark<br>11. rejoice (feat. Eldon Somers)<br>12. abomunist manifesto<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album R.A.P. Ferreira - Bob's Son (2021) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly93d3cuZmlsZWNyeXB0LmNjL0NvbnRhaW5lci9GMkU1OEY0MTMwLmh0bWw%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Rap / Hip-Hop]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 08:30:51 +0100</pubDate>
</item><item>
<title>Key! - The Alpha Jerk (2021)</title>
<guid isPermaLink="true">https://music4newgen.org/8235-key-the-alpha-jerk-2021.html</guid>
<link>https://music4newgen.org/8235-key-the-alpha-jerk-2021.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://music4newgen.org/uploads/posts/2021-01/1609486166_500x500-000000-80-0-0.jpg" style="max-width:100%;" alt="Key! - The Alpha Jerk (2021)"><br><b><span style="color:#51ac00">Artist:</span></b> Key!<br><b><span style="color:#51ac00">Album:</span></b> The Alpha Jerk<br><b><span style="color:#51ac00">Release:</span></b> 2021<br><b><span style="color:#51ac00">Country:</span></b> USA<br><b><span style="color:#51ac00">Genre:</span></b> Hip-Hop<br><b><span style="color:#51ac00">Quality:</span></b> mp3, 320 kbps<br><b><span style="color:#51ac00">Size:</span></b> 99 MB<br><br><b><span style="color:#51ac00">Tracklist:</span></b><br>01. Like That<br>02. Clinical<br>03. My Puppy<br>04. Change<br>05. Xylophone<br>06. No Sirski (feat. Lil Yachty)<br>07. Voodoo<br>08. Wylin'<br>09. Fashion Week (feat. Sonny Digital)<br>10. Rida Rida<br>11. Destruction<br>12. I Know<br>13. Heart (feat. 448rasta)<br>14. Narcissistic<br>15. Ima Star (feat. Quadie Diesel)<br>16. Scott Disick<br>17. Send It<br><br><span style="font-size:12pt;"><span style="color:#51ac00"><b>Download Album Key! - The Alpha Jerk (2021) Free</b></span></span><br>----------------------------------------------------------<br>-- <a href="https://music4newgen.org/index.php?do=go&amp;url=aHR0cHM6Ly9maWxlY3J5cHQuY2MvQ29udGFpbmVyL0Y1NEMwN0E0MUQuaHRtbA%3D%3D" target="_blank"><span style="font-family:Georgia"><span style="font-size:14pt;">DOWNLOAD</span></span></a> --<br>----------------------------------------------------------</div>]]></description>
<category><![CDATA[Rap / Hip-Hop]]></category>
<dc:creator>admin2</dc:creator>
<pubDate>Fri, 01 Jan 2021 08:28:53 +0100</pubDate>
</item></channel></rss>
`
