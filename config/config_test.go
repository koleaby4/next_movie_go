package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpass")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("TMDB_API_KEY", "testapikey")
	os.Setenv("TMDB_BASE_URL", "testbaseurl")
	os.Setenv("TMDB_BACKLOAD_HIGH_WATERMARK_DATE", "2023-01-28")

	// Call GetConfig
	cfg, err := GetConfig()

	// Assert no error
	assert.NoError(t, err)

	// Assert correct values
	assert.Equal(t, "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable", cfg.DbDsn)
	assert.Equal(t, "testapikey", cfg.TmdbConfig.APIKey)
	assert.Equal(t, "testbaseurl", cfg.TmdbConfig.BaseURL)
	assert.Equal(t, "2023-01-28", cfg.TmdbConfig.BackloadHighWatermarkDate)

	// Unset environment variables
	os.Unsetenv("DB_USER")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("TMDB_API_KEY")
	os.Unsetenv("TMDB_BASE_URL")
	os.Unsetenv("TMDB_BACKLOAD_HIGH_WATERMARK_DATE")
}
