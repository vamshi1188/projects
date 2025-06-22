package controllers

import (
	"testing"

	"jwtauth/database"
	"jwtauth/models"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	var err error
	database.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}
	err = database.DB.AutoMigrate(&models.User{})
	assert.NoError(t, err)
}

func TestCreateAndReadUser(t *testing.T) {
	setupTestDB(t)

	password, _ := bcrypt.GenerateFromPassword([]byte("test123"), 14)

	user := models.User{
		Name:     "Integration User",
		Email:    "integration@example.com",
		Password: password,
	}

	result := database.DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.Id)

	var fetched models.User
	err := database.DB.First(&fetched, user.Id).Error
	assert.NoError(t, err)

	assert.Equal(t, user.Name, fetched.Name)
	assert.Equal(t, user.Email, fetched.Email)
}

func TestUserNotFound(t *testing.T) {
	setupTestDB(t)

	var user models.User
	err := database.DB.First(&user, 999).Error

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
