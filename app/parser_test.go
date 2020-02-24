package main

import (
	"strings"
	"testing"
)

func Test_test(t *testing.T) {
	info := "Leaked Release - February 25, 2020 Genre - Progressive Deathcore Quality - MP3, 320 kbps CBR &nbsp; Tracklist: 01. Athenas (5:51) &nbsp; Download &nbsp; Support! Facebook / iTunes"
	t.Log(info)
	info = strings.ReplaceAll(info, "&nbsp;", "\n")
	t.Log(info)
}
