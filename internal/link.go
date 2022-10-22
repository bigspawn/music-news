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

var unusedSuffixRegexp = regexp.MustCompile(`\(\d+\)|([\[(][singleSINGLEpP]+[])])`)

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
		return "", nil, err
	}

	resp, err := api.GetSongLink(ctx, id)
	if err != nil {
		return "", nil, err
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
	id, err := getID(resp.Results.Results, title)
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

func getID(results []itunes.Result, title string) (string, error) {
	lowerTitle := strings.ToLower(title)
	count := 0
	for _, item := range results {
		var tokens []string
		tokens = append(tokens,
			append(
				strings.Split(strings.ToLower(item.ArtistName), " "),
				strings.Split(strings.ToLower(item.CollectionName), " ")...,
			)...,
		)
		for _, token := range tokens {
			if strings.Contains(lowerTitle, token) {
				count++
			}
		}
		percent := float64(count) / float64(len(tokens)) * float64(100)
		if percent >= 65 {
			return strconv.Itoa(item.CollectionId), nil
		}
	}

	return "", fmt.Errorf("albums in iTunes not found: title=%s", title)
}
