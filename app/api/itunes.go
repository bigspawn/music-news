package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
)

type SearchMap struct {
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

type SearchResult struct {
	ResultCount int         `json:"resultCount"`
	Results     []SearchMap `json:"results"`
}

type LookupMap struct {
	Advisories                       []string `json:"advisories"`
	AppletvScreenshotUrls            []string `json:"appletvScreenshotUrls"`
	ArtistId                         int      `json:"artistId"`
	ArtistName                       string   `json:"artistName"`
	ArtistViewUrl                    string   `json:"artistViewUrl"`
	ArtworkUrl60                     string   `json:"artworkUrl60"`
	ArtworkUrl100                    string   `json:"artworkUrl100"`
	ArtworkUrl512                    string   `json:"artworkUrl512"`
	BundleId                         string   `json:"bundleId"`
	ContentAdvisoryRating            string   `json:"contentAdvisoryRating"`
	Currency                         string   `json:"currency"`
	CurrentVersionReleaseDate        string   `json:"currentVersionReleaseDate"`
	Description                      string   `json:"description"`
	Features                         []string `json:"features"`
	FileSizeBytes                    string   `json:"fileSizeBytes"`
	FormattedPrice                   string   `json:"formattedPrice"`
	Genres                           []string `json:"genres"`
	GenreIds                         []string `json:"genreIds"`
	IpadScreenshotUrls               []string `json:"ipadScreenshotUrls"`
	IsGameCenterEnabled              bool     `json:"isGameCenterEnabled"`
	IsVppDeviceBasedLicensingEnabled bool     `json:"isVppDeviceBasedLicensingEnabled"`
	Kind                             string   `json:"kind"`
	LanguageCodesISO2A               []string `json:"languageCodesISO2A"`
	MinimumOsVersion                 string   `json:"minimumOsVersion"`
	Price                            float64  `json:"price"`
	PrimaryGenreId                   int      `json:"primaryGenreId"`
	PrimaryGenreName                 string   `json:"primaryGenreName"`
	ReleaseDate                      string   `json:"releaseDate"`
	ReleaseNotes                     string   `json:"releaseNotes"`
	ScreenshotUrls                   []string `json:"screenshotUrls"`
	SellerName                       string   `json:"sellerName"`
	SellerUrl                        string   `json:"sellerUrl"`
	SupportedDevices                 []string `json:"supportedDevices"`
	TrackCensoredName                string   `json:"trackCensoredName"`
	TrackContentRating               string   `json:"trackContentRating"`
	TrackId                          int      `json:"trackId"`
	TrackName                        string   `json:"trackName"`
	TrackViewUrl                     string   `json:"trackViewUrl"`
	Version                          string   `json:"version"`
	WrapperType                      string   `json:"wrapperType"`
}

type LookupResult struct {
	ResultCount int         `json:"resultCount"`
	Results     []LookupMap `json:"results"`
}

func Search(query, country, entity string) (*SearchResult, error) {
	u := url2.Values{}
	u["term"] = []string{query}
	u["country"] = []string{country}
	u["entity"] = []string{entity}
	res, err := http.Get("https://itunes.apple.com/search?" + u.Encode())
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("[ERROR] close response body: %v\n", err)
		}
	}()
	result := SearchResult{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Lookup(searchTerm string, searchTermValue string, entity string, limit string, sort string) (*LookupResult, error) {
	u := url2.Values{}
	u[searchTerm] = []string{searchTermValue}
	u["entity"] = []string{entity}
	u["limit"] = []string{limit}
	u["sort"] = []string{sort}
	res, err := http.Get("https://itunes.apple.com/lookup?" + u.Encode())
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("[ERROR] close response body: %v\n", err)
		}
	}()
	result := LookupResult{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
