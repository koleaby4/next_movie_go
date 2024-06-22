package tmdb

import (
	"github.com/koleaby4/next_movie_go/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_expandPosterURL_expanded(t *testing.T) {
	m := db.Movie{PosterUrl: "/poster.jpg"}
	expandPosterURL(&m)
	expected := "https://image.tmdb.org/t/p/original/poster.jpg"
	assert.Equal(t, expected, m.PosterUrl)
}

func Test_expandPosterURL_nothingToExpand(t *testing.T) {
	m := db.Movie{}
	expandPosterURL(&m)
	assert.Empty(t, m.PosterUrl)
}
