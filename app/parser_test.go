package main

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	desc := "Leaked Release - February 25, 2020 Genre - Alternative Rock Quality - MP3, 320 kbps CBR &nbsp; Tracklist: 01. Run (4:00) 02. Girl (2:34) 03. Violence And Riots (3:31) 04. Soldiers (4:16) 05. Wayward Signs (3:43) 06. Lie To Me (3:17) 07. Rise (3:34) 08. Bubblegum (2:54) 09. Only One (3:12) 10. Sugar (3:34) 11. Shattered Suns (3:23) 12. I Need A Friend Tonight (3:50) 13. Between You and I (3:31) 14. Demons (3:55) &nbsp; Download &nbsp; Support! Facebook / iTunes &nbsp;"
	desc = split(desc)
	fmt.Printf("\n%s\n", desc)
}
