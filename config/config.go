package config

import (
	"errors"
	"fmt"
	"github.com/koleaby4/next_movie_go/internal/database"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

type Config struct {
	Db *gorm.DB
}

func SetDb(config Config, dsn string) {
	db, err := database.New(dsn)
	if err != nil {
		log.Fatalln("error setting up db:", err)
	}

	config.Db = db
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
