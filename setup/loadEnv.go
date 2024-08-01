package setup

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"SUPA_DATABASE_HOST"`
	DBPort     string `mapstructure:"SUPA_DATABASE_PORT"`
	DBName     string `mapstructure:"SUPA_DATABASE_NAME"`
	DBUserName string `mapstructure:"SUPA_DATABASE_USER"`
	DBPassword string `mapstructure:"SUPA_DATABASE_PASSWORD"`

	ServerAddr string `mapstructure:"SERVER_ADDR"`

	APIRef string `mapstructure:"SUPA_API_REF"`
	APIKey string `mapstructure:"SUPA_API_ANON_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	// Load .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
	}
	config.DBHost = viper.GetString("SUPA_DATABASE_HOST")
	config.DBPort = viper.GetString("SUPA_DATABASE_PORT")
	config.DBName = viper.GetString("SUPA_DATABASE_NAME")
	config.DBUserName = viper.GetString("SUPA_DATABASE_USER")
	config.DBPassword = viper.GetString("SUPA_DATABASE_PASSWORD")
	config.ServerAddr = viper.GetString("SERVER_ADDR")
	config.APIRef = viper.GetString("SUPA_API_REF")
	config.APIKey = viper.GetString("SUPA_API_ANON_KEY")
	return
}
