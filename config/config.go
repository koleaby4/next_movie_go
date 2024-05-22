package config

import (
	"errors"
	"fmt"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	DbDsn string
	Tmdb  tmdb.Config
}

func GetFromFile() (AppConfig, error) {
	appConfig := AppConfig{}
	contents, err := os.ReadFile("config/.env")
	if err != nil {
		log.Println("error reading config file")
		return AppConfig{}, err
	}

	for _, line := range strings.Split(string(contents), "\n") {
		parts := strings.Split(line, "=")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "DSN":
			appConfig.DbDsn = value
		case "TMDB_API_KEY":
			appConfig.Tmdb.ApiKey = value
		default:
			log.Printf("unknown key: %v\n", key)
		}
	}

	return appConfig, nil
}

func GetDsn() (string, error) {

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Println("DSN env var not found - fetching from config file")

		example := "host=localhost user=dbUsername password=dbPassword dbname=dbName port=5432 sslmode=disable"
		msg := fmt.Sprintf("please declare env var 'DSN' with the following format: %v", example)
		log.Println(msg)
		return "", errors.New(msg)
	}

	dsn = strings.Trim(dsn, `"`)
	dsn = strings.Trim(dsn, `'`)
	return dsn, nil
}

func GetTmdbApiKey() string {
	token := os.Getenv("TMDB_API_KEY")
	if token == "" {
		log.Fatalln("env var TMDB_API_KEY not found")
	}
	return token
}
