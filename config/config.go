package config

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfig struct {
	DbDsn      string     `json:"DSN"`
	TmdbConfig TmdbConfig `json:"TMDB_CONFIG"`
}

type TmdbConfig struct {
	TmdbApiKey  string `json:"API_KEY"`
	TmdbBaseUrl string `json:"BASE_URL"`
}

func ReadFromFile() (AppConfig, error) {
	appConfig := AppConfig{}
	contents, err := os.ReadFile("config/.env.json")
	if err != nil {
		log.Println("error reading config file")
		return AppConfig{}, err
	}

	err = json.Unmarshal(contents, &appConfig)
	if err != nil {
		log.Println("error unmarshalling json", err)
		return AppConfig{}, err
	}

	return appConfig, nil
}
