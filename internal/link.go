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

var (
	unusedSuffixRegexp = regexp.MustCompile(`\(\d+\)|([\[(][singleSINGLEpP]+[])])|(\s+-\sEP$)|(\s+\[EP\]$)|(\s+\[[^\]]*EP[^\]]*\])|(\s+\[[^\]]*CD[^\]]*\])`)
	splitEPRegexp      = regexp.MustCompile(`(?i)\s*\[split\s+EP\]\s*`)
	cdMarkingRegexp    = regexp.MustCompile(`(?i)\s*\[[0-9]*CD\]\s*`)
	cyrillicYearRegexp = regexp.MustCompile(`\s*\(\d{4}\)\s*$`)
	discographyRegexp  = regexp.MustCompile(`(?i)\s*-?\s*(дискография|discography)\s*$`)
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
	cleanedTitle := clearTitle(title)
	variants := generateSearchVariants(cleanedTitle)

	api.Lgr.Logf("[DEBUG] trying search variants for title [%s]: %v", title, variants)

	// Пробуем каждый вариант поиска
	for i, variant := range variants {
		api.Lgr.Logf("[DEBUG] trying variant %d: %s", i+1, variant)

		resp, err := api.Itunes.Search(ctx, itunes.SearchRequest{
			Term:    variant,
			Country: itunes.US,
		})
		if err != nil {
			api.Lgr.Logf("[WARN] iTunes search failed for variant [%s]: %v", variant, err)
			continue
		}

		if len(resp.Results.Results) == 0 {
			api.Lgr.Logf("[DEBUG] no results for variant: %s", variant)
			continue
		}

		api.Lgr.Logf("[DEBUG] found %d results for variant [%s]", len(resp.Results.Results), variant)

		// Пробуем с разными порогами distance
		thresholds := []int{5, 10, 15, 20}
		for _, threshold := range thresholds {
			id, err := findCollectionIDFromResultsByTitleWithThreshold(api.Lgr, resp.Results.Results, cleanedTitle, threshold)
			if err == nil {
				api.Lgr.Logf("[INFO] found match with threshold %d for variant [%s]: %s", threshold, variant, id)
				return id, nil
			}
		}

		// Если не нашли точное совпадение, логируем результаты для диагностики
		api.Lgr.Logf("[INFO] iTunes response for variant [%s]: %d results", variant, len(resp.Results.Results))
		for j, result := range resp.Results.Results {
			if j >= 3 { // Логируем только первые 3 результата
				break
			}
			api.Lgr.Logf("[DEBUG] result %d: Artist=%s, Album=%s, Kind=%s", j+1, result.ArtistName, result.CollectionName, result.Kind)
		}
	}

	return "", fmt.Errorf("albums in iTunes not found after trying all variants: title=%s", title)
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
	// Удаляем года в скобках
	title = cyrillicYearRegexp.ReplaceAllString(title, "")
	// Удаляем слово "дискография"
	title = discographyRegexp.ReplaceAllString(title, "")
	// Удаляем split EP маркировки
	title = splitEPRegexp.ReplaceAllString(title, "")
	// Удаляем CD маркировки
	title = cdMarkingRegexp.ReplaceAllString(title, "")
	// Удаляем остальные суффиксы
	title = unusedSuffixRegexp.ReplaceAllString(title, "")
	// Убираем лишние пробелы
	title = strings.TrimSpace(title)
	return title
}

// Генерирует варианты поиска для fallback стратегий
func generateSearchVariants(title string) []string {
	variants := []string{title}

	// Вариант без года
	withoutYear := cyrillicYearRegexp.ReplaceAllString(title, "")
	withoutYear = strings.TrimSpace(withoutYear)
	if withoutYear != title && withoutYear != "" {
		variants = append(variants, withoutYear)
	}

	// Вариант без дискографии
	withoutDiscography := discographyRegexp.ReplaceAllString(title, "")
	withoutDiscography = strings.TrimSpace(withoutDiscography)
	if withoutDiscography != title && withoutDiscography != "" {
		variants = append(variants, withoutDiscography)
	}

	// Если есть " - ", попробуем только имя артиста
	if parts := strings.Split(title, " - "); len(parts) >= 2 {
		artistOnly := strings.TrimSpace(parts[0])
		if artistOnly != "" {
			variants = append(variants, artistOnly)
		}
	}

	return variants
}

func findCollectionIDFromResultsByTitle(l lgr.L, results []itunes.Result, title string) (string, error) {
	return findCollectionIDFromResultsByTitleWithThreshold(l, results, title, 10)
}

func findCollectionIDFromResultsByTitleWithThreshold(l lgr.L, results []itunes.Result, title string, threshold int) (string, error) {
	title = strings.ToLower(title)
	title = clearTitle(title)

	bestMatch := ""
	bestDistance := threshold + 1

	for _, item := range results {
		if item.Kind != itunes.KindAlbum {
			continue
		}

		// Пробуем разные комбинации
		variants := []string{
			item.ArtistName + " - " + item.CollectionName,
			item.CollectionName,
			item.ArtistName,
		}

		for _, variant := range variants {
			t := clearTitle(variant)
			t = strings.ToLower(t)
			distance := levenshteinDistance(t, title)

			l.Logf("[DEBUG] comparing [%s] with [%s] = distance %d", t, title, distance)

			if distance <= threshold && distance < bestDistance {
				bestDistance = distance
				bestMatch = strconv.Itoa(item.CollectionId)
				l.Logf("[DEBUG] new best match: distance=%d, id=%s", distance, bestMatch)
			}
		}
	}

	if bestMatch != "" {
		l.Logf("[INFO] found best match with distance %d", bestDistance)
		return bestMatch, nil
	}

	return "", fmt.Errorf("albums in iTunes not found with threshold %d: title=%s", threshold, title)
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
