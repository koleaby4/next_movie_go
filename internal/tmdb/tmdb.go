package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Movie struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
	Rating      float64 `json:"vote_average"`
	PosterURL   string  `json:"poster_path"`
	TrailerURL  string  `json:"trailer_path"`
}

type MovieSearchResults struct {
	Movies []Movie `json:"results"`
}

type VideoResult struct {
	Results []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
	} `json:"results"`
}

var baseUrl = "https://api.themoviedb.org"
var imageBaseUrl = "https://image.tmdb.org/t/p/original"

func GetRecentMovies(apiKey string, minRating float64) ([]Movie, error) {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	url := fmt.Sprintf("%s/3/discover/movie?primary_release_date.gte=%v&vote_average.gte=%v&sort_by=release_date.desc&api_key=%v", baseUrl, sevenDaysAgo, minRating, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results MovieSearchResults
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}

	return results.Movies, nil
}

func EnrichMovies(apiKey string, movies []Movie) []Movie {
	wg := sync.WaitGroup{}
	wg.Add(len(movies))

	for i := range movies {
		movie := &movies[i]
		go addTrailerUrl(apiKey, movie, &wg)
	}
	wg.Wait()

	return movies
}

func addTrailerUrl(apiKey string, m *Movie, wg *sync.WaitGroup) {
	defer wg.Done()

	videoUrl := fmt.Sprintf("%s/3/movie/%d/videos?api_key=%s", baseUrl, m.Id, apiKey)
	videoResp, err := http.Get(videoUrl)
	if err != nil {
		log.Println("error fetching video url=", videoUrl, err)
	}
	defer videoResp.Body.Close()

	videoBody, err := io.ReadAll(videoResp.Body)
	if err != nil {
		log.Println("error reading videoResp.Body", err)
	}

	var videoResult VideoResult
	err = json.Unmarshal(videoBody, &videoResult)
	if err != nil {
		log.Println("error unmarshalling videoBody", err)
	}

	for _, video := range videoResult.Results {
		if video.Type == "Trailer" {
			m.TrailerURL = fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.Key)
			break
		}
	}
}
