package internal

import "strings"

var skipGenders = []string{
	"Retro Pop", "R&B", "Pop Music", "Pop Rock", "City Pop", "Disco", "Eurodance", "Hip-Hop",
	"retro pop", "r&b", "pop music", "pop rock", "city pop", "disco", "eurodance", "hip-hop",
	"Pop", "pop", "Rap / Hip Hop", "R'n'B", "Dance / Electronic / House",
	"Rockabilly", "rockabilly",
	"Punkabilly", "punkabilly",
	"Psychobilly", "psychobilly",
	"Street Punk", "street punk",
	"Blues Rock", "blues rock",
	"Country / Cow Punk / Hillbilly", "Cow Punk",
	"Hillbilly", "hillbilly",
	"Gothabilly", "gothabilly",
	"Death 'n' Roll",
	"Country",
	"Oi! / Punk Rock",
	"Deutschpunk",
	"Skate Punk",
}

func isSkippedGender(data string) bool {
	for _, s := range skipGenders {
		if strings.Contains(data, s) || strings.Contains(data, s) {
			return true
		}
	}
	return false
}
