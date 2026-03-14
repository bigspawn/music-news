package internal

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	itunes "github.com/bigspawn/go-itunes-api"
	odesli "github.com/bigspawn/go-odesli"
	"github.com/go-pkgz/lgr"
)

var unusedSuffixRegexp = regexp.MustCompile(`(?i)` +
	`\(\d{4}\)` + // (2020)
	`|[\[(]single[)\]]` + // [Single] or (Single)
	`|\s+\[EP\]` + // [EP]
	`|\(EP\)` + // (EP)
	`|\s+-\s+EP(?:\s|$)` + // - EP (at end or before space)
	`|\s+-\s+Single(?:\s|$)` + // - Single (iTunes format)
	`|\(Bonus Track Version\)` +
	`|\(Deluxe Version\)` +
	`|\(Deluxe Edition\)` +
	`|\(Deluxe\)` +
	`|\[Deluxe Edition\]` +
	`|\[Deluxe\]` +
	`|\[Deluxe Version\]` +
	`|\(Remastered\)` +
	`|\[Remastered\]` +
	`|\(Remaster\)` +
	`|\(Special Edition\)` +
	`|\[Special Edition\]` +
	`|\(Expanded Edition\)` +
	`|\[Expanded Edition\]` +
	`|\[Self-titled\]` + // [Self-titled]
	`|\(Remixed\s*&\s*Remastered\s*\d{4}\)` + // (Remixed & Remastered 2026)
	`|\(\d{4}\s+Remixed\s+and\s+Remastered\s+Version\)` + // (2026 Remixed and Remastered Version)
	`|\[DJ Mix\]` + // [DJ Mix]
	`|\(feat\.[^)]*\)` + // (feat. Artist Name)
	`|\s+\([A-Z]{2}\)`, // (US), (UK) country codes
)

type LinksApiParams struct {
	Lgr          lgr.L
	ITunesClient itunes.API
	OdesliClient odesli.API
}

func (p *LinksApiParams) Validate() error {
	if p.Lgr == nil {
		return fmt.Errorf("lgr is required")
	}
	if p.ITunesClient == nil {
		return fmt.Errorf("itunes client is required")
	}
	if p.OdesliClient == nil {
		return fmt.Errorf("odesli client is required")
	}
	return nil
}

type LinksApi struct {
	Lgr     lgr.L
	Itunes  itunes.API
	Odesli  odesli.API
	Spotify *SpotifyApi
}

func NewLinksApi(params LinksApiParams) (*LinksApi, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &LinksApi{
		Lgr:    params.Lgr,
		Itunes: params.ITunesClient,
		Odesli: params.OdesliClient,
	}, nil
}

func (api *LinksApi) GetLinks(ctx context.Context, title string) (string, map[odesli.Platform]string, error) {
	searchTitle := clearTitle(title)

	id, err := api.getIDiTunes(ctx, searchTitle)
	if err != nil {
		return "", nil, fmt.Errorf("get itunes id for %s: %w", title, err)
	}

	resp, err := api.GetSongLink(ctx, id)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get links for %s: %w", title, err)
	}

	links := make(map[odesli.Platform]string, len(resp.LinksByPlatform))
	for p, l := range resp.LinksByPlatform {
		if l.Url != "" {
			links[p] = l.Url
		}
	}

	if _, hasSpotify := links[odesli.PlatformSpotify]; !hasSpotify && api.Spotify != nil {
		artist, album := splitArtistAlbum(searchTitle)
		if artist != "" && album != "" {
			spotifyURL, sErr := api.Spotify.SearchAlbumURL(ctx, artist, album)
			if sErr != nil {
				api.Lgr.Logf("[DEBUG] Spotify fallback search failed for %s: %v", title, sErr)
			} else if spotifyURL != "" {
				links[odesli.PlatformSpotify] = spotifyURL
				api.Lgr.Logf("[INFO] Spotify link found via fallback for %s", title)
			}
		}
	}

	return resp.PageUrl, links, nil
}

