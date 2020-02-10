package main

import (
	"fmt"

	"github.com/bigspawn/music-news/api"
)

func main() {
	l, err := api.Get("Bring Me the Horizon", "That's the Spirit ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("-------")
	fmt.Println(l)
}
