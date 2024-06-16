package handlers

import (
	"github.com/koleaby4/next_movie_go"
)

type Handlers struct {
	AppConfig next_movie_go.AppConfig
}

func New(cfg next_movie_go.AppConfig) *Handlers {
	return &Handlers{
		AppConfig: cfg,
	}
}
