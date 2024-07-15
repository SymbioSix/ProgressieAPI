package setup

import (
	"log"

	api "github.com/SymbioSix/ProgressieAPI/utils"
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
