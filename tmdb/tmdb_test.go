package tmdb

import (
	"github.com/koleaby4/next_movie_go/db"
	"testing"
)

func Test_expandPosterURL_expanded(t *testing.T) {
	m := db.Movie{PosterUrl: "/poster.jpg"}
	expandPosterURL(&m)
	expected := "https://image.tmdb.org/t/p/original/poster.jpg"
	if m.PosterUrl != expected {
		t.Error("Expected:", expected, " got:", m.PosterUrl)
	}
}

func Test_expandPosterURL_nothingToExpand(t *testing.T) {
	m := db.Movie{}
	expandPosterURL(&m)
	if m.PosterUrl != "" {
		t.Error("Expected empty PosterURL, got", m.PosterUrl)
	}
}
