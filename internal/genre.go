package internal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-pkgz/lgr"
)

var genres = []string{
	`Blues Rock`,
	`Blues Rock`,
	`Celtic Punk`,
	`City Pop`,
	`Cow Punk`,
	`Dance`,
	`Death 'n' Roll`,
	`Deutschpunk`,
	`Disco`,
	`Eurodance`,
	`Gothabilly`,
	`Hillbilly`,
	`Hip-Hop`,
	`Oi!`,
	`Pop Music`,
	`Pop Punk`,
	`Pop Rock`,
	`Psychobilly`,
	`Punk 'N' Roll`,
	`Punkabilly`,
	`R&B`,
	`R'n'B`,
	`RnB`,
	`Rap \/ Hip Hop`,
	`Retro Pop`,
	`Rockabilly`,
	`Skate Punk`,
	`Street Punk`,
}

var genreRegexpStr = fmt.Sprintf(`Genre:.*(%s).*\n`, strings.Join(genres, "|"))

var genreToSkip = regexp.MustCompile(genreRegexpStr)

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
