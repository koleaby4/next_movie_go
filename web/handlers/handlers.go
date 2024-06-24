package handlers

import (
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/db"
)

type Handlers struct {
	Config  config.Config
	Queries *db.Queries
}

func New(cfg config.Config, queries *db.Queries) *Handlers {
	return &Handlers{
		Config:  cfg,
		Queries: queries,
	}
}
