package next_movie_go

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// AppConfig is the configuration for the app
type AppConfig struct {
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

func GetEnvar(key string) string {
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

// GetAppConfig reads the config from a file
func GetAppConfig() (AppConfig, error) {
	appConfig := AppConfig{}

	dbDsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", GetEnvar("DB_USER"), GetEnvar("POSTGRES_PASSWORD"), GetEnvar("DB_HOST"), GetEnvar("DB_NAME"))
	log.Println("dbDsn", dbDsn)
	appConfig.DbDsn = dbDsn

	appConfig.TmdbConfig = TmdbConfig{APIKey: GetEnvar("TMDB_API_KEY"), BaseURL: GetEnvar("TMDB_BASE_URL"), BackloadHighWatermarkDate: GetEnvar("TMDB_BACKLOAD_HIGH_WATERMARK_DATE")}
	return appConfig, nil
}
