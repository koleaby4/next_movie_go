package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	email string `gorm:"uniqueIndex"`
}
