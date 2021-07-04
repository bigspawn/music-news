package internal

import "strings"

var skipGenders = []string{
	"Blues Rock",
	"Celtic Punk",
	"City Pop",
	"Country",
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

func isSkippedGender(data string) bool {
	if len(data) == 0 {
		return false
	}

	for i := range skipGenders {
		if containsAny(data, skipGenders[i], strings.ToLower(skipGenders[i])) {
			return true
		}
	}

	return false
}

func containsAny(s string, substr ...string) bool {
	for _, ss := range substr {
		if strings.Contains(s, ss) {
			return true
		}
	}
	return false
}
