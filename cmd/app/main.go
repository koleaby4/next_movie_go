package main

import (
	"fmt"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/database/user"
	"log"
)

func main() {
	fmt.Println("starting...")

	dsn, err := config.GetDsn()

	if err != nil {
		log.Fatalln("error fetching dsn:", err)
	}

	alex := &user.User{Email: "alex@gmail.com"}

	err = user.Create(dsn, alex)
	if err != nil {
		log.Fatalln("error creating user alex:", err)
	}

	log.Println("user alex:", *alex)

}
