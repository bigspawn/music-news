package api

import (
	"encoding/json"
	"errors"
	"net/http"
	url2 "net/url"
	"strconv"
)

const (
	country = "US"
	entity  = "song,album,podcast"
)

var ErrITunesNotFound = errors.New("no results from itunes")

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
		return nil, ErrITunesNotFound
	}

	v := url2.Values{}
	v.Set("platform", string(ItunesPlatform))
	v.Set("type", "album")
	v.Set("id", strconv.Itoa(search.Results[0].CollectionId))
	v.Set("key", key)

	reqUrl := "https://api.song.link/v1-alpha.1/links?" + v.Encode()
	res, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	links := &SongLinkResponse{}
	err = json.NewDecoder(res.Body).Decode(links)
	if err != nil {
		return nil, err
	}
	return links, nil
}
