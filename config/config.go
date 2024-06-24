package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config holds the configuration for the application
type Config struct {
	DbDsn      string
	TmdbConfig TmdbConfig
	SessionKey string
}

// TmdbConfig holds the configuration for the TMDB API
type TmdbConfig struct {
	APIKey                    string
	BaseURL                   string
	BackloadHighWatermarkDate string
}

func getEnvar(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return val, nil
}

// GetConfig reads the configuration from the environment variables
func GetConfig() (Config, error) {
	err := godotenv.Load(`.env`)
	if err != nil {
		log.Println("Error loading .env file", err)
		return Config{}, err
	}

	cfg := Config{}

	dbUser, err := getEnvar("DB_USER")
	if err != nil {
		log.Println("Error getting DB_USER dbUser", err)
		return cfg, err
	}

	dbPass, err := getEnvar("POSTGRES_PASSWORD")
	if err != nil {
		log.Println("Error getting POSTGRES_PASSWORD dbPass", err)
		return cfg, err
	}

	dbHost, err := getEnvar("DB_HOST")
	if err != nil {
		log.Println("Error getting DB_HOST dbHost", err)
		return cfg, err
	}

	dbName, err := getEnvar("DB_NAME")
	if err != nil {
		log.Println("Error getting DB_NAME dbName", err)
		return cfg, err
	}

	dbDsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbUser, dbPass, dbHost, dbName)
	cfg.DbDsn = dbDsn

	apiKey, err := getEnvar("TMDB_API_KEY")
	if err != nil {
		return cfg, err
	}

	baseURL, err := getEnvar("TMDB_BASE_URL")
	if err != nil {
		return cfg, err
	}

	backloadHighWatermarkDate, err := getEnvar("TMDB_BACKLOAD_HIGH_WATERMARK_DATE")
	if err != nil {
		return cfg, err
	}

	cfg.TmdbConfig = TmdbConfig{
		APIKey:                    apiKey,
		BaseURL:                   baseURL,
		BackloadHighWatermarkDate: backloadHighWatermarkDate,
	}

	return cfg, nil
}
