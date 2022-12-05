package configs

import (
	"log"

	"github.com/spf13/viper"
)

type envConfig struct {
	Port string `mapstructure:"PORT"`
	Db_User string `mapstructure:"DB_USER"`
	Db_Host string `mapstructure:"DB_HOST"`
	Db_Name string `mapstructure:"DB_NAME"`
	Db_Password string `mapstructure:"DB_PASSWORD"`
	Db_Port string `mapstructure:"DB_PORT"`
}

func Envload() (config *envConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error loading config file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("error unmarshaling config file", err)
	}

	return config
}
