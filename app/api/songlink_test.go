package api

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	l, err := Get("Bring Me the Horizon", "That's the Spirit ")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("-------")
	fmt.Println(l)
}
