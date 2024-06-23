package handlers

import (
	"github.com/koleaby4/next_movie_go"
)

// Handlers is the struct for the handlers
type Handlers struct {
	AppConfig next_movie_go.Config
}

// New creates a new Handlers struct
func New(cfg next_movie_go.Config) *Handlers {
	return &Handlers{
		AppConfig: cfg,
	}
}
