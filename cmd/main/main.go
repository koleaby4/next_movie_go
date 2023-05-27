package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	loadDotEnv()
}

func loadDotEnv() {

	dotEnvContent, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalf("error reading .env file: %v", err)
	}

	lines := strings.Split(string(dotEnvContent), "\n")
	for _, line := range lines {
		if line == "" || line == "\n" {
			continue
		}
		args := strings.Split(line, "=")
		name := strings.TrimSpace(args[0])
		value := strings.TrimSpace(args[1])
		os.Setenv(name, value)
	}
}

func getMoviesByRating() {

	url := "https://moviesminidatabase.p.rapidapi.com/movie/order/byRating/"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("error building request to fetch moves: %v", err)
	}

	rapid_api_key := os.Getenv("RAPID_API_KEY")

	req.Header.Add("X-RapidAPI-Key", rapid_api_key)
	req.Header.Add("X-RapidAPI-Host", "moviesminidatabase.p.rapidapi.com")

    client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("error fetching movies by rating: %v", err)
	}

	defer res.Body.Close()

	if statusCode := res.StatusCode; statusCode != 200 {
		log.Fatalf("unexpected status code: %v", statusCode)
	}

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

}

func main() {
	getMoviesByRating()

}
