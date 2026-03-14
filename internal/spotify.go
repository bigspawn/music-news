package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-pkgz/lgr"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyApi struct {
	lgr    lgr.L
	client *spotify.Client
}

func NewSpotifyApi(lgr lgr.L, clientID, clientSecret string) (*SpotifyApi, error) {
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("spotify client ID and secret are required")
	}

	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}

	httpClient := config.Client(context.Background())
	client := spotify.New(httpClient, spotify.WithRetry(true))

	return &SpotifyApi{
		lgr:    lgr,
		client: client,
	}, nil
}

func (s *SpotifyApi) SearchAlbumURL(ctx context.Context, artist, album string) (string, error) {
	query := fmt.Sprintf("artist:%s album:%s", artist, album)

	results, err := s.client.Search(ctx, query, spotify.SearchTypeAlbum, spotify.Limit(5))
	if err != nil {
		return "", fmt.Errorf("spotify search: %w", err)
	}

	if results.Albums == nil || len(results.Albums.Albums) == 0 {
		return "", fmt.Errorf("no albums found in Spotify for %s - %s", artist, album)
	}

	searchTitle := strings.ToLower(clearTitle(artist + " - " + album))

	for _, a := range results.Albums.Albums {
		var artistName string
		if len(a.Artists) > 0 {
			artistName = a.Artists[0].Name
		}

		resultTitle := strings.ToLower(clearTitle(artistName + " - " + a.Name))
		dist := levenshteinDistance(searchTitle, resultTitle)
		if dist <= 15 {
			if url, ok := a.ExternalURLs["spotify"]; ok {
				return url, nil
			}
		}

		s.lgr.Logf("[DEBUG] Spotify match: %s vs %s, distance=%d", resultTitle, searchTitle, dist)
	}

	return "", fmt.Errorf("no matching album found in Spotify for %s - %s", artist, album)
}
