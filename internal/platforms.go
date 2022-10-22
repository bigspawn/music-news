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
}

func CheckRequiredPlatforms(platforms map[goOdesli.Platform]string) (map[goOdesli.Platform]string, error) {
	result := make(map[goOdesli.Platform]string)
	for _, platform := range requiredPlatforms {
		if link, ok := platforms[platform]; ok {
			result[platform] = link
			continue
		}
		if platform == goOdesli.PlatformTidal || platform == goOdesli.PlatformSpotify || platform == goOdesli.PlatformItunes {
			return nil, fmt.Errorf("link for platform=%s not found", platform)
		}
	}
	return result, nil
}
