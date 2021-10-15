package internal

import (
	"context"
	"strings"
	"testing"

	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/require"
)

func TestGetMetalParser_Parse(t *testing.T) {
	t.Skip()

	feed, err := gofeed.NewParser().Parse(strings.NewReader(alterportalRss))
	require.NoError(t, err)

	// feed, err := gofeed.NewParser().ParseURL(rssFeed)
	// require.NoError(t, err)

	p := NewAlterportalParser(lgr.New())
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

const alterportalRss = `
<?xml version="1.0" encoding="utf-8"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:media="http://search.yahoo.com/mrss/" xmlns:turbo="http://turbo.yandex.ru" version="2.0">
<channel>
<title>Alterportal.net V2</title>
<link>https://alterportal.net/</link>
<language>ru</language>
<description>Alterportal.net V2</description>
<generator>DataLife Engine</generator><item turbo="true">
<title>Mechina - Siege (2021)</title>
<guid isPermaLink="true">https://alterportal.net/2021_albums/170465-mechina-siege-2021.html</guid>
<link>https://alterportal.net/2021_albums/170465-mechina-siege-2021.html</link>
<description><![CDATA[<img src="https://funkyimg.com/i/39LZs.jpg" style="max-width:100%;" alt="Mechina - Siege (2021)"><br><span style="font-family:Georgia"><b>Стиль: <span style="color:#FF6600">Modern Death Metal / Electronic / Atmospheric</span><br>Треклист:</b><br>01. King Breeder<br>02. The Worst in Us<br>03. Shock Doctrine<br>04. Purity Storm<br>05. Siege<br>06. Claw at the Dirt<br>07. Blood Feud Erotica<br>08. Freedom Foregone<br><br><b>CD2: Siege: Instrumental<br>CD3: Siege: Core</b></span>]]></description>
<turbo:content><![CDATA[ <div style="background:#000000;"><br><div><a href="https://funkyimg.com/i/39LZs.jpg" target="_blank" rel="noopener external noreferrer"><img src="https://funkyimg.com/i/39LZs.jpg" style="max-width:100%;" alt="Mechina - Siege (2021)"></a><br><span style="font-family:Georgia"><b><span style="font-size:10pt;"><a href="https://www.facebook.com/pages/Mechina/114376106370?id=114376106370&amp;sk=info" target="_blank" rel="noopener external noreferrer"><span style="color:#33CC00">Facebook</span></a></span><br><br>Родина: <span style="color:#FF6600"> Chicago, USA</span><br>Стиль: <span style="color:#FF6600">Modern Death Metal / Electronic / Atmospheric</span><br>Формат: <span style="color:#FF6600">FLAC (tracks) / MP3 CBR 320 kbps</span><br>Лейбл: <span style="color:#FF6600"> -</span><br>Дата релиза: <span style="color:#FF6600">01 Января 2021</span></b><br><br><b>Треклист:</b><br>01. King Breeder<br>02. The Worst in Us<br>03. Shock Doctrine<br>04. Purity Storm<br>05. Siege<br>06. Claw at the Dirt<br>07. Blood Feud Erotica<br>08. Freedom Foregone<br><br>CD2: Siege: Instrumental<br>CD3: Siege: Core<br><br><div style="text-align:center;"><a href="https://www34.zippyshare.com/v/VUnLfQ9D/file.html" target="_blank" rel="noopener external noreferrer"><b><span style="font-size:12pt;"><span style="color:#FF0000">Siege</span></span></b></a><span style="color:#000000"> | </span><a href="https://www34.zippyshare.com/v/d6sxqQbW/file.html" target="_blank" rel="noopener external noreferrer"><b><span style="font-size:12pt;"><span style="color:#FF0000">Core</span></span></b></a><span style="color:#000000"> | </span><a href="https://www34.zippyshare.com/v/8oGI3qqX/file.html" target="_blank" rel="noopener external noreferrer"><b><span style="font-size:12pt;"><span style="color:#FF0000">Instrumental</span></span></b></a></div><br><br><div style="text-align:center;"><iframe style="border:0;width:400px;height:406px;" src="https://bandcamp.com/EmbeddedPlayer/album=1152799763/size=large/bgcol=333333/linkcol=ffffff/artwork=small/transparent=true/"><a href="https://mechinamusic.bandcamp.com/album/siege" target="_blank" rel="noopener external noreferrer">Siege by Mechina</a></iframe></div><br><div style="text-align:center;"><iframe style="border:0;width:400px;height:373px;" src="https://bandcamp.com/EmbeddedPlayer/album=4211283451/size=large/bgcol=333333/linkcol=ffffff/artwork=small/transparent=true/"><a href="https://mechinamusic.bandcamp.com/album/siege-core" target="_blank" rel="noopener external noreferrer">Siege [Core] by Mechina</a></iframe></div><br><div style="text-align:center;"><iframe style="border:0;width:400px;height:406px;" src="https://bandcamp.com/EmbeddedPlayer/album=523366626/size=large/bgcol=333333/linkcol=ffffff/artwork=small/transparent=true/"><a href="https://mechinamusic.bandcamp.com/album/siege-instrumental" target="_blank" rel="noopener external noreferrer">Siege [Instrumental] by Mechina</a></iframe></div></div><br></div> ]]></turbo:content>
<category><![CDATA[Альбомы 2021 | Metal]]></category>
<dc:creator>Смок</dc:creator>
<pubDate>Sat, 02 Jan 2021 00:28:56 +0300</pubDate>
</item><item turbo="true">
<title>Prozak</title>
<guid isPermaLink="true">https://alterportal.net/neformat/170464-prozak.html</guid>
<link>https://alterportal.net/neformat/170464-prozak.html</link>
<description><![CDATA[<div style="text-shadow:1px 1px #000000;"><div style="text-align:center;"><img src="https://i114.fastpic.ru/big/2021/0101/2a/d3e31d0473c18f8f999962b499b5812a.png" style="border:none;" alt="Society 1" title="Society 1"><br><br><b>Стиль:</b> <span style="color:#33FF33">Horrorcore | Hardcore Rap | Rap Metal | Rapcore</span><br><b>Формат:</b> <span style="color:#33FF33">M4A</span><br><br><b>Список Альбомов:</b><br>2003 - Aftabirth<br>2008 - Tales From The Sick<br>2012 - Nocturnal (EP)<br>2012 - Paranormal<br>2013 - We All Fall Down<br>2015 - Black Ink<br><br><b>Протекция Alternativshik_6, январь 1/3</b></div></div>]]></description>
<turbo:content><![CDATA[ <div><div style="background:url(&quot;http://i46.fastpic.ru/big/2013/0716/21/5bda6404bcce8e8334491f235c29c721.jpg&quot;);border-radius:10px 10px 10px 10px;border:1px solid #666666;"><br><a href="https://fastpic.ru/view/114/2021/0101/17ed23ea43946cf41b08c6b909126eac.jpg.html" target="_blank" rel="noopener external noreferrer"><img src="https://i114.fastpic.ru/big/2021/0101/ac/17ed23ea43946cf41b08c6b909126eac.jpg" style="max-width:100%;" alt="Prozak"></a><br><br><a href="https://www.facebook.com/therealProzak" target="_blank" rel="noopener external noreferrer"><img src="http://i48.fastpic.ru/big/2013/0716/18/1a55c10ba823ca1dd070705807178e18.png" style="border:none;" alt="Society 1" title="Society 1"></a> <a href="https://music.apple.com/ru/artist/prozak/4258290" target="_blank" rel="noopener external noreferrer"><img src="http://i46.fastpic.ru/big/2013/0716/c2/f2e2f7ba5d417f81a9cd3d341d0e51c2.png" style="border:none;" alt="Society 1" title="Society 1"></a><br><br><span style="font-family:Century Gothic"><b><span style="color:#33FF33">Стиль: </span></b> Horrorcore | Hardcore Rap | Rap Metal | Rapcore <b><span style="color:#33FF33"><br>Страна: </span></b> USA, Saginaw <b><span style="color:#33FF33"><br>Год выпуска: </span></b>2003-2021<b><span style="color:#33FF33"><br>Формат: </span></b>M4A</span><br><br><div style="background-color:#000000;width:490px;padding:0px;border-radius:15px 15px 15px 15px;"><br><span style="font-family:Tahoma"></span><a href="https://fastpic.ru/view/114/2021/0101/dc8e5364f9f596f8296d5301d13b751e.jpg.html" target="_blank" rel="noopener external noreferrer"><img src="https://i114.fastpic.ru/big/2021/0101/1e/dc8e5364f9f596f8296d5301d13b751e.jpg" style="max-width:100%;" alt=""></a><br><br><b><span style="color:#33FF33">Качество (Размер):</span> m4a ~256 kbps (61.8 Mb) </b><br><br><b><span style="color:#33FF33">Треклист:</span></b><br><span style="font-family:Georgia"><br>01 - I'm Back<br>02 - Aftabirth<br>03 - Burnt Offeringz (Feat. Cap One)<br>04 - Fallin' Down<br>05 - Wicked 4 Life<br>06 - Facez Of Evil<br>07 - Callin' Me<br>08 - Optionz<br>09 - Aftabirth 2000</span><br><br><br></div><br><div style="background-color:#000000;width:490px;padding:0px;border-radius:15px 15px 15px 15px;"><br><span style="font-family:Tahoma"></span><a href="https://fastpic.ru/view/114/2021/0101/dc949cc89b420c675bb71b9c70690d3f.jpg.html" target="_blank" rel="noopener external noreferrer"><img src="https://i114.fastpic.ru/big/2021/0101/3f/dc949cc89b420c675bb71b9c70690d3f.jpg" style="max-width:100%;" alt=""></a><br><br><b><span style="color:#33FF33">Качество (Размер):</span> m4a ~256 kbps (141 Mb) </b><br><br><b><span style="color:#33FF33">Треклист:</span></b><br><span style="font-family:Georgia"><br>01 - The Hitchcock Of Hip Hop<br>02 - Fun 'N' Games<br>03 - Keep Grindin' (Feat. Krizz Kaliko)<br>04 - Scapegoat<br>05 - Go To Hell<br>06 - Crossin' Over<br>07 - It Was You (Intro)<br>08 - It Was You (Feat. Krizz Kaliko)<br>09 - Why??? (Feat. Tech N9ne &amp; Twista)<br>10 - Run Away (Feat. Tech N9ne &amp; Krizz Kaliko)<br>11 - Corporate Genocide<br>12 - Bombs Away<br>13 - Holy War (Feat. Mike E.Clark)<br>14 - It's Too Late Now<br>15 - Insane (Feat. Insane Clown Posse)<br>16 - Bodies Fall (feat. Blaze Ya Dead Homie, Kutt Calhoun, Tech N9ne &amp; The R.O.C.)<br>17 - Psycho, Psycho, Psycho! (Feat. Bizarre &amp; King Gordy)<br>18 - Drugs<br>19 - Living In The Fog (Feat. Cypress Hill)<br>20 - Fading... (Feat. Twiztid &amp; Krizz Kaliko)<br>21 - Good Enough (Feat. Mike E.Clark)<br>22 - Under The Rain (Feat. Krizz Kaliko)</span><br><br><br></div><br><div style="background-color:#000000;width:490px;padding:0px;border-radius:15px 15px 15px 15px;"><br><span style="font-family:Tahoma"></span><a href="https://fastpic.ru/view/114/2021/0101/_e114334212c863a42c84b9c48ab3da77.jpg.html" target="_blank" rel="noopener external noreferrer"><img src="https://i114.fastpic.ru/thumb/2021/0101/77/_e114334212c863a42c84b9c48ab3da77.jpeg" style="max-width:100%;" alt=""></a><br><br><b><span style="color:#33FF33">Качество (Размер):</span> m4a ~256 kbps (51.9 Mb) </b><br><br><b><span style="color:#33FF33">Треклист:</span></b><br><span style="font-family:Georgia"><br>01. Intro<br>02. Molotov<br>03. Shadow Of Death<br>04. Giving Up<br>05. 2012<br>06. Vigilante (feat. Twiztid)<br>07. Knuckle Up (feat. R.O.C. &amp; Dirtball)<br>08. No More</span><br><br><br></div><br><div style="background-color:#000000;width:490px;padding:0px;border-radius:15px 15px 15px 15px;"><br><span style="font-family:Tahoma"></span><a href="https://fastpic.ru/view/114/2021/0101/_b62275f3cddd6864d9add100544d00b5.jpg.html" target="_blank" rel="noopener external noreferrer"><img src="https://i114.fastpic.ru/thumb/2021/0101/b5/_b62275f3cddd6864d9add100544d00b5.jpeg" style="max-width:100%;" alt=""></a><br><br><b><span style="color:#33FF33">Качество (Размер):</span> m4a ~256 kbps (115 Mb) </b><br><br><b><span style="color:#33FF33">Треклист:</span></b><br><span style="font-family:Georgia"><br>01 - Paranormal<br>02 - Prepare For The Worst<br>03 - The End Of Us (Feat. Sid Wilson)<br>04 - Tell A Tale Of Two Hearts<br>05 - Line In The Middle (Feat. Twiztid)<br>06 - Farewell<br>07 - Fuck You<br>08 - Perception Deception<br>09 - Enemy (Feat. Tech N9ne)<br>10 - Wake Up You're Dead<br>11 - Hate<br>12 - Until Then<br>13 - Million Miles Away<br>14 - Turn Back<br>15 - Full Moon<br>16 - One Of These Days (Feat. Tech N9ne &amp; Krizz Kaliko)<br>17 - Alien<br>18 - Last Will And Testament (Strange Music Pre-Order Digital Bonus Track) (192 kbps)</span><br><br><br></div><br><div style="background-color:#000000;width:490px;padding:0px;border-radius:15px 15px 15px 15px;"><br><span style="font-family:Tahoma"></span><a href="https://fastpic.ru/view/114/2021/0101/_880a31d8bd7de3987b3616555bf98763.jpg.html" target="_blank" rel="noopener external noreferrer"><img src="https://i114.fastpic.ru/thumb/2021/0101/63/_880a31d8bd7de3987b3616555bf98763.jpeg" style="max-width:100%;" alt=""></a><br><br><b><span style="color:#33FF33">Качество (Размер):</span> m4a ~256 kbps (96.2 Mb) </b><br><br><b><span style="color:#33FF33">Треклист:</span></b><br><span style="font-family:Georgia"><br>01. Divided We Stand<br>02. Audio Barricade<br>03. Just Like Nothing<br>04. The Ghost Of Injustice (Interlude) <br>05. Blood Paved Road <br>06. Fading Away <br>07. Three, Two, One <br>08. Vendetta <br>09. Nowhere To Run <br>10. Distress Call <br>11. Darkest Shade Of Grey<br>12. We All Fall Down <br>13. The Shadow Of Mortality (Interlude) <br>14. Time <br>15. Before We Say Goodbye<br>16. Catacomb (128 kbps)</span><br><br><br></div><br><div style="background-color:#000000;width:490px;padding:0px;border-radius:15px 15px 15px 15px;"><br><span style="font-family:Tahoma"></span><a href="https://fastpic.ru/view/114/2021/0101/_9c7308fefaf774f926ba70ecc342be8e.jpg.html" target="_blank" rel="noopener external noreferrer"><img src="https://i114.fastpic.ru/thumb/2021/0101/8e/_9c7308fefaf774f926ba70ecc342be8e.jpeg" style="max-width:100%;" alt=""></a><br><br><b><span style="color:#33FF33">Качество (Размер):</span> m4a ~256 kbps (95.8 Mb) </b><br><br><b><span style="color:#33FF33">Треклист:</span></b><br><span style="font-family:Georgia"><br>01. The Abyss (Intro)<br>02. Purgatory feat. Tech N9ne, Krizz Kaliko<br>03. War Within feat. Ces Cru<br>04. Tomorrow feat. Krizz Kaliko<br>05. Do You Know Where You Are feat. Tech N9ne, Twiztid<br>06. House of Cards feat. Kate Rose<br>07. Erased feat. Mackenzie O'Guin<br>08. Killing Me feat. Krizz Kaliko, Blaze Ya Dead Homie, The R.O.C<br>09. The Plague feat. Madchild, Ubiquitous<br>10. My Life feat. Wrekonize, Bernz<br>11. Black Ink feat. Mackenzie O'Guin<br>12. People of the Outside feat. Tyler Lyon<br>13. Your Creation<br>14. Nobody's Fool (256 kbps)</span><br><br><br></div><br><div style="background-color:#000000;width:100px;border-radius:10px 10px 10px 10px;border:1px solid #666666;"><br><a href="https://www.zippyshare.com/Alternativshik_6/aqr8zwwc/dir.html" target="_blank" rel="noopener external noreferrer"><b><span style="font-family:Georgia"><span style="color:#ffffff"><span style="font-size:10pt;">Zippyshare</span></span></span></b></a><br><br></div><br><br><iframe width="490" height="315" src="//www.youtube.com/embed/XFTPg8k1PIA" frameborder="0" allowfullscreen></iframe><br><iframe width="490" height="315" src="//www.youtube.com/embed/-3wjTVH0ceg" frameborder="0" allowfullscreen></iframe><br></div><br><br></div> ]]></turbo:content>
<category><![CDATA[Неформат]]></category>
<dc:creator>Alternativshik_6</dc:creator>
<pubDate>Fri, 01 Jan 2021 21:23:36 +0300</pubDate>
</item><item turbo="true">
<title>Hako Yamasaki - 飛・び・ま・す (Tobimasu) (1975)</title>
<guid isPermaLink="true">https://alterportal.net/neformat/170463-hako-yamasaki-tobimasu-1975.html</guid>
<link>https://alterportal.net/neformat/170463-hako-yamasaki-tobimasu-1975.html</link>
<description><![CDATA[<div style="text-align:center;"><img src="https://a.radikal.ru/a11/2101/e5/4ef7fb13f8f5.jpg" style="max-width:100%;" alt="Hako Yamasaki - 飛・び・ま・す (Tobimasu) (1975)"></div><div style="text-align:center;"><span style="font-family:Century Gothic"><b>Стиль:</b> <span style="color:#CC66CC">Acoustic / Folk / Singer-Songwriter</span><br><b>Треклист:</b><br>01. 望郷<br>02. さすらい<br>03. かざぐるま<br>04. 橋向こうの家<br>05. サヨナラの鐘<br>06. 竹とんぼ<br>07. 影が見えない<br>08. 気分を変えて<br>09. 飛びます<br>10. 子守唄<br>11. 男と女の部屋（ボーナストラック）<br><br>Протекция I Am Not Jesus, январь 1/1</span></div>]]></description>
<turbo:content><![CDATA[ <div class="quote_single"><div style="text-align:center;"><img src="https://a.radikal.ru/a14/2101/1a/bc2c9d513401.jpg" style="max-width:100%;" alt="Hako Yamasaki - 飛・び・ま・す (Tobimasu) (1975)"></div><br><div><a href="http://www.hako.esy.es/index.html" target="_blank" rel="noopener external noreferrer">official</a></div><br><div><span style="font-family:Century Gothic"><b>Стиль:</b> <span style="color:#CC66CC">Acoustic / Folk / Singer-Songwriter</span><br><b>Страна:</b><span style="color:#CC66CC"> Japan</span><br><b>Формат / Качество:</b> <span style="color:#CC66CC">MP3 CBR 320kbps</span><br><b>Треклист:</b><br>01. 望郷<br>02. さすらい<br>03. かざぐるま<br>04. 橋向こうの家<br>05. サヨナラの鐘<br>06. 竹とんぼ<br>07. 影が見えない<br>08. 気分を変えて<br>09. 飛びます<br>10. 子守唄<br>11. 男と女の部屋（ボーナストラック） </span></div><br><div><a href="https://www.mediafire.com/file/n8uuovwee4foamv/Hako_Yamasaki_-_%25E9%25A3%259B%25C2%25B7%25E3%2581%25B3%25C2%25B7%25E3%2581%25BE%25C2%25B7%25E3%2581%2599_%25281975%2529.rar/file" target="_blank" rel="noopener external noreferrer">Mediafire</a></div><br><iframe width="500" height="281" src="https://www.youtube.com/embed/hboXXVfj24M?controls=1" frameborder="0" allowfullscreen></iframe><br></div> ]]></turbo:content>
<category><![CDATA[Неформат]]></category>
<dc:creator>I Am Not Jesus</dc:creator>
<pubDate>Fri, 01 Jan 2021 21:11:21 +0300</pubDate>
</item><item turbo="true">
<title>Gridiron - The Other Side of Suffering (2021)</title>
<guid isPermaLink="true">https://alterportal.net/2021_albums/170462-gridiron-the-other-side-of-suffering-2021.html</guid>
<link>https://alterportal.net/2021_albums/170462-gridiron-the-other-side-of-suffering-2021.html</link>
<description><![CDATA[<img src="https://funkyimg.com/i/39KYT.jpg" style="max-width:100%;" alt="Gridiron - The Other Side of Suffering (2021)"><br><b>Стиль</b> :  <span style="color:#FFC69C">Metalcore</span><br><b>Треклист :</b><br>01. Inception (Intro by Misstiq) (1:07)<br>02. A Sight to Behold (3:25)<br>03. Become (4:30)<br>04. Lazarus (3:48)<br>05. Afterlife (4:13)<br>06. Eyes Wide Shut (5:15)<br>07. As the Pictures Burn (feat. Deanne Oliver) (4:45)<br>08. Blame It on the Fire (5:16)<br>09. Wretched Earth (feat. Jesse Zaraska) (3:52)<br>10. The Other Side of Suffering (5:47)<br>11. Justice in Honor (3:20)<br>12. World's End (6:47)<br>13. Of Blood &amp; Bone (4:08)]]></description>
<turbo:content><![CDATA[ <div style="background:#000000;"><br><img src="https://funkyimg.com/i/39KYT.jpg" style="max-width:100%;" alt="Gridiron - The Other Side of Suffering (2021)"><br><a href="https://www.facebook.com/wearegridiron/" target="_blank" rel="noopener external noreferrer"><b><span style="font-family:Tahoma"><span style="font-size:18pt;"><span style="color:#9C9C94">facebook</span></span></span></b></a><br><b>Стиль</b> :  <span style="color:#FFC69C">Metalcore</span><br><b>Страна</b> :  <span style="color:#FFC69C">USA</span><br><b>Дата релиза</b> :  <span style="color:#FFC69C">January 1, 2021</span><br><b>Формат</b> :  <span style="color:#FFC69C">MP3, CBR 320 kbps</span><br><b>Размер</b> :  <span style="color:#FFC69C">129 mb</span><br><b>Треклист :</b><br>01. Inception (Intro by Misstiq) (1:07)<br>02. A Sight to Behold (3:25)<br>03. Become (4:30)<br>04. Lazarus (3:48)<br>05. Afterlife (4:13)<br>06. Eyes Wide Shut (5:15)<br>07. As the Pictures Burn (feat. Deanne Oliver) (4:45)<br>08. Blame It on the Fire (5:16)<br>09. Wretched Earth (feat. Jesse Zaraska) (3:52)<br>10. The Other Side of Suffering (5:47)<br>11. Justice in Honor (3:20)<br>12. World's End (6:47)<br>13. Of Blood &amp; Bone (4:08)<br><br>Download<br><a href="https://www33.zippyshare.com/v/J0nmtP16/file.html" target="_blank" rel="noopener external noreferrer"><b><span style="font-family:Tahoma"><span style="font-size:10pt;"><span style="color:#9C9C94">ZIPPYSHARE</span></span></span></b></a>   I   <a href="https://yadi.sk/d/smV_aqcyN6zGSw" target="_blank" rel="noopener external noreferrer"><b><span style="font-family:Tahoma"><span style="font-size:10pt;"><span style="color:#9C9C94">YADISK</span></span></span></b></a><br><br><iframe width="420" height="255" src="https://www.youtube.com/embed/YdrbOLeOhkU" frameborder="0" allowfullscreen></iframe><iframe width="420" height="255" src="https://www.youtube.com/embed/rnCzVX3GwbY" frameborder="0" allowfullscreen></iframe></div> ]]></turbo:content>
<category><![CDATA[Альбомы 2021    | Сore]]></category>
<dc:creator>izuver</dc:creator>
<pubDate>Fri, 01 Jan 2021 01:52:20 +0300</pubDate>
</item><item turbo="true">
<title>Alterportal&#039;s ТОР 2020. Голосование</title>
<guid isPermaLink="true">https://alterportal.net/raznoe/170420-alterportals-tor-2020-golosovanie.html</guid>
<link>https://alterportal.net/raznoe/170420-alterportals-tor-2020-golosovanie.html</link>
<description><![CDATA[<span style="color:#33CCFF"><span style="font-size:12pt;"><b>Голосование за лучшие альбомы 2020 года</b></span></span><br><br>------------------------------------------------------------------------------<br><br>Правила такие же, как и в прошлом году, но обязательно их прочтите, нажав Далее. <br><br><span style="color:#00FF00"><b>Форматные альбомы/ЕР русских групп и групп бывшего СССР допускаются к голосованию!</b></span><br><br>Флуд запрещен! Любые вопросы по поводу темы, обсудить топы можно как обычно в <a href="https://alterportal.net/raznoe/164743-bolt0logija.html"><span style="color:#FFFF66"><b>болтологии</b></span></a>.]]></description>
<turbo:content><![CDATA[ <div><div style="width:590px;padding:10px;background-color:#0b0b0b;border-radius:10px 10px 10px 10px;"><br><span style="font-size:14pt;"><span style="color:#FF0000">Правила:</span></span><br><br>1. Голосование продлится с 1-го января 2021 по 31 января 2021 года (включительно). Итоги будут подведены позднее.<br>2. К голосованию допускаются пользователи, зарегистрировавшие свой аккаунт до 1 января 2021 года.<br>3. <b><span style="color:#3366FF">Каждый участник предлагает для голосования не менее 3 и не более 20 альбомов.</span></b><br>4. Альбомы должны быть расположены в порядке убывания, лучшие сверху.<br>5. К голосованию принимаются официальные "номерные" альбомы и EP , официальная дата выхода которых с 1 января по 31 декабря 2020 года. <span style="color:#CC0000">Альбомы русских групп и групп бывшего СССР допускаются к голосованию!</span><br>6. К голосованию не принимаются: сборники (ранее издававшегося материала), переиздания, концерты, синглы и релизы из раздела "Неформат". Если есть сомнения о попадании альбома в рамки правил, можно узнать официальную дату выхода альбома например на <b>itunes.apple.com</b>.<br>7. Если ваше мнение изменилось, вы смело можете изменить, дополнить свою двадцатку до окончания голосования обратившись за помощью к модератору или администратору. После этого тема закрывается и начинается подсчет голосов.<br>8. Баллы начисляются: за первое место - 20 баллов, за 2-е - 19..., за 20-е место - 1 балл. Если у вас не двадцать альбомов, а меньше, то первое место также получает 20 баллов, остальные - 19, 18 и т.д.<br>9. Победители вычисляются арифметическим подсчетом суммы баллов. Подсчет начнется после закрытия голосования.<br><br><span style="font-size:12pt;"><span style="color:#FF0000">Свой топ пишем тут, в комментариях!<br>Не в личку, не через добавление новостей, не через подпись!</span><br><br><span style="color:#3366FF"><b>Пример оформления топа Вы можете увидеть в первом комментарии данной темы.</b></span></span><br><br><span style="font-size:12pt;">Админ-модерам просьба редактировать топы, уже присланные и новые, приводя к формату для программы и исправляя ошибки в написании групп.</span><br><br><b><span style="color:#FF0000">Замечание</b></span><br>Не будьте обезьянами, изобретающими велосипеды! Посмотрите, как другие оформили топ, и делайте так же! После 1. ПРОБЕЛ, а потом идет название и т.д.<br><br><b><span style="color:#FF0000">Правильное написание некоторых альбомов:</span></b><br><span style="color:#66FFFF">...</span><br><br><b><span style="color:#FF0000">Альбомы, которые можно включать в топ 2020:</span></b><br><span style="color:#66FFFF">...</span><br><br><b><span style="color:#FF0000">Альбомы, которые нельзя включать в топ 2020:</span></b><br><span style="color:#66FFFF">...</span><br></div></div> ]]></turbo:content>
<category><![CDATA[Разное]]></category>
<dc:creator>Vetal</dc:creator>
<pubDate>Fri, 01 Jan 2021 00:00:01 +0300</pubDate>
</item><item turbo="true">
<title>Олександр Василенко - До 11-річчя Збройних Сил України (2002)</title>
<guid isPermaLink="true">https://alterportal.net/neformat/170441-oleksandr-vasilenko-do-11-richchja-zbrojnih-sil-ukrajini-2002.html</guid>
<link>https://alterportal.net/neformat/170441-oleksandr-vasilenko-do-11-richchja-zbrojnih-sil-ukrajini-2002.html</link>
<description><![CDATA[<img src="https://c.radikal.ru/c35/2012/62/ef344a98e4dc.jpg" style="max-width:100%;" alt="Олександр Василенко - До 11-річчя Збройних Сил України (2002)"><br><b>Cтиль: Retro Pop<br>Треклист:</b><br>1. Дорога моя земля (03:55)<br>2. Яблуневий дзвін (03:50)<br>]]></description>
<turbo:content><![CDATA[ <img src="https://d.radikal.ru/d09/2012/d2/f5204eefc433.jpg" style="max-width:100%;" alt="Олександр Василенко - До 11-річчя Збройних Сил України (2002)"><br><br><b>Исполнитель: Александр Василенко<br>Страна: СССР/Украина<br>Стиль: Retro Pop, Vocal Pop<br>Дата Релиза: 2002<br>Формат: MP3<br>Битрейт: сbr 320 kbps<br>Размер: 14 mb<br>Треклист:</b><br><br>1. Дорога моя земля (03:55)<br>— муз. Володимир Павліковський, сл. Зоя Ружин<br>2. Яблуневий дзвін (03:50)<br>— муз. Володимир Павліковський, сл. Зоя Ружин<br><br><a href="https://www49.zippyshare.com/v/DiNId7rd/file.html" target="_blank" rel="noopener external noreferrer">Прослушка!</a><br><br><a href="https://www32.zippyshare.com/v/2XGu4eqI/file.html" target="_blank" rel="noopener external noreferrer">Cкачать</a><br><br><a href="http://www.uaestrada.org/spivaki/vasilenko_oleksandr/" target="_blank" rel="noopener external noreferrer">официальный сайт</a> ]]></turbo:content>
<category><![CDATA[Неформат]]></category>
<dc:creator>Wild.E.Coyote</dc:creator>
<pubDate>Thu, 31 Dec 2020 20:07:04 +0300</pubDate>
</item><item turbo="true">
<title>Witchcraft - В Твоём Море</title>
<guid isPermaLink="true">https://alterportal.net/video/170444-witchcraft-v-tvoem-more-2020.html</guid>
<link>https://alterportal.net/video/170444-witchcraft-v-tvoem-more-2020.html</link>
<description><![CDATA[<img src="https://d.radikal.ru/d29/2012/32/77469dc55d65.jpg" style="max-width:100%;" alt="Witchcraft - В Твоём Море"><br><b>Стиль: Gothic Rock<br>Формат: MP4 1080p<br>Размер: 95 mb</b>]]></description>
<turbo:content><![CDATA[ <img src="https://b.radikal.ru/b01/2012/db/9ff51b292d1b.jpg" style="max-width:100%;" alt="Witchcraft - В Твоём Море"><br><br><b>Исполнитель: Witchcraft<br>Cтрана: Россия<br>Стиль: Gothic Rock<br>Дата Релиза: 2020<br>Формат: MP4 1080p<br>Размер: 95 mb<br>Треклист: </b><br><br>1.В Твоём Море<br><br><center><iframe width="356" height="200" src="https://www.youtube.com/embed/tu0Xne6nxto?feature=oembed" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></center><br><br><a href="https://www97.zippyshare.com/v/yZT6Jeds/file.html" target="_blank" rel="noopener external noreferrer">Скачать!</a><br><br><a href="https://vk.com/witchcraft" target="_blank" rel="noopener external noreferrer">В Твоём ВК!</a> ]]></turbo:content>
<category><![CDATA[Video       | ex-USSR       | Rock]]></category>
<dc:creator>Wild.E.Coyote</dc:creator>
<pubDate>Thu, 31 Dec 2020 20:05:51 +0300</pubDate>
</item><item turbo="true">
<title>Владимир Пресняков - Слушая Тишину (2020)</title>
<guid isPermaLink="true">https://alterportal.net/neformat/170446-vladimir-presnjakov-slushaja-tishinu-2020.html</guid>
<link>https://alterportal.net/neformat/170446-vladimir-presnjakov-slushaja-tishinu-2020.html</link>
<description><![CDATA[<img src="https://c.radikal.ru/c26/2012/99/1b0010fb616d.jpg" style="max-width:100%;" alt="Владимир Пресняков - Слушая Тишину (2020)"><br><b>Cтиль: Pop<br>Треклист:</b><br>01. Снег<br>02. Слушая тишину<br>03. Ты у меня одна<br>04. Достучаться до небес<br>05. Странная<br>06. Если нет рядом тебя<br>07. Неземная<br>08. Бабочка и мотылёк<br>09. Только где ты<br>10. Я в облака<br>11. Зурбаган 2.0<br>12. Два сердечка<br>13. Дыши<br>14. Kissлород (&amp; Наталья Подольская) <br>15. Я всё помню (&amp; Наталья Подольская)<br>16. Сердце пленных не берёт (&amp; Никита Пресняков) <br>17. Я помню<br>18. Первый снег (cover)<br><div style="text-align:center;">19. Почему небо плачет (&amp; Наталья Подольская) <br>20. Странник (&amp; Влад Соколовский) <br>21. Ты не птица (&amp; Дмитрий Колдун)</div>]]></description>
<turbo:content><![CDATA[ <img src="https://a.radikal.ru/a01/2012/c5/729737cf0dee.jpg" style="max-width:100%;" alt="Владимир Пресняков - Слушая Тишину (2020)"><br><br><b>Исполнитель: Владимир Пресняков<br>Cтиль: Pop Music<br>Страна: Россия<br>Дата Релиза: 2020<br>Формат: MP3<br>Битрейт: сbr 320 kbps<br>Размер: 185 mb<br>Треклист:</b><br><br>01. Снег<br>02. Слушая тишину<br>03. Ты у меня одна<br>04. Достучаться до небес<br>05. Странная<br>06. Если нет рядом тебя<br>07. Неземная<br>08. Бабочка и мотылёк<br>09. Только где ты<br>10. Я в облака<br>11. Зурбаган 2.0<br>12. Два сердечка<br>13. Дыши<br>14. Kissлород (&amp; Наталья Подольская) <br>15. Я всё помню (&amp; Наталья Подольская)<br>16. Сердце пленных не берёт (&amp; Никита Пресняков) <br>17. Я помню<br>18. Первый снег (cover)<br>19. Почему небо плачет (&amp; Наталья Подольская) <br>20. Странник (&amp; Влад Соколовский) <br>21. Ты не птица (&amp; Дмитрий Колдун)<br><br><center><iframe width="267" height="200" src="https://www.youtube.com/embed/loSebeLuyrw?feature=oembed" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></center><br><br><a href="https://www31.zippyshare.com/v/BPoO4c8p/file.html" target="_blank" rel="noopener external noreferrer">Скачать!</a><br><br><a href="https://vk.com/vl_presnyakov" target="_blank" rel="noopener external noreferrer">ВК!</a> ]]></turbo:content>
<category><![CDATA[Неформат]]></category>
<dc:creator>Wild.E.Coyote</dc:creator>
<pubDate>Thu, 31 Dec 2020 20:04:37 +0300</pubDate>
</item><item turbo="true">
<title>Cult Of Luna - Three Bridges (single) (2020)</title>
<guid isPermaLink="true">https://alterportal.net/newsongs/170460-cult-of-luna-three-bridges-single-2020.html</guid>
<link>https://alterportal.net/newsongs/170460-cult-of-luna-three-bridges-single-2020.html</link>
<description><![CDATA[<img src="https://funkyimg.com/i/39KuG.jpg" style="max-width:100%;" alt="Cult Of Luna - Three Bridges (single) (2020)"><br><b>Стиль: <span style="color:#E76363">Post-Metal / Experimental</span></b><br><b>Треклист:</b><br>1. Three Bridges]]></description>
<turbo:content><![CDATA[ <img src="https://funkyimg.com/i/39KuG.jpg" style="max-width:100%;" alt="Cult Of Luna - Three Bridges (single) (2020)"><br><b><a href="https://www.cultofluna.com" target="_blank" rel="noopener external noreferrer">Official Website</a> <span style="color:#00FFFF">|</span> <a href="https://www.facebook.com/cultoflunamusic" target="_blank" rel="noopener external noreferrer">Facebook</a> <span style="color:#00FFFF">|</span> <a href="https://www.instagram.com/cultofluna" target="_blank" rel="noopener external noreferrer">Instagram</a> <span style="color:#00FFFF">|</span> <a href="https://twitter.com/Cultofluna_off" target="_blank" rel="noopener external noreferrer">Twitter</a> <span style="color:#00FFFF">|</span> <a href="https://open.spotify.com/artist/7E7fJJpdVgr1F3pfAfRtHe" target="_blank" rel="noopener external noreferrer">Spotify</a> <span style="color:#00FFFF">|</span> <a href="https://music.apple.com/us/artist/cult-of-luna/42407392" target="_blank" rel="noopener external noreferrer">Apple Music</a><br><br>Страна: <span style="color:#E76363">Sweden, Umeå</span><br>Стиль: <span style="color:#E76363">Post-Metal / Experimental</span> <br>Формат: <span style="color:#E76363">MP3 CBR 320 kbps</span><br>Лейбл: <span style="color:#E76363">Self-released</span><br>Дата релиза: <span style="color:#E76363">Dec 9, 2020</span></b><br><br><b>Треклист:</b><br>1. Three Bridges<br><br><iframe width="460" height="215" src="https://www.youtube.com/embed/wgN3OV6XtTY" frameborder="0" allowfullscreen></iframe><br><br><b><a href="https://www84.zippyshare.com/v/TdXKibHN/file.html" target="_blank" rel="noopener external noreferrer">Download</a></b> ]]></turbo:content>
<category><![CDATA[Новые треки         | Metal]]></category>
<dc:creator>Lestat</dc:creator>
<pubDate>Thu, 31 Dec 2020 20:03:53 +0300</pubDate>
</item><item turbo="true">
<title>Cold Subject - Singles (2020)</title>
<guid isPermaLink="true">https://alterportal.net/core/170459-cold-subject-singles-2020.html</guid>
<link>https://alterportal.net/core/170459-cold-subject-singles-2020.html</link>
<description><![CDATA[<img src="https://funkyimg.com/i/39Kp9.jpg" style="max-width:100%;" alt="Cold Subject - Singles (2020)"><br><b>Стиль: <span style="color:#E76363">Metalcore / Hard Rock</span></b><br><b>Треклист:</b><br>1. UFO<br>2. Past]]></description>
<turbo:content><![CDATA[ <img src="https://funkyimg.com/i/39Kp9.jpg" style="max-width:100%;" alt="Cold Subject - Singles (2020)"><br><b><a href="http://www.coldsubject.com" target="_blank" rel="noopener external noreferrer">Official Website</a> <span style="color:#00FFFF">|</span> <a href="https://www.facebook.com/ColdSubejct" target="_blank" rel="noopener external noreferrer">Facebook</a> <span style="color:#00FFFF">|</span> <a href="https://www.instagram.com/coldsubjectfl" target="_blank" rel="noopener external noreferrer">Instagram</a> <span style="color:#00FFFF">|</span> <a href="https://open.spotify.com/artist/0c25fpYwnfaUskPSPaOE9w" target="_blank" rel="noopener external noreferrer">Spotify</a> <span style="color:#00FFFF">|</span> <a href="https://music.apple.com/us/artist/cold-subject/1459827143" target="_blank" rel="noopener external noreferrer">Apple Music</a><br><br>Страна: <span style="color:#E76363">USA, Frorida, Orlando</span><br>Стиль: <span style="color:#E76363">Metalcore / Hard Rock</span> <br>Формат: <span style="color:#E76363">MP3 CBR 320 kbps</span><br>Лейбл: <span style="color:#E76363">Chilled Records</span><br>Дата релиза: <span style="color:#E76363">2020</span></b><br><br><b>Треклист:</b><br>1. <a href="https://www84.zippyshare.com/v/Q8Xgw1cT/file.html" target="_blank" rel="noopener external noreferrer">UFO</a><br>2. <a href="https://www84.zippyshare.com/v/MUyD5FQo/file.html" target="_blank" rel="noopener external noreferrer">Past</a><br><br><iframe width="460" height="215" src="https://www.youtube.com/embed/XULlFqkiG9Y" frameborder="0" allowfullscreen></iframe><br><br><iframe width="460" height="215" src="https://www.youtube.com/embed/cPcUvnwRGiU" frameborder="0" allowfullscreen></iframe> ]]></turbo:content>
<category><![CDATA[Сore          | Новые треки]]></category>
<dc:creator>Lestat</dc:creator>
<pubDate>Thu, 31 Dec 2020 20:03:31 +0300</pubDate>
</item></channel></rss>
`
