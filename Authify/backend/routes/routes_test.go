package routes_test

import (
	"jwtauth/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()
	routes.Setup(app)

	// You can test just the route existence, not logic
	req := httptest.NewRequest(http.MethodPost, "/api/register", nil)
	resp, err := app.Test(req)

	// We are only verifying that the route exists and doesn't 404
	assert.Nil(t, err)
	assert.NotEqual(t, http.StatusNotFound, resp.StatusCode)
}
