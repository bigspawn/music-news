package internal

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	itunes "github.com/bigspawn/go-itunes-api"
	odesli "github.com/bigspawn/go-odesli"
	"github.com/go-pkgz/lgr"
)

var unusedSuffixRegexp = regexp.MustCompile(`\(\d+\)|([\[(][singleSINGLEpP]+[])])|(\s+-\sEP$)|(\s+\[EP\]$)`)

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
	Lgr    lgr.L
	Itunes itunes.API
	Odesli odesli.API
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
	id, err := api.getIDiTunes(ctx, clearTitle(title))
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
	return resp.PageUrl, links, nil
}

func (api *LinksApi) getIDiTunes(ctx context.Context, title string) (string, error) {
	resp, err := api.Itunes.Search(ctx, itunes.SearchRequest{
		Term:    title,
		Country: itunes.US,
	})
	if err != nil {
		return "", err
	}

	if len(resp.Results.Results) == 0 {
		return "", fmt.Errorf("albums in iTunes not found: title=%s", title)
	}

	id, err := findCollectionIDFromResultsByTitle(api.Lgr, resp.Results.Results, title)
	if err != nil {
		api.Lgr.Logf("[INFO] iTunes response: %v", resp)
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
		return nil, err
	}
	return &resp, nil
}

func clearTitle(title string) string {
	title = unusedSuffixRegexp.ReplaceAllString(title, "")
	title = strings.TrimSpace(title)
	return title
}

func findCollectionIDFromResultsByTitle(l lgr.L, results []itunes.Result, title string) (string, error) {
	title = strings.ToLower(title)
	title = clearTitle(title)
	for _, item := range results {
		if item.Kind != itunes.KindAlbum {
			continue
		}
		t := item.ArtistName + " - " + item.CollectionName
		t = clearTitle(t)
		t = strings.ToLower(t)
		n := levenshteinDistance(t, title)
		if n <= 10 {
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
