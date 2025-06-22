package database_test

import (
	"jwtauth/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	database.Connect()
	assert.NotNil(t, database.DB)
}

func TestConnect(t *testing.T) {
	database.Connect()
	assert.NotNil(t, database.DB)
}
