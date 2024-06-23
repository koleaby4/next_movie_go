package next_movie_go

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config is the configuration for the app
type Config struct {
	DbDsn      string
	TmdbConfig TmdbConfig
	SessionKey string
}

// TmdbConfig is the configuration for the tmdb api
type TmdbConfig struct {
	APIKey                    string
	BaseURL                   string
	BackloadHighWatermarkDate string
}

func getEnvar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file", err)
	}
	if val := os.Getenv(key); val != "" {
		return val
	}
	log.Fatalf("Environment variable %s not set", key)
	return ""
}

// GetConfig reads the config from a file
func GetConfig() (Config, error) {
	cfg := Config{}

	dbDsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", getEnvar("DB_USER"), getEnvar("POSTGRES_PASSWORD"), getEnvar("DB_HOST"), getEnvar("DB_NAME"))
	cfg.DbDsn = dbDsn

	cfg.TmdbConfig = TmdbConfig{APIKey: getEnvar("TMDB_API_KEY"), BaseURL: getEnvar("TMDB_BASE_URL"), BackloadHighWatermarkDate: getEnvar("TMDB_BACKLOAD_HIGH_WATERMARK_DATE")}
	return cfg, nil
}
