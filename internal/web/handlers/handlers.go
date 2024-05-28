package handlers

import (
	"github.com/koleaby4/next_movie_go/config"
)

type Handlers struct {
	AppConfig config.AppConfig
}

func New(cfg config.AppConfig) *Handlers {
	return &Handlers{
		AppConfig: cfg,
	}
}
