package main

import (
	"fmt"
	"github.com/koleaby4/next_movie_go/config"
	"log"
)

func main() {

	cfg := config.Config{}
	dsn, err := config.GetDsn()

	if err != nil {
		log.Fatalln("error fetching dsn:", err)
	}
	config.SetDb(cfg, dsn)

	fmt.Println("working...")
}
