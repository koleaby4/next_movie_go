package main

import (
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/database"
	"github.com/koleaby4/next_movie_go/internal/database/user"
	"gorm.io/gorm"
	"log"
)

func migrate(db *gorm.DB) {

	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatalln("error migrating users:", err)
	}

}

func main() {
	log.Println("running migrations...")

	dsn, err := config.GetDsn()
	if err != nil {
		log.Fatalln("error fetching dsn:", err)
	}

	db := database.New(dsn)

	migrate(db)

	log.Println("db migration completed.")
}
