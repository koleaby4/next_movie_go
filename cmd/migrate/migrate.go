package main

import (
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/database"
	"gorm.io/gorm"
	"log"
)

func migrate(db *gorm.DB) {

	if err := db.AutoMigrate(&database.User{}); err != nil {
		log.Fatalln("error migrating users:", err)
	}

}

func main() {
	log.Println("running migrations...")

	dsn, err := config.GetDsn()
	if err != nil {
		log.Fatalln("error fetching dsn:", err)
	}

	db, err := database.New(dsn)
	if err != nil {
		log.Fatalln("error connecting to db:", err)
	}

	migrate(db)

	log.Println("db migration completed.")
}
