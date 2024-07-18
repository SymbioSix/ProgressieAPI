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
	router.Get("/liveness-check",
		healthcheck.NewHealthChecker(
			healthcheck.Config{
				Probe: func(c fiber.Ctx) bool { return true },
			},
		),
	)
	router.Get("/healthcheck", func(c fiber.Ctx) error {
		var database_status string = "ready"
		var supabase_api_status string = "ready"
		var overall_status string = "super healthy"
		healthmap := fiber.Map{
			"database_status":     database_status,
			"supabase_api_status": supabase_api_status,
			"overall_status":      overall_status,
		}
		if s.DB.Error != nil && !s.Client.Rest.Ping() {
			database_status = "error"
			supabase_api_status = "error"
			overall_status = "having issue(s) : database and supabase"
			return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
		}
		if s.DB.Error != nil {
			database_status = "error"
			overall_status = "having issue(s) : database"
			return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
		}
		if !s.Client.Rest.Ping() {
			supabase_api_status = "error"
			overall_status = "having issue(s) : supabase"
			return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
		}
		return c.Status(fiber.StatusOK).JSON(healthmap)
	})

	// Connect all the routes
	AuthRouter.AuthRoutes(router)

	// Serve The API
	s.StartServerWithGracefulShutdown(app, &config)
}
