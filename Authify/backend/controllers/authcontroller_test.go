package controllers_test

import (
	"bytes"
	"encoding/json"
	"jwtauth/controllers"
	"jwtauth/database"
	"jwtauth/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	var err error
	database.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		panic("failed to migrate test database")
	}
	code := m.Run()
	os.Exit(code)
}

func setup() *fiber.App {
	database.DB.Exec("DELETE FROM users")
	app := fiber.New()
	return app
}

func TestRegister(t *testing.T) {
	app := setup()
	app.Post("/api/register", controllers.Register)

	body := map[string]string{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var user models.User
	json.NewDecoder(resp.Body).Decode(&user)
	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
	assert.NotZero(t, user.Id)
}

func TestRegisterInvalidJSON(t *testing.T) {
	app := setup()
	app.Post("/api/register", controllers.Register)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer([]byte(`invalid json`)))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestLogin(t *testing.T) {
	app := setup()
	app.Post("/api/login", controllers.Login)

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	user := models.User{Name: "Test User", Email: "test@example.com", Password: password}
	database.DB.Create(&user)

	body := map[string]string{"email": "test@example.com", "password": "password123"}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "success", response["message"])
}

func TestLoginIncorrectPassword(t *testing.T) {
	app := setup()
	app.Post("/api/login", controllers.Login)

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	user := models.User{Name: "Test User", Email: "wrong@example.com", Password: password}
	database.DB.Create(&user)

	body := map[string]string{"email": "wrong@example.com", "password": "wrongpass"}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestLoginUserNotFound(t *testing.T) {
	app := setup()
	app.Post("/api/login", controllers.Login)

	body := map[string]string{"email": "notfound@example.com", "password": "any"}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestUser(t *testing.T) {
	app := setup()
	app.Get("/api/user", controllers.User)

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	user := models.User{Name: "Test User", Email: "test@example.com", Password: password}
	database.DB.Create(&user)

	claims := jwt.MapClaims{
		"issuer":    user.Id,
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte("secretkey"))

	req := httptest.NewRequest(http.MethodGet, "/api/user", nil)
	req.Header.Set("Cookie", "jwt="+tokenStr)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var got models.User
	json.NewDecoder(resp.Body).Decode(&got)
	assert.Equal(t, user.Id, got.Id)
	assert.Equal(t, user.Email, got.Email)
}

func TestUserInvalidJWT(t *testing.T) {
	app := setup()
	app.Get("/api/user", controllers.User)

	req := httptest.NewRequest(http.MethodGet, "/api/user", nil)
	req.Header.Set("Cookie", "jwt=invalidtoken")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

func TestLogout(t *testing.T) {
	app := setup()
	app.Post("/api/logout", controllers.Logout)

	req := httptest.NewRequest(http.MethodPost, "/api/logout", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "success", response["message"])
}
