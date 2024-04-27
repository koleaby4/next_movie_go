package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	DbDsn string
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
