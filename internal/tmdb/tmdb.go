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
	Results    []Movie `json:"results"`
	Page       int     `json:"page"`
	TotalPages int     `json:"total_pages"`
}

type VideoResult struct {
	Results []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
	} `json:"results"`
}

type Config struct {
	BaseUrl string
	ApiKey  string
}

func GetMovies(cfg Config, url string) ([]Movie, error) {
	page := 1
	var movies []Movie
	for {
		url = fmt.Sprintf("%s&page=%d&api_key=%s", url, page, cfg.ApiKey)
		resp, err := http.Get(url)
		if err != nil {
			return movies, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return movies, err
		}

		var pageResults MovieSearchResults
		err = json.Unmarshal(body, &pageResults)
		if err != nil {
			return movies, err
		}

		movies = append(movies, pageResults.Results...)
		if pageResults.Page >= pageResults.TotalPages {
			break
		}
		page++
	}
	fmt.Println("number of movies fetched:", len(movies))
	return movies, nil
}

func GetRecentMovies(cfg Config, minRating float64, releasedAfter time.Time) ([]Movie, error) {
	url := fmt.Sprintf("%s/3/discover/movie?primary_release_date.gte=%v&vote_average.gte=%v&sort_by=release_date.desc", cfg.BaseUrl, releasedAfter.Format("2006-01-02"), minRating)
	movies, err := GetMovies(cfg, url)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func EnrichMoviesInfo(cfg Config, movies []Movie) []Movie {
	wg := sync.WaitGroup{}
	wg.Add(len(movies))

	for i := range movies {
		movie := &movies[i]
		go addTrailerUrl(cfg, movie, &wg)
		expandPosterUrl(movie)
	}
	wg.Wait()

	return movies
}

func expandPosterUrl(movie *Movie) {
	const imageBaseUrl = "https://image.tmdb.org/t/p/original"
	if movie.PosterURL != "" {
		movie.PosterURL = fmt.Sprintf("%s%s", imageBaseUrl, movie.PosterURL)
	}

}

func addTrailerUrl(cfg Config, movie *Movie, wg *sync.WaitGroup) {
	defer wg.Done()

	videoUrl := fmt.Sprintf("%s/3/movie/%d/videos?api_key=%s", cfg.BaseUrl, movie.Id, cfg.ApiKey)
	videoResp, err := http.Get(videoUrl)
	if err != nil {
		log.Println("error fetching video url=", videoUrl, err)
		return
	}
	defer videoResp.Body.Close()

	videoBody, err := io.ReadAll(videoResp.Body)
	if err != nil {
		log.Println("error reading videoResp.Body", err)
		return
	}

	var videoResult VideoResult
	err = json.Unmarshal(videoBody, &videoResult)
	if err != nil {
		log.Println("error unmarshalling videoBody", err)
		return
	}

	for _, video := range videoResult.Results {
		if video.Type == "Trailer" {
			movie.TrailerURL = fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.Key)
			break
		}
	}
}

func GetMostPopularMovies(cfg Config, minRating float64, page int) ([]Movie, error) {
	url := fmt.Sprintf("%s/3/discover/movie?vote_average.gte=%v&sort_by=vote_average.desc&page=%v&api_key=%v", cfg.BaseUrl, minRating, page, cfg.ApiKey)
	fmt.Println("url", url)
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

	return results.Results, nil
}
