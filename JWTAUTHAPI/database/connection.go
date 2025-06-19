package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"jwtauth/models"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
