package user

import (
	"github.com/koleaby4/next_movie_go/internal/database"
	"gorm.io/gorm"
	"log"
	"strings"
)

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex"`
}

func Create(dsn string, user *User) error {
	db := database.New(dsn)

	result := db.Create(user)
	if err := result.Error; err != nil {
		prefix := "ERROR: duplicate key value violates unique constraint"
		if strings.HasPrefix(err.Error(), prefix) {
			log.Printf("user with email=%v already exists\n", user.Email)
			Get(db, user)
			return nil
		}
		return err
	}
	return nil
}

func Update(db *gorm.DB, user *User) (*User, error) {
	result := db.Model(user).Updates(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func Get(db *gorm.DB, user *User) {
	db.Where("email = ?", user.Email).First(user)
}
