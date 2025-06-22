package main

import (
	"jwtauth/database"
	"jwtauth/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":8000"))
}
