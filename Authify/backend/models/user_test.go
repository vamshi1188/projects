package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jwtauth/models"
)

func TestUserModel(t *testing.T) {
	user := models.User{
		Name:  "Test User",
		Email: "user@example.com",
	}

	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, "user@example.com", user.Email)
}
