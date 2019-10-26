package main

import (
	"fmt"
	"github.com/mirrorsge/go_steam_api"
)

func main() {
	gsa := go_steam_api.New()
	loginURL := gsa.Authentication.GenLoginUrl("https://baidu.com")
	fmt.Println(loginURL)
	return
}
