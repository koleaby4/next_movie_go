package main

import (
	"fmt"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/database"
	"log"
)

func main() {

	dsn, err := config.GetDsn()
	if err != nil {
		log.Fatalln("error fetching db dsn:", err)
	}

	cfg := config.Config{
		DbDsn: dsn,
	}

	_, err = database.New(cfg.DbDsn)
	if err != nil {
		log.Fatalln("error connecting to db:", err)
	}

	fmt.Println("working...")
}
