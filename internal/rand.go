package internal

import "math/rand"

func RandBetween(max int, min int) int {
	//nolint:gosec // no crypto
	return rand.Intn(max-min+1) + min
}
