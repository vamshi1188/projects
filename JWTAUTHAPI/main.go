package main

import (
	"jwtauth/database"
	"jwtauth/routes"

	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	// Initialize a new Fiber app
	app := fiber.New()

	routes.Setup(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":8000"))
}
