package main

import (
	"log"

	s "github.com/SymbioSix/ProgressieAPI/setup"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

var (
	app *fiber.App

	// TODO: Create New Service Controller and Router Variables
)

func init() {
	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Initialize Database and API Connectivity
	s.ConnectDatabase(&config)
	s.ConnectViaAPI(&config)

	// TODO: Initialize Routers and Controllers

	app = fiber.New()
}

func main() {
	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	corsConfig := cors.Config{
		// Allow Origins Will Be Updated With Our Web Domain
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}

	app.Use(cors.New(corsConfig))

	router := app.Group("/v1")
	router.Get("/healthcheck", func(c fiber.Ctx) error {
		status := "healthy"
		description := "API is Ready"
		healthmap := fiber.Map{
			status:      status,
			description: description,
		}
		if s.DB.Error != nil || !s.Client.Rest.Ping() {
			status = "unhealthy"
			description = "API has some error"
			return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
		}
		return c.JSON(healthmap)
	})

	// Connect all the routes

	// Serve The API
	app.Listen(config.ServerAddr)
	// app.Server().ListenAndServe(config.ServerAddr)
}
