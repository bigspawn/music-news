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
	for _, platform := range requiredPlatforms {
		if link, ok := platforms[platform]; ok {
			result[platform] = link
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no required platforms found")
	}
	return result, nil
}
