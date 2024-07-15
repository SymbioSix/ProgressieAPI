package setup

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"SUPA_DATABASE_HOST"`
	DBPort     string `mapstructure:"SUPA_DATABASE_PORT"`
	DBName     string `mapstructure:"SUPA_DATABASE_NAME"`
	DBUserName string `mapstructure:"SUPA_DATABASE_USER"`
	DBPassword string `mapstructure:"SUPA_DATABASE_PASSWORD"`

	ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