// fallbackCountries defines the order of iTunes store countries to search.
// US has the largest catalog, GB is the second largest English-speaking store,
// DE covers German metal scene, SE covers Scandinavian metal/rock.
var fallbackCountries = []string{"US", "GB", "DE", "SE"}

func (api *LinksApi) getIDiTunes(ctx context.Context, title string) (string, error) {
	var lastErr error
	for i, country := range fallbackCountries {
		if i > 0 {
			time.Sleep(3 * time.Second)
		}
		id, err := api.searchITunesCountry(ctx, title, country)
		if err == nil {
			return id, nil
		}
		lastErr = err
		api.Lgr.Logf("[DEBUG] iTunes search failed for country=%s, title=%s: %v", country, title, err)
	}
	return "", fmt.Errorf("search iTunes in all countries: %w", lastErr)
}

func (api *LinksApi) searchITunesCountry(ctx context.Context, title, country string) (string, error) {
	resp, err := api.Itunes.Search(ctx, itunes.SearchRequest{
		Term:    title,
		Country: country,
		Entity:  itunes.EntityAlbum,
		Limit:   10,
		Media:   itunes.MediaTypeMusic,
	})
	if err != nil {
		return "", fmt.Errorf("search iTunes country=%s: %w", country, err)
	}

	if len(resp.Results.Results) == 0 {
		return "", fmt.Errorf("no results from iTunes for %s (country=%s)", title, country)
	}

	id, err := findCollectionIDFromResultsByTitle(api.Lgr, resp.Results.Results, title)
	if err != nil {
		api.Lgr.Logf("[INFO] iTunes response (country=%s): %v", country, resp)
		return "", err
	}
	return id, nil
}

func (api *LinksApi) GetSongLink(ctx context.Context, id string) (*odesli.GetLinksResponse, error) {
	resp, err := api.Odesli.GetLinks(ctx, odesli.GetLinksRequest{
		ID:          id,
		UserCountry: "US",
		Platform:    odesli.PlatformItunes,
		Type:        odesli.EntityTypeAlbum,
	})
	if err != nil {
		return nil, fmt.Errorf("get links from odesli: %w", err)
	}
	return &resp, nil
}

var multipleSpacesRegexp = regexp.MustCompile(`\s{2,}`)

func clearTitle(title string) string {
	title = strings.ReplaceAll(title, "—", "-")
	title = strings.ReplaceAll(title, "\u2013", "-") // en-dash
	title = unusedSuffixRegexp.ReplaceAllString(title, "")
	title = multipleSpacesRegexp.ReplaceAllString(title, " ")
	title = strings.TrimSpace(title)
	return title
}

func findCollectionIDFromResultsByTitle(l lgr.L, results []itunes.Result, title string) (string, error) {
	title = strings.ToLower(title)
	title = clearTitle(title)
	for _, item := range results {
		// if item.Kind != itunes.KindAlbum {
		// continue
		// }
		t := item.ArtistName + " - " + item.CollectionName
		t = clearTitle(t)
		t = strings.ToLower(t)
		n := levenshteinDistance(t, title)
		if n <= 15 {
			return strconv.Itoa(item.CollectionId), nil
		}
		l.Logf("[INFO] levenshteinDistance(%s, %s) = %d", t, title, n)
	}
	return "", fmt.Errorf("albums in iTunes not found: title=%s", title)
}

func levenshteinDistance(s1, s2 string) int {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	if len(runes1) == 0 {
		return len(runes2)
	}
	if len(runes2) == 0 {
		return len(runes1)
	}

	matrix := make([][]int, len(runes1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(runes2)+1)
	}

	for i := 0; i <= len(runes1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(runes2); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(runes1); i++ {
		for j := 1; j <= len(runes2); j++ {
			cost := 1
			if runes1[i-1] == runes2[j-1] {
				cost = 0
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,
				min(
					matrix[i][j-1]+1,
					matrix[i-1][j-1]+cost,
				),
			)
		}
	}

	return matrix[len(runes1)][len(runes2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitArtistAlbum(title string) (artist, album string) {
	parts := strings.SplitN(title, " - ", 2)
	if len(parts) != 2 {
		return "", ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}
