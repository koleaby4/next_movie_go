package config

import (
	"encoding/json"
	"log"
	"os"
)

// AppConfig is the configuration for the app
type AppConfig struct {
	DbDsn      string     `json:"DSN"`
	TmdbConfig TmdbConfig `json:"TMDB_CONFIG"`
	SessionKey string     `json:"SESSION_KEY"`
}

// TmdbConfig is the configuration for the tmdb api
type TmdbConfig struct {
	APIKey                    string `json:"API_KEY"`
	BaseURL                   string `json:"BASE_URL"`
	BackloadHighWatermarkDate string `json:"BACKLOAD_HIGH_WATERMARK_DATE"`
}

// ReadFromFile reads the config from a file
func ReadFromFile(filePath string) (AppConfig, error) {
	appConfig := AppConfig{}
	contents, err := os.ReadFile(filePath)
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
