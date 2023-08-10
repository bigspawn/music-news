package internal

import (
	"regexp"

	"github.com/go-pkgz/lgr"
)

var genreToSkip = regexp.MustCompile(
	`^Genre:\s*\w*(` +
		`Blues Rock|` +
		`Celtic Punk|` +
		`City Pop|` +
		`Cow Punk|` +
		`Dance|` +
		`Death 'n' Roll|` +
		`Deutschpunk|` +
		`Disco|` +
		`Eurodance|` +
		`Gothabilly|` +
		`Hillbilly|` +
		`Hip-Hop|` +
		`Oi!|` +
		`Pop Music|` +
		`Pop Rock|` +
		`Psychobilly|` +
		`Punkabilly|` +
		`R&B|` +
		`R'n'B|` +
		`Rap \/ Hip Hop|` +
		`Retro Pop|` +
		`Rockabilly|` +
		`Skate Punk|` +
		`Street Punk|` +
		`Punk 'N' Roll).*$`,
)

func isSkippedGenre(l lgr.L, data string) bool {
	if len(data) == 0 {
		return false
	}
	if genreToSkip.MatchString(data) {
		l.Logf("skipped by gender: %s", data)
		return true
	}
	return false
}
