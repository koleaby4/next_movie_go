package plexapi

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetUserToken(appToken string) string {
	url := "http://hp-server2:32400/security/token?type=delegation&scope=all"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Add("X-Plex-Token", appToken)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`token="([^"]*)"`)
	token := r.FindStringSubmatch(string(body))

	if len(token) < 1 {
		log.Fatalln("error fetching token from response", string(body))
	}
	return token[1]
}

type MovieInfo struct {
	Title      string  `xml:"title,attr"`
	Summary    string  `xml:"summary,attr"`
	PosterURL  string  `xml:"thumb,attr"`
	TrailerURL string  `xml:"extras,attr"`
	Rating     float64 `xml:"rating,attr"`
}

type MediaContainer struct {
	Movies []MovieInfo `xml:"Video"`
}

func GetNewestMovies(userTolen string, minRating float64) ([]MovieInfo, error) {
	plexNewestUrl := "http://hp-server2:32400/library/sections/1/newest?X-Plex-Token=" + userTolen
	resp, err := http.Get(plexNewestUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var fetchedMovies MediaContainer
	err = xml.Unmarshal(body, &fetchedMovies)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var movies []MovieInfo

	for _, movie := range fetchedMovies.Movies {
		if movie.Rating >= minRating {
			movies = append(movies, movie)
		}
	}
	return movies, nil
}
