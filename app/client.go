package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	url "net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	iTunesSearchUrl = "https://itunes.apple.com/search?"
	songLinksUrl    = "https://api.song.link/v1-alpha.1/links?"
)

var (
	country  = []string{"US"}
	entity   = []string{"album"}
	platform = []string{string(ItunesPlatform)}

	unusedSuffixRegexp = regexp.MustCompile(`\(\d+\)|\[EP]|\[[singleSINGLE]+]|\([singleSINGLE]+\)|(\sEP\s+)`)
)

type SearchResult struct {
	ResultCount int    `json:"resultCount"`
	Results     []Item `json:"results"`
}

type Item struct {
	ArtistId               int     `json:"artistId"`
	ArtistName             string  `json:"artistName"`
	ArtistViewUrl          string  `json:"artistViewUrl"`
	ArtworkUrl100          string  `json:"artworkUrl100"`
	ArtworkUrl30           string  `json:"artworkUrl30"`
	ArtworkUrl60           string  `json:"artworkUrl60"`
	CollectionCensoredName string  `json:"collectionCensoredName"`
	CollectionExplicitness string  `json:"collectionExplicitness"`
	CollectionId           int     `json:"collectionId"`
	CollectionName         string  `json:"collectionName"`
	CollectionPrice        float64 `json:"collectionPrice"`
	CollectionViewUrl      string  `json:"collectionViewUrl"`
	CollectionType         string  `json:"collectionType"`
	Country                string  `json:"country"`
	Currency               string  `json:"currency"`
	DiscCount              int     `json:"discCount"`
	DiscNumber             int     `json:"discNumber"`
	Kind                   string  `json:"kind"`
	PreviewUrl             string  `json:"previewUrl"`
	PrimaryGenreName       string  `json:"primaryGenreName"`
	RadioStationUrl        string  `json:"radioStationUrl"`
	ReleaseDate            string  `json:"releaseDate"`
	TrackCensoredName      string  `json:"trackCensoredName"`
	TrackCount             int     `json:"trackCount"`
	TrackExplicitness      string  `json:"trackExplicitness"`
	TrackId                int     `json:"trackId"`
	TrackName              string  `json:"trackName"`
	TrackNumber            int     `json:"trackNumber"`
	TrackPrice             float64 `json:"trackPrice"`
	TrackTimeMillis        int     `json:"trackTimeMillis"`
	TrackViewUrl           string  `json:"trackViewUrl"`
	WrapperType            string  `json:"wrapperType"`
}

type Platform string

const (
	SpotifyPlatform      Platform = "spotify"
	ItunesPlatform       Platform = "itunes"
	AppleMusicPlatform   Platform = "appleMusic"
	YoutubePlatform      Platform = "youtube"
	YoutubeMusicPlatform Platform = "youtubeMusic"
	GooglePlatform       Platform = "google"
	GoogleStorePlatform  Platform = "googleStore"
	PandoraPlatform      Platform = "pandora"
	DeezerPlatform       Platform = "deezer"
	TidalPlatform        Platform = "tidal"
	AmazonStorePlatform  Platform = "amazonStore"
	AmazonMusicPlatform  Platform = "amazonMusic"
	SoundcloudPlatform   Platform = "soundcloud"
	NapsterPlatform      Platform = "napster"
	YandexPlatform       Platform = "yandex"
	SpinrillaPlatform    Platform = "spinrilla"
)

type ApiProvider string

const (
	SpotifyProvider    ApiProvider = "spotify"
	ItunesProvider     ApiProvider = "itunes"
	YoutubeProvider    ApiProvider = "youtube"
	GoogleProvider     ApiProvider = "google"
	PandoraProvider    ApiProvider = "pandora"
	DeezerProvider     ApiProvider = "deezer"
	TidalProvider      ApiProvider = "tidal"
	AmazonProvider     ApiProvider = "amazon"
	SoundcloudProvider ApiProvider = "soundcloud"
	NapsterProvider    ApiProvider = "napster"
	YandexProvider     ApiProvider = "yandex"
	SpinrillaProvider  ApiProvider = "spinrilla"
)

type SongLinkResponse struct {
	EntityUniqueId     string            `json:"entityUniqueId"`
	UserCountry        string            `json:"userCountry"`
	PageUrl            string            `json:"pageUrl"`
	EntitiesByUniqueId map[string]Entity `json:"entitiesByUniqueId"`
	LinksByPlatform    map[Platform]Link `json:"linksByPlatform"`
	Platform           Platform          `json:"platform"`
}

type Entity struct {
	Type            string      `json:"type"`
	Title           string      `json:"title,omitempty"`
	ArtistName      string      `json:"artistName,omitempty"`
	ThumbnailUrl    string      `json:"thumbnailUrl,omitempty"`
	ThumbnailWidth  int         `json:"thumbnailWidth,omitempty"`
	ThumbnailHeight int         `json:"thumbnailHeight,omitempty"`
	ApiProvider     ApiProvider `json:"apiProvider"`
	Platforms       []Platform  `json:"platforms"`
}

type Link struct {
	EntityUniqueId      string `json:"entityUniqueId"`
	Url                 string `json:"url"`
	NativeAppUriMobile  string `json:"nativeAppUriMobile,omitempty"`
	NativeAppUriDesktop string `json:"nativeAppUriDesktop,omitempty"`
}

type LinksApi struct {
	client http.Client
	key    string
}

func NewLinkApi(key string) *LinksApi {
	return &LinksApi{
		client: http.Client{},
		key:    key,
	}
}

func (a *LinksApi) GetLink(ctx context.Context, title string) (string, error) {
	title = clearTitle(title)
	id, err := a.getIDiTunes(ctx, title)
	if err != nil {
		return "", err
	}
	return a.getLinkByiTunesID(ctx, id)
}

func (a *LinksApi) getIDiTunes(_ context.Context, title string) (string, error) {
	u := url.Values{
		"term":    []string{title},
		"country": country,
		"entity":  entity,
	}

	resp, err := a.client.Get(iTunesSearchUrl + u.Encode())
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	result := SearchResult{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	for _, item := range result.Results {
		resultName := strings.ToLower(fmt.Sprintf("%s - %s", item.ArtistName, item.CollectionName))
		if strings.Contains(resultName, strings.ToLower(title)) {
			return strconv.Itoa(item.CollectionId), nil
		}
	}

	Lgr.Logf("[INFO] iTunes response: %v", result)

	return "", fmt.Errorf("albums in iTunes not found: title=%s", title)
}

func (a *LinksApi) getLinkByiTunesID(_ context.Context, id string) (string, error) {
	u := url.Values{
		"platform":    platform,
		"type":        entity,
		"id":          []string{id},
		"key":         []string{a.key},
		"userCountry": country,
	}

	resp, err := a.client.Get(songLinksUrl + u.Encode())
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("bad response: %d %s", resp.StatusCode, resp.Status)
	}

	result := SongLinkResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.PageUrl == "" {
		return "", errors.New("songs link is empty")
	}

	return result.PageUrl, nil
}

func clearTitle(title string) string {
	title = unusedSuffixRegexp.ReplaceAllString(title, "")
	title = strings.TrimSpace(title)
	return title
}
