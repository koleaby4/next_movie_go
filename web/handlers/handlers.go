package handlers

import (
	"github.com/koleaby4/next_movie_go"
)

// Handlers is the struct for the handlers
type Handlers struct {
	AppConfig next_movie_go.AppConfig
}

// New creates a new Handlers struct
func New(cfg next_movie_go.AppConfig) *Handlers {
	return &Handlers{
		AppConfig: cfg,
	}
}
