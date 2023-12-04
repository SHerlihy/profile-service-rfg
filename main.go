package main

import (
	"os"

	"github.com/SHerlihy/profile-service-rfg/database"
	"github.com/SHerlihy/profile-service-rfg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ALLOWED_ORIGINS string
var DATABASE_ADDRESS string
var DATABASE_USER string
var DATABASE_PASSWORD string

func main() {
	initEnvVars()

	database.Connect()

	app := fiber.New()

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowCredentials: true,
	}))

	api := app.Group("/api/v1")

	routes.Setup(api)

	app.Listen(":5010")
}

func initEnvVars() {
	os.Setenv("ALLOWED_ORIGINS", ALLOWED_ORIGINS)
	os.Setenv("DATABASE_ADDRESS", DATABASE_ADDRESS)
	os.Setenv("DATABASE_USER", DATABASE_USER)
	os.Setenv("DATABASE_PASSWORD", DATABASE_PASSWORD)
}
