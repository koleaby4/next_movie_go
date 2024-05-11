package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Movie struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
	Rating      float64 `json:"vote_average"`
	PosterURL   string  `json:"poster_path"`
	TrailerURL  *string `json:"trailer_path"`
}
type TMDbResults struct {
	Results []Movie `json:"results"`
}

func GetRecentMovies(apiKey string, minRating float64) ([]Movie, error) {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	url := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?primary_release_date.gte=%v&vote_average.gte=%v&sort_by=release_date.desc&api_key=%v", sevenDaysAgo, minRating, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tmdbResponse TMDbResults
	err = json.Unmarshal(body, &tmdbResponse)
	if err != nil {
		return nil, err
	}

	return tmdbResponse.Results, nil
}
