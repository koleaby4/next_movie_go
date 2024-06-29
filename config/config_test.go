package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnvars(envars map[string]string) {
	for k, v := range envars {
		os.Setenv(k, v)
	}
}

func unsetEnvars(envars map[string]string) {
	for k := range envars {
		os.Unsetenv(k)
	}
}

func TestGetConfig(t *testing.T) {
	envars := map[string]string{
		"DB_USER":                           "testuser",
		"POSTGRES_PASSWORD":                 "testpass",
		"DB_HOST":                           "localhost",
		"DB_NAME":                           "testdb",
		"TMDB_API_KEY":                      "testapikey",
		"TMDB_BASE_URL":                     "testbaseurl",
		"TMDB_BACKLOAD_HIGH_WATERMARK_DATE": "2023-01-28",
	}

	setEnvars(envars)

	cfg, err := GetConfig()

	assert.NoError(t, err)

	assert.Equal(t, "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable", cfg.DbDsn)
	assert.Equal(t, "testapikey", cfg.TmdbConfig.APIKey)
	assert.Equal(t, "testbaseurl", cfg.TmdbConfig.BaseURL)
	assert.Equal(t, "2023-01-28", cfg.TmdbConfig.BackloadHighWatermarkDate)

	unsetEnvars(envars)
}
