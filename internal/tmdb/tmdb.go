package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
	RawData     string  `json:"-"`
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

func GetMovies(cfg Config, prefixUrl string, top int) ([]Movie, error) {
	page := 1
	var movies []Movie
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

func GetMoviesReleasedBetween(cfg Config, from time.Time, to time.Time, minRating float64) ([]Movie, error) {
	urlPattern := "%s/discover/movie?primary_release_date.gte=%v&primary_release_date.lte=%v&vote_average.gte=%v&sort_by=release_date.desc"
	url := fmt.Sprintf(urlPattern, cfg.BaseUrl, from.Format("2006-01-02"), to.Format("2006-01-02"), minRating)

	movies, err := GetMovies(cfg, url, 0)
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
	if movie.PosterURL != "" {
		movie.PosterURL = "https://image.tmdb.org/t/p/original" + "/" + strings.Trim(movie.PosterURL, "/")
	}
}

func addTrailerUrl(cfg Config, movie *Movie, wg *sync.WaitGroup) {
	defer wg.Done()

	videosUrl := fmt.Sprintf("%s/movie/%d/videos?api_key=%s", cfg.BaseUrl, movie.Id, cfg.ApiKey)
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
			movie.TrailerURL = fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.Key)
			break
		}
	}
}

func GetMostPopularMovies(cfg Config, minRating float64) ([]Movie, error) {
	url := fmt.Sprintf("%s/discover/movie?vote_average.gte=%v", cfg.BaseUrl, minRating)
	return GetMovies(cfg, url, 20)
}
