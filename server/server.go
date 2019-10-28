package main

import (
	"fmt"
	"github.com/mirrorsge/go_steam_api"
	"log"
	"net/http"
	"strings"
)

func receiveSteamOpenIDValidation(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()   //解析参数，默认是不会解析的
	openIDParams := map[string]string{}
	for k, v := range r.Form {
		value := strings.Join(v, "")
		openIDParams[k] = value
	}
	gsa := go_steam_api.New()
	steamID, err := gsa.Authentication.ValidateOpenID(openIDParams)
	if err != nil {
		_, _ = fmt.Fprint(w, err)
	}
	_, _ = fmt.Fprintf(w, "steamID is: %s", steamID)
	return
}

func main(){
	http.HandleFunc("/steamOpenIDValidation",receiveSteamOpenIDValidation)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}}
