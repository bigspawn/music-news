package internal

import (
	"fmt"

	goOdesli "github.com/bigspawn/go-odesli"
)

var requiredPlatforms = []goOdesli.Platform{
	goOdesli.PlatformTidal,
	goOdesli.PlatformSpotify,
	goOdesli.PlatformItunes,
	goOdesli.PlatformYandex,
	goOdesli.PlatformDeezer,
}

func CheckRequiredPlatforms(platforms map[goOdesli.Platform]string) (map[goOdesli.Platform]string, error) {
	result := make(map[goOdesli.Platform]string)
	foundCount := 0

	for _, platform := range requiredPlatforms {
		if link, ok := platforms[platform]; ok {
			result[platform] = link
			foundCount++
		}
	}

	// Require at least one platform instead of mandatory iTunes and Spotify
	if foundCount == 0 {
		return nil, fmt.Errorf("no platform links found")
	}

	return result, nil
}
