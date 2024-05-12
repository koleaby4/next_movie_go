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

func GetDsn() (string, error) {

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Println("DSN env var not found")
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
