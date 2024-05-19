package tmdb

import (
	"encoding/json"
	"fmt"
	"github.com/koleaby4/next_movie_go/internal/models"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type MovieSearchResults struct {
	Results    []models.Movie `json:"results"`
	Page       int            `json:"page"`
	TotalPages int            `json:"total_pages"`
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

func GetMovie(config Config, movieID string) (models.Movie, error) {
	url := fmt.Sprintf("%s/movie/%s", config.BaseUrl, movieID)
	movies, err := GetMovies(config, url, 1)
	if err != nil {
		return models.Movie{}, err
	}

	if len(movies) == 0 {
		return models.Movie{}, fmt.Errorf("movie with id %s not found", movieID)
	}

	return movies[0], nil
}

func GetMovies(cfg Config, prefixUrl string, top int) ([]models.Movie, error) {
	page := 1
	var movies []models.Movie
	var url string

	for {
		tail := fmt.Sprintf("page=%d&api_key=%s", page, cfg.ApiKey)

		if strings.Contains(prefixUrl, "?") {
			url = prefixUrl + "&" + tail
		} else {
			url = prefixUrl + "?" + tail
		}

		log.Println("url", url)
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

		if pageResults.TotalPages == 0 {
			var movieResult models.Movie
			err = json.Unmarshal(body, &movieResult)
			return []models.Movie{movieResult}, err
		}

		for i := range pageResults.Results {
			rawData, err := json.Marshal(pageResults.Results[i])
			if err != nil {
				return movies, err
			}
			pageResults.Results[i].RawData = string(rawData)
		}

		movies = append(movies, pageResults.Results...)
		if top > 0 && len(movies) >= top {
			break
		}
		if pageResults.Page >= pageResults.TotalPages {
			break
		}
		page++
	}
	fmt.Println("number of movies fetched:", len(movies))
	return EnrichMoviesInfo(cfg, movies), nil
}

func GetMoviesReleasedBetween(cfg Config, from time.Time, to time.Time, minRating float64) ([]models.Movie, error) {
	urlPattern := "%s/discover/movie?primary_release_date.gte=%v&primary_release_date.lte=%v&vote_average.gte=%v&sort_by=release_date.desc"
	url := fmt.Sprintf(urlPattern, cfg.BaseUrl, from.Format("2006-01-02"), to.Format("2006-01-02"), minRating)

	movies, err := GetMovies(cfg, url, 0)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func EnrichMoviesInfo(cfg Config, movies []models.Movie) []models.Movie {
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

func expandPosterUrl(movie *models.Movie) {
	if movie.PosterUrl != "" {
		movie.PosterUrl = "https://image.tmdb.org/t/p/original" + "/" + strings.Trim(movie.PosterUrl, "/")
	}
}

func addTrailerUrl(cfg Config, movie *models.Movie, wg *sync.WaitGroup) {
	defer wg.Done()

	videosUrl := fmt.Sprintf("%s/movie/%d/videos?api_key=%s", cfg.BaseUrl, movie.ID, cfg.ApiKey)
	videoResp, err := http.Get(videosUrl)
	if err != nil {
		log.Println("error fetching video url=", videosUrl, err)
		return
	}
	defer videoResp.Body.Close()

	videoBody, err := io.ReadAll(videoResp.Body)
	if err != nil {
		log.Printf("error reading videoResp.Body for videosUrl=%v, err=%v\n", videosUrl, err)
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
			movie.TrailerUrl = fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.Key)
			break
		}
	}
}

func GetMostPopularMovies(cfg Config, minRating float64) ([]models.Movie, error) {
	url := fmt.Sprintf("%s/discover/movie?vote_average.gte=%v", cfg.BaseUrl, minRating)
	return GetMovies(cfg, url, 20)
}
