package internal

import (
	"context"
	"github.com/go-pkgz/lgr"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGetMetalParser_Parse(t *testing.T) {
	feed, err := gofeed.NewParser().Parse(strings.NewReader(rssExample))
	require.NoError(t, err)

	//feed, err := gofeed.NewParser().ParseURL(rssFeed)
	//require.NoError(t, err)

	p := NewAlterportalParser(lgr.New())
	ctx := context.Background()
	for _, i := range feed.Items {
		result, err := p.Parse(ctx, i)
		if err == errSkipItem {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("result %v\n", result)
	}
}

const rssExample = `
<?xml version="1.0" encoding="utf-8"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:content="http://purl.org/rss/1.0/modules/content/"
    xmlns:media="http://search.yahoo.com/mrss/"
    xmlns:turbo="http://turbo.yandex.ru" version="2.0">
    <channel>
        <title>Alterportal.net V2</title>
        <link>https://alterportal.net/</link>
        <language>ru</language>
        <description>Alterportal.net V2</description>
        <generator>DataLife Engine</generator>
        <item turbo="true">
            <title>Gridiron - The Other Side of Suffering (2021)</title>
            <guid isPermaLink="true">https://alterportal.net/2021_albums/170462-gridiron-the-other-side-of-suffering-2021.html</guid>
            <link>https://alterportal.net/2021_albums/170462-gridiron-the-other-side-of-suffering-2021.html</link>
            <description><![CDATA[<img src="https://funkyimg.com/i/39KYT.jpg" style="max-width:100%;" alt="Gridiron - The Other Side of Suffering (2021)"><br><b>Стиль</b> :  <span style="color:#FFC69C">Metalcore</span><br><b>Треклист :</b><br>01. Inception (Intro by Misstiq) (1:07)<br>02. A Sight to Behold (3:25)<br>03. Become (4:30)<br>04. Lazarus (3:48)<br>05. Afterlife (4:13)<br>06. Eyes Wide Shut (5:15)<br>07. As the Pictures Burn (feat. Deanne Oliver) (4:45)<br>08. Blame It on the Fire (5:16)<br>09. Wretched Earth (feat. Jesse Zaraska) (3:52)<br>10. The Other Side of Suffering (5:47)<br>11. Justice in Honor (3:20)<br>12. World's End (6:47)<br>13. Of Blood &amp; Bone (4:08)]]></description>
            <turbo:content><![CDATA[ <div style="background:#000000;"><br><img src="https://funkyimg.com/i/39KYT.jpg" style="max-width:100%;" alt="Gridiron - The Other Side of Suffering (2021)"><br><a href="https://www.facebook.com/wearegridiron/" target="_blank" rel="noopener external noreferrer"><b><span style="font-family:Tahoma"><span style="font-size:18pt;"><span style="color:#9C9C94">facebook</span></span></span></b></a><br><b>Стиль</b> :  <span style="color:#FFC69C">Metalcore</span><br><b>Страна</b> :  <span style="color:#FFC69C">USA</span><br><b>Дата релиза</b> :  <span style="color:#FFC69C">January 1, 2021</span><br><b>Формат</b> :  <span style="color:#FFC69C">MP3, CBR 320 kbps</span><br><b>Размер</b> :  <span style="color:#FFC69C">129 mb</span><br><b>Треклист :</b><br>01. Inception (Intro by Misstiq) (1:07)<br>02. A Sight to Behold (3:25)<br>03. Become (4:30)<br>04. Lazarus (3:48)<br>05. Afterlife (4:13)<br>06. Eyes Wide Shut (5:15)<br>07. As the Pictures Burn (feat. Deanne Oliver) (4:45)<br>08. Blame It on the Fire (5:16)<br>09. Wretched Earth (feat. Jesse Zaraska) (3:52)<br>10. The Other Side of Suffering (5:47)<br>11. Justice in Honor (3:20)<br>12. World's End (6:47)<br>13. Of Blood &amp; Bone (4:08)<br><br>Download<br><a href="https://www33.zippyshare.com/v/J0nmtP16/file.html" target="_blank" rel="noopener external noreferrer"><b><span style="font-family:Tahoma"><span style="font-size:10pt;"><span style="color:#9C9C94">ZIPPYSHARE</span></span></span></b></a>   I   <a href="https://yadi.sk/d/smV_aqcyN6zGSw" target="_blank" rel="noopener external noreferrer"><b><span style="font-family:Tahoma"><span style="font-size:10pt;"><span style="color:#9C9C94">YADISK</span></span></span></b></a><br><br><iframe width="420" height="255" src="https://www.youtube.com/embed/YdrbOLeOhkU" frameborder="0" allowfullscreen></iframe><iframe width="420" height="255" src="https://www.youtube.com/embed/rnCzVX3GwbY" frameborder="0" allowfullscreen></iframe></div> ]]></turbo:content>
            <category><![CDATA[Альбомы 2021 | Сore]]></category>
            <dc:creator>izuver</dc:creator>
            <pubDate>Fri, 01 Jan 2021 01:52:20 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Alterportal&#039;s ТОР 2020. Голосование</title>
            <guid isPermaLink="true">https://alterportal.net/raznoe/170420-alterportals-tor-2020-golosovanie.html</guid>
            <link>https://alterportal.net/raznoe/170420-alterportals-tor-2020-golosovanie.html</link>
            <description><![CDATA[<span style="color:#33CCFF"><span style="font-size:12pt;"><b>Голосование за лучшие альбомы 2020 года</b></span></span><br><br>------------------------------------------------------------------------------<br><br>Правила такие же, как и в прошлом году, но обязательно их прочтите, нажав Далее. <br><br><span style="color:#00FF00"><b>Форматные альбомы/ЕР русских групп и групп бывшего СССР допускаются к голосованию!</b></span><br><br>Флуд запрещен! Любые вопросы по поводу темы, обсудить топы можно как обычно в <a href="https://alterportal.net/raznoe/164743-bolt0logija.html"><span style="color:#FFFF66"><b>болтологии</b></span></a>.]]></description>
            <turbo:content><![CDATA[ <div><div style="width:590px;padding:10px;background-color:#0b0b0b;border-radius:10px 10px 10px 10px;"><br><span style="font-size:14pt;"><span style="color:#FF0000">Правила:</span></span><br><br>1. Голосование продлится с 1-го января 2021 по 31 января 2021 года (включительно). Итоги будут подведены позднее.<br>2. К голосованию допускаются пользователи, зарегистрировавшие свой аккаунт до 1 января 2021 года.<br>3. <b><span style="color:#3366FF">Каждый участник предлагает для голосования не менее 3 и не более 20 альбомов.</span></b><br>4. Альбомы должны быть расположены в порядке убывания, лучшие сверху.<br>5. К голосованию принимаются официальные "номерные" альбомы и EP , официальная дата выхода которых с 1 января по 31 декабря 2020 года. <span style="color:#CC0000">Альбомы русских групп и групп бывшего СССР допускаются к голосованию!</span><br>6. К голосованию не принимаются: сборники (ранее издававшегося материала), переиздания, концерты, синглы и релизы из раздела "Неформат". Если есть сомнения о попадании альбома в рамки правил, можно узнать официальную дату выхода альбома например на <b>itunes.apple.com</b>.<br>7. Если ваше мнение изменилось, вы смело можете изменить, дополнить свою двадцатку до окончания голосования обратившись за помощью к модератору или администратору. После этого тема закрывается и начинается подсчет голосов.<br>8. Баллы начисляются: за первое место - 20 баллов, за 2-е - 19..., за 20-е место - 1 балл. Если у вас не двадцать альбомов, а меньше, то первое место также получает 20 баллов, остальные - 19, 18 и т.д.<br>9. Победители вычисляются арифметическим подсчетом суммы баллов. Подсчет начнется после закрытия голосования.<br><br><span style="font-size:12pt;"><span style="color:#FF0000">Свой топ пишем тут, в комментариях!<br>Не в личку, не через добавление новостей, не через подпись!</span><br><br><span style="color:#3366FF"><b>Пример оформления топа Вы можете увидеть в первом комментарии данной темы.</b></span></span><br><br><span style="font-size:12pt;">Админ-модерам просьба редактировать топы, уже присланные и новые, приводя к формату для программы и исправляя ошибки в написании групп.</span><br><br><b><span style="color:#FF0000">Замечание</b></span><br>Не будьте обезьянами, изобретающими велосипеды! Посмотрите, как другие оформили топ, и делайте так же! После 1. ПРОБЕЛ, а потом идет название и т.д.<br><br><b><span style="color:#FF0000">Правильное написание некоторых альбомов:</span></b><br><span style="color:#66FFFF">...</span><br><br><b><span style="color:#FF0000">Альбомы, которые можно включать в топ 2020:</span></b><br><span style="color:#66FFFF">...</span><br><br><b><span style="color:#FF0000">Альбомы, которые нельзя включать в топ 2020:</span></b><br><span style="color:#66FFFF">...</span><br></div></div> ]]></turbo:content>
            <category><![CDATA[Разное]]></category>
            <dc:creator>Vetal</dc:creator>
            <pubDate>Fri, 01 Jan 2021 00:00:01 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Олександр Василенко - До 11-річчя Збройних Сил України (2002)</title>
            <guid isPermaLink="true">https://alterportal.net/neformat/170441-oleksandr-vasilenko-do-11-richchja-zbrojnih-sil-ukrajini-2002.html</guid>
            <link>https://alterportal.net/neformat/170441-oleksandr-vasilenko-do-11-richchja-zbrojnih-sil-ukrajini-2002.html</link>
            <description><![CDATA[<img src="https://c.radikal.ru/c35/2012/62/ef344a98e4dc.jpg" style="max-width:100%;" alt="Олександр Василенко - До 11-річчя Збройних Сил України (2002)"><br><b>Cтиль: Retro Pop<br>Треклист:</b><br>1. Дорога моя земля (03:55)<br>2. Яблуневий дзвін (03:50)<br>]]></description>
            <turbo:content><![CDATA[ <img src="https://d.radikal.ru/d09/2012/d2/f5204eefc433.jpg" style="max-width:100%;" alt="Олександр Василенко - До 11-річчя Збройних Сил України (2002)"><br><br><b>Исполнитель: Александр Василенко<br>Страна: СССР/Украина<br>Стиль: Retro Pop, Vocal Pop<br>Дата Релиза: 2002<br>Формат: MP3<br>Битрейт: сbr 320 kbps<br>Размер: 14 mb<br>Треклист:</b><br><br>1. Дорога моя земля (03:55)<br>— муз. Володимир Павліковський, сл. Зоя Ружин<br>2. Яблуневий дзвін (03:50)<br>— муз. Володимир Павліковський, сл. Зоя Ружин<br><br><a href="https://www49.zippyshare.com/v/DiNId7rd/file.html" target="_blank" rel="noopener external noreferrer">Прослушка!</a><br><br><a href="https://www32.zippyshare.com/v/2XGu4eqI/file.html" target="_blank" rel="noopener external noreferrer">Cкачать</a><br><br><a href="http://www.uaestrada.org/spivaki/vasilenko_oleksandr/" target="_blank" rel="noopener external noreferrer">официальный сайт</a> ]]></turbo:content>
            <category><![CDATA[Неформат]]></category>
            <dc:creator>Wild.E.Coyote</dc:creator>
            <pubDate>Thu, 31 Dec 2020 20:07:04 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Witchcraft - В Твоём Море</title>
            <guid isPermaLink="true">https://alterportal.net/video/170444-witchcraft-v-tvoem-more-2020.html</guid>
            <link>https://alterportal.net/video/170444-witchcraft-v-tvoem-more-2020.html</link>
            <description><![CDATA[<img src="https://d.radikal.ru/d29/2012/32/77469dc55d65.jpg" style="max-width:100%;" alt="Witchcraft - В Твоём Море"><br><b>Стиль: Gothic Rock<br>Формат: MP4 1080p<br>Размер: 95 mb</b>]]></description>
            <turbo:content><![CDATA[ <img src="https://b.radikal.ru/b01/2012/db/9ff51b292d1b.jpg" style="max-width:100%;" alt="Witchcraft - В Твоём Море"><br><br><b>Исполнитель: Witchcraft<br>Cтрана: Россия<br>Стиль: Gothic Rock<br>Дата Релиза: 2020<br>Формат: MP4 1080p<br>Размер: 95 mb<br>Треклист: </b><br><br>1.В Твоём Море<br><br><center><iframe width="356" height="200" src="https://www.youtube.com/embed/tu0Xne6nxto?feature=oembed" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></center><br><br><a href="https://www97.zippyshare.com/v/yZT6Jeds/file.html" target="_blank" rel="noopener external noreferrer">Скачать!</a><br><br><a href="https://vk.com/witchcraft" target="_blank" rel="noopener external noreferrer">В Твоём ВК!</a> ]]></turbo:content>
            <category><![CDATA[Video    | ex-USSR    | Rock]]></category>
            <dc:creator>Wild.E.Coyote</dc:creator>
            <pubDate>Thu, 31 Dec 2020 20:05:51 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Владимир Пресняков - Слушая Тишину (2020)</title>
            <guid isPermaLink="true">https://alterportal.net/neformat/170446-vladimir-presnjakov-slushaja-tishinu-2020.html</guid>
            <link>https://alterportal.net/neformat/170446-vladimir-presnjakov-slushaja-tishinu-2020.html</link>
            <description><![CDATA[<img src="https://c.radikal.ru/c26/2012/99/1b0010fb616d.jpg" style="max-width:100%;" alt="Владимир Пресняков - Слушая Тишину (2020)"><br><b>Cтиль: Pop<br>Треклист:</b><br>01. Снег<br>02. Слушая тишину<br>03. Ты у меня одна<br>04. Достучаться до небес<br>05. Странная<br>06. Если нет рядом тебя<br>07. Неземная<br>08. Бабочка и мотылёк<br>09. Только где ты<br>10. Я в облака<br>11. Зурбаган 2.0<br>12. Два сердечка<br>13. Дыши<br>14. Kissлород (&amp; Наталья Подольская) <br>15. Я всё помню (&amp; Наталья Подольская)<br>16. Сердце пленных не берёт (&amp; Никита Пресняков) <br>17. Я помню<br>18. Первый снег (cover)<br><div style="text-align:center;">19. Почему небо плачет (&amp; Наталья Подольская) <br>20. Странник (&amp; Влад Соколовский) <br>21. Ты не птица (&amp; Дмитрий Колдун)</div>]]></description>
            <turbo:content><![CDATA[ <img src="https://a.radikal.ru/a01/2012/c5/729737cf0dee.jpg" style="max-width:100%;" alt="Владимир Пресняков - Слушая Тишину (2020)"><br><br><b>Исполнитель: Владимир Пресняков<br>Cтиль: Pop Music<br>Страна: Россия<br>Дата Релиза: 2020<br>Формат: MP3<br>Битрейт: сbr 320 kbps<br>Размер: 185 mb<br>Треклист:</b><br><br>01. Снег<br>02. Слушая тишину<br>03. Ты у меня одна<br>04. Достучаться до небес<br>05. Странная<br>06. Если нет рядом тебя<br>07. Неземная<br>08. Бабочка и мотылёк<br>09. Только где ты<br>10. Я в облака<br>11. Зурбаган 2.0<br>12. Два сердечка<br>13. Дыши<br>14. Kissлород (&amp; Наталья Подольская) <br>15. Я всё помню (&amp; Наталья Подольская)<br>16. Сердце пленных не берёт (&amp; Никита Пресняков) <br>17. Я помню<br>18. Первый снег (cover)<br>19. Почему небо плачет (&amp; Наталья Подольская) <br>20. Странник (&amp; Влад Соколовский) <br>21. Ты не птица (&amp; Дмитрий Колдун)<br><br><center><iframe width="267" height="200" src="https://www.youtube.com/embed/loSebeLuyrw?feature=oembed" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></center><br><br><a href="https://www31.zippyshare.com/v/BPoO4c8p/file.html" target="_blank" rel="noopener external noreferrer">Скачать!</a><br><br><a href="https://vk.com/vl_presnyakov" target="_blank" rel="noopener external noreferrer">ВК!</a> ]]></turbo:content>
            <category><![CDATA[Неформат]]></category>
            <dc:creator>Wild.E.Coyote</dc:creator>
            <pubDate>Thu, 31 Dec 2020 20:04:37 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Cult Of Luna - Three Bridges (single) (2020)</title>
            <guid isPermaLink="true">https://alterportal.net/newsongs/170460-cult-of-luna-three-bridges-single-2020.html</guid>
            <link>https://alterportal.net/newsongs/170460-cult-of-luna-three-bridges-single-2020.html</link>
            <description><![CDATA[<img src="https://funkyimg.com/i/39KuG.jpg" style="max-width:100%;" alt="Cult Of Luna - Three Bridges (single) (2020)"><br><b>Стиль: <span style="color:#E76363">Post-Metal / Experimental</span></b><br><b>Треклист:</b><br>1. Three Bridges]]></description>
            <turbo:content><![CDATA[ <img src="https://funkyimg.com/i/39KuG.jpg" style="max-width:100%;" alt="Cult Of Luna - Three Bridges (single) (2020)"><br><b><a href="https://www.cultofluna.com" target="_blank" rel="noopener external noreferrer">Official Website</a> <span style="color:#00FFFF">|</span> <a href="https://www.facebook.com/cultoflunamusic" target="_blank" rel="noopener external noreferrer">Facebook</a> <span style="color:#00FFFF">|</span> <a href="https://www.instagram.com/cultofluna" target="_blank" rel="noopener external noreferrer">Instagram</a> <span style="color:#00FFFF">|</span> <a href="https://twitter.com/Cultofluna_off" target="_blank" rel="noopener external noreferrer">Twitter</a> <span style="color:#00FFFF">|</span> <a href="https://open.spotify.com/artist/7E7fJJpdVgr1F3pfAfRtHe" target="_blank" rel="noopener external noreferrer">Spotify</a> <span style="color:#00FFFF">|</span> <a href="https://music.apple.com/us/artist/cult-of-luna/42407392" target="_blank" rel="noopener external noreferrer">Apple Music</a><br><br>Страна: <span style="color:#E76363">Sweden, Umeå</span><br>Стиль: <span style="color:#E76363">Post-Metal / Experimental</span> <br>Формат: <span style="color:#E76363">MP3 CBR 320 kbps</span><br>Лейбл: <span style="color:#E76363">Self-released</span><br>Дата релиза: <span style="color:#E76363">Dec 9, 2020</span></b><br><br><b>Треклист:</b><br>1. Three Bridges<br><br><iframe width="460" height="215" src="https://www.youtube.com/embed/wgN3OV6XtTY" frameborder="0" allowfullscreen></iframe><br><br><b><a href="https://www84.zippyshare.com/v/TdXKibHN/file.html" target="_blank" rel="noopener external noreferrer">Download</a></b> ]]></turbo:content>
            <category><![CDATA[Новые треки      | Metal]]></category>
            <dc:creator>Lestat</dc:creator>
            <pubDate>Thu, 31 Dec 2020 20:03:53 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Cold Subject - Singles (2020)</title>
            <guid isPermaLink="true">https://alterportal.net/core/170459-cold-subject-singles-2020.html</guid>
            <link>https://alterportal.net/core/170459-cold-subject-singles-2020.html</link>
            <description><![CDATA[<img src="https://funkyimg.com/i/39Kp9.jpg" style="max-width:100%;" alt="Cold Subject - Singles (2020)"><br><b>Стиль: <span style="color:#E76363">Metalcore / Hard Rock</span></b><br><b>Треклист:</b><br>1. UFO<br>2. Past]]></description>
            <turbo:content><![CDATA[ <img src="https://funkyimg.com/i/39Kp9.jpg" style="max-width:100%;" alt="Cold Subject - Singles (2020)"><br><b><a href="http://www.coldsubject.com" target="_blank" rel="noopener external noreferrer">Official Website</a> <span style="color:#00FFFF">|</span> <a href="https://www.facebook.com/ColdSubejct" target="_blank" rel="noopener external noreferrer">Facebook</a> <span style="color:#00FFFF">|</span> <a href="https://www.instagram.com/coldsubjectfl" target="_blank" rel="noopener external noreferrer">Instagram</a> <span style="color:#00FFFF">|</span> <a href="https://open.spotify.com/artist/0c25fpYwnfaUskPSPaOE9w" target="_blank" rel="noopener external noreferrer">Spotify</a> <span style="color:#00FFFF">|</span> <a href="https://music.apple.com/us/artist/cold-subject/1459827143" target="_blank" rel="noopener external noreferrer">Apple Music</a><br><br>Страна: <span style="color:#E76363">USA, Frorida, Orlando</span><br>Стиль: <span style="color:#E76363">Metalcore / Hard Rock</span> <br>Формат: <span style="color:#E76363">MP3 CBR 320 kbps</span><br>Лейбл: <span style="color:#E76363">Chilled Records</span><br>Дата релиза: <span style="color:#E76363">2020</span></b><br><br><b>Треклист:</b><br>1. <a href="https://www84.zippyshare.com/v/Q8Xgw1cT/file.html" target="_blank" rel="noopener external noreferrer">UFO</a><br>2. <a href="https://www84.zippyshare.com/v/MUyD5FQo/file.html" target="_blank" rel="noopener external noreferrer">Past</a><br><br><iframe width="460" height="215" src="https://www.youtube.com/embed/XULlFqkiG9Y" frameborder="0" allowfullscreen></iframe><br><br><iframe width="460" height="215" src="https://www.youtube.com/embed/cPcUvnwRGiU" frameborder="0" allowfullscreen></iframe> ]]></turbo:content>
            <category><![CDATA[Сore       | Новые треки]]></category>
            <dc:creator>Lestat</dc:creator>
            <pubDate>Thu, 31 Dec 2020 20:03:31 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Cold Subject - Crystal Clear (EP) (2020)</title>
            <guid isPermaLink="true">https://alterportal.net/2020_albums/170458-cold-subject-crystal-clear-ep-2020.html</guid>
            <link>https://alterportal.net/2020_albums/170458-cold-subject-crystal-clear-ep-2020.html</link>
            <description><![CDATA[<img src="https://funkyimg.com/i/39KkP.jpg" style="max-width:100%;" alt="Cold Subject - Crystal Clear (EP) (2020)"><br><b>Стиль: <span style="color:#E76363">Metalcore / Hard Rock</span></b><br><b>Треклист:</b><br>1. Hourglass<br>2. Fragile<br>3. Candlewax<br>4. Alice<br>5. New Life]]></description>
            <turbo:content><![CDATA[ <img src="https://funkyimg.com/i/39KkP.jpg" style="max-width:100%;" alt="Cold Subject - Crystal Clear (EP) (2020)"><br><b><a href="http://www.coldsubject.com" target="_blank" rel="noopener external noreferrer">Official Website</a> <span style="color:#00FFFF">|</span> <a href="https://www.facebook.com/ColdSubejct" target="_blank" rel="noopener external noreferrer">Facebook</a> <span style="color:#00FFFF">|</span> <a href="https://www.instagram.com/coldsubjectfl" target="_blank" rel="noopener external noreferrer">Instagram</a> <span style="color:#00FFFF">|</span> <a href="https://open.spotify.com/artist/0c25fpYwnfaUskPSPaOE9w" target="_blank" rel="noopener external noreferrer">Spotify</a> <span style="color:#00FFFF">|</span> <a href="https://music.apple.com/us/artist/cold-subject/1459827143" target="_blank" rel="noopener external noreferrer">Apple Music</a><br><br>Страна: <span style="color:#E76363">USA, Frorida, Orlando</span><br>Стиль: <span style="color:#E76363">Metalcore / Hard Rock</span> <br>Формат: <span style="color:#E76363">MP3 CBR 320 kbps</span><br>Лейбл: <span style="color:#E76363">Chilled Records</span><br>Дата релиза: <span style="color:#E76363">Oct 25, 2020</span></b><br><br><b>Треклист:</b><br>1. Hourglass<br>2. Fragile<br>3. Candlewax<br>4. Alice<br>5. New Life<br><br><iframe width="460" height="215" src="https://www.youtube.com/embed/48MP3ojoyqE" frameborder="0" allowfullscreen></iframe><br><br><b><a href="https://www84.zippyshare.com/v/6aYOuwdC/file.html" target="_blank" rel="noopener external noreferrer">Download</a></b> ]]></turbo:content>
            <category><![CDATA[Альбомы 2020        | Сore]]></category>
            <dc:creator>Lestat</dc:creator>
            <pubDate>Thu, 31 Dec 2020 20:03:09 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>С Новым 2021 Годом!</title>
            <guid isPermaLink="true">https://alterportal.net/raznoe/170461-s-novym-2021-godom.html</guid>
            <link>https://alterportal.net/raznoe/170461-s-novym-2021-godom.html</link>
            <description><![CDATA[<img src="https://i114.fastpic.ru/big/2020/1231/00/3bc17fa8e3eee9de2c9c787d5842f700.png" style="max-width:100%;" alt="С Новым 2021 Годом!"><br><span style="font-size:14pt;"><span style="color:#FF0000">Дорогие пользователи и гости нашего сайта!</span></span><br><br><span style="font-size:12pt;">Администрация сайта, поздравляет всех с наступающим Новым Годом!<br>Желаем всем крепкого здоровья, настоящего счастья, хорошего настроения и конечно же годных и крутых релизов в 2021  <img alt="yahoo" class="emoji" src="https://alterportal.net/engine/data/emoticons/yahoo.gif"> </span>]]></description>
            <turbo:content><![CDATA[ <div class="quote_single"><img src="https://i.ibb.co/zVBw6vH/99px-ru-animacii-32863-na-zimnij-gorod-opustilas-noch-v-domah-gorit.gif" style="max-width:100%;" alt="С Новым 2021 Годом!"><br><br><span style="font-size:14pt;">Отмечать желаю славно,<br>Не ударить в грязь лицом,<br>Запивать икру винцом,<br>Заедать коньяк тунцом,<br>Водку – супер-огурцом!<br>Захмелеть чтобы слегка,<br>Честь блюсти наверняка,<br>Встретить бодренько восход,<br>Как-никак же - Новый Год!</span><br><br><span style="font-size:12pt;">Наш любимый портал пережил сложный период восстановления, надеюсь, что со временем, сайт будет продолжать двигаться только вперед и постоянно радовать пользователей самыми свежими и годными релизами! </span><br><br><span style="color:#FF0000"><span style="font-size:18pt;">С Новым 2021 годом друзья! </span></span><br><br><span style="color:#636363">если музыка не звучит - нажми сюда </span><br><br><img src="https://i111.fastpic.ru/big/2019/1231/d8/181b92f0e91417d0edd748effa0028d8.png" style="max-width:100%;" alt=""><br><br><embed src="http://www.youtube.com/v/3Uo0JAUWijM?&amp;rel=1&amp;autoplay=1&amp;color1=0xb1b1b1&amp;color2=0x0000ff" type="application/x-shockwave-flash" width="25" height="25" allowscriptaccess="never" allownetworking="internal"></div> ]]></turbo:content>
            <category><![CDATA[Разное]]></category>
            <dc:creator>lollibub</dc:creator>
            <pubDate>Thu, 31 Dec 2020 19:55:20 +0300</pubDate>
        </item>
        <item turbo="true">
            <title>Anacondaz feat. Кис-Кис - Сядь Мне На Лицо</title>
            <guid isPermaLink="true">https://alterportal.net/video/170455-anacondaz-feat-kis-kis-sjad-mne-na-lico-2020.html</guid>
            <link>https://alterportal.net/video/170455-anacondaz-feat-kis-kis-sjad-mne-na-lico-2020.html</link>
            <description><![CDATA[<img src="https://a.radikal.ru/a02/2012/fa/46e601bbf851.jpg" style="max-width:100%;" alt="Anacondaz feat. Кис-Кис - Сядь Мне На Лицо"><br><b>Cтиль: Pop-Punk / Pop-Rock <br>Формат: MP4 1080p<br>Размер: 79 MB</b>]]></description>
            <turbo:content><![CDATA[ <img src="https://funkyimg.com/i/39KBs.jpg" style="max-width:100%;" alt="Anacondaz feat. Кис-Кис - Сядь Мне На Лицо"><br><br><b>Исполнитель: Anacondaz/Кис - Кис<br>Cтиль: Pop-Punk / Pop-Rock<br>Страна: Россия<br>Дата Релиза: 2020<br>Формат: MP4 1080p<br>Размер: 79 MB</b><br><br><a href="https://www40.zippyshare.com/v/Jd9NAW15/file.html" target="_blank" rel="noopener external noreferrer"><b>Cкачать!</b></a><br><br><center><iframe width="356" height="200" src="https://www.youtube.com/embed/nW33_DripcQ?feature=oembed" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></center> ]]></turbo:content>
            <category><![CDATA[Video          | ex-USSR          | Punk]]></category>
            <dc:creator>Wild.E.Coyote</dc:creator>
            <pubDate>Thu, 31 Dec 2020 19:11:02 +0300</pubDate>
        </item>
    </channel>
</rss>
`
