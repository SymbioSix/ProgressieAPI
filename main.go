package main

import (
	"log"

	au_r "github.com/SymbioSix/ProgressieAPI/routers/auth"
	au_s "github.com/SymbioSix/ProgressieAPI/services/auth"
	s "github.com/SymbioSix/ProgressieAPI/setup"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

var (
	app *fiber.App

	// TODO: Create New Service Controller and Router Variables
	AuthController au_s.AuthController
	AuthRouter     au_r.AuthRouter
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

	AuthController = au_s.NewAuthController(s.DB, s.Client)
	AuthRouter = au_r.NewRouteAuthController(AuthController)

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
		database_status := "ready"
		supabase_api_status := "ready"
		overall_status := "super healthy"
		healthmap := fiber.Map{
			"database_status":     database_status,
			"supabase_api_status": supabase_api_status,
			"overall_status":      overall_status,
		}
		if s.DB.Error != nil {
			database_status = "error"
			if !s.Client.Rest.Ping() {
				supabase_api_status = "error"
			}
			return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
		}
		return c.JSON(healthmap)
	}, healthcheck.NewHealthChecker(healthcheck.Config{Probe: func(c fiber.Ctx) bool { return true }}))

	// Connect all the routes
	AuthRouter.AuthRoutes(router)

	// Serve The API
	app.Listen(config.ServerAddr)
	// app.Server().ListenAndServe(config.ServerAddr)
}
