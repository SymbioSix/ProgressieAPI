package setup

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *Config) {
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
}
