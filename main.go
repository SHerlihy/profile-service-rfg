package main

import (
	"flag"
	"os"

	"github.com/SHerlihy/profile-service-rfg/database"
	"github.com/SHerlihy/profile-service-rfg/env_vars"
	"github.com/SHerlihy/profile-service-rfg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	env := flag.String("e", "dev", "environment: dev|prod")
	flag.Parse()

	if *env == "dev" {
		env_vars.DevEnv()
	}

	if *env == "prod" {
		env_vars.ProdEnv()
	}

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
