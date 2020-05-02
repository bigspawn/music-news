package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	country = "US"
	entity  = "song,album,podcast"
)

var (
	ErriTunesNotFound    = errors.New("no results from itunes")
	ErriSongLinkNotFound = errors.New("no results from song links")
	ErrWrongTitle        = errors.New("wrong title")
)

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

func GetByTitle(title, key string) (*SongLinkResponse, error) {
	search, err := Search(title, country, entity)
	if err != nil {
		return nil, err
	}

	if search.ResultCount == 0 {
		return nil, ErriTunesNotFound
	}

	splits := strings.Split(strings.ToLower(title), " - ")
	if len(splits) != 2 {
		return nil, ErrWrongTitle
	}
	artist := splits[0]
	album := splits[1]

	var result SearchMap
	var isSearchResultsContainNeedsResult bool
	for _, result = range search.Results {
		if result.WrapperType == "collection" &&
			strings.Contains(strings.ToLower(result.ArtistName), artist) &&
			strings.Contains(strings.ToLower(result.CollectionName), album) {
			isSearchResultsContainNeedsResult = true
			break
		}
	}

	if !isSearchResultsContainNeedsResult {
		return nil, ErriTunesNotFound
	}

	v := url.Values{}
	v.Set("platform", string(ItunesPlatform))
	v.Set("type", "album")
	v.Set("id", strconv.Itoa(result.CollectionId))
	v.Set("key", key)
	v.Set("userCountry", "US")

	reqUrl := "https://api.song.link/v1-alpha.1/links?" + v.Encode()
	res, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != 200 {
		return nil, ErriSongLinkNotFound
	}

	response := &SongLinkResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}
	if response.PageUrl == "" {
		return nil, errors.New("songs link is empty")
	}
	return response, nil
}
