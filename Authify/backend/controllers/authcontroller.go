package controllers

import (
	"jwtauth/database"
	"jwtauth/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "secretkey"

func Register(c *fiber.Ctx) error {
	var input map[string]string
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	name := input["name"]
	email := input["email"]
	password := input["password"]

	if name == "" || email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// Check if user already exists
	var existing models.User
	if err := database.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	database.DB.Create(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var input map[string]string
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	email := input["email"]
	password := input["password"]

	var user models.User
	database.DB.Where("email = ?", email).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	claims := jwt.MapClaims{
		"issuer":    user.Id,
		"expiresAt": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization", // changed from jwt
		Value:    tokenStr,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("Authorization") // changed from jwt

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.MapClaims)
	userId := int((*claims)["issuer"].(float64))

	var user models.User
	database.DB.First(&user, userId)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "Logged out",
	})
}
