package go_steam_api

import (
	"github.com/mirrorsge/go_steam_api/authentication"
)

type GlobalConfig struct {
}




type Engine struct {
	Config GlobalConfig
	Authentication *authentication.Authentication
}

func New() *Engine {
	newEngine := &Engine{}
	return newEngine
}




