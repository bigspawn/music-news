package internal

import (
	"strings"

	"github.com/go-pkgz/lgr"
)

var skipGenders = []string{
	"Blues Rock",
	"Celtic Punk",
	"City Pop",
	"Cow Punk",
	"Dance",
	"Death 'n' Roll",
	"Deutschpunk",
	"Disco",
	"Eurodance",
	"Gothabilly",
	"Hillbilly",
	"Hip-Hop",
	"House",
	"Oi!",
	"Pop Music",
	"Pop Rock",
	"Pop",
	"Psychobilly",
	"Punkabilly",
	"R&B",
	"R'n'B",
	"Rap / Hip Hop",
	"Retro Pop",
	"Rockabilly",
	"Skate Punk",
	"Street Punk",
}

func isSkippedGender(l lgr.L, data string) bool {
	if len(data) == 0 {
		return false
	}

	for i := range skipGenders {
		if ok, reason := containsAny(data, skipGenders[i], strings.ToLower(skipGenders[i])); ok {
			l.Logf("%s is skipped by %s", data, reason)
			return ok
		}
	}

	return false
}

func containsAny(s string, substr ...string) (bool, string) {
	for _, ss := range substr {
		if strings.Contains(s, ss) {
			return true, ss
		}
	}

	return false, ""
}
