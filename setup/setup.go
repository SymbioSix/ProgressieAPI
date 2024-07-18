package setup

import (
	"log"
	"os"
	"os/signal"

	api "github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	Client *api.Client
)

func ConnectDatabase(config *Config) error {
	var err error
	dbName := config.DBName
	dbHost := config.DBHost
	dbPort := config.DBPort
	dbUser := config.DBUserName
	dbPass := config.DBPassword
	database, err := gorm.Open(postgres.Open("user="+dbUser+" password="+dbPass+" host="+dbHost+" port="+dbPort+" dbname="+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Can't connect to database!")
	}

	log.Println("Connected to Database!")

	DB = database
	return err
}

func ConnectViaAPI(config *Config) {
	apiKey := config.APIKey
	apiRef := config.APIRef

	client, err := api.NewClient(apiRef, apiKey, &api.ClientOptions{})
	if err != nil {
		log.Fatal(err)
		panic("Can't connect to API Gateway")
	}

	log.Println("Connected to API Gateway")

	Client = client
}

func StartServerWithGracefulShutdown(a *fiber.App, config *Config) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(config.ServerAddr); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
