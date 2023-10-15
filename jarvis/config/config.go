package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server ServerConfig
	Gpt2   ServerConfig
}

type ServerConfig struct {
	Host string
	Port string
}

func ReadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Error reading config file, %s", err)
	}

	viper.SetDefault("database.dbname", "test_db")

	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal("Unable to read configs, %v", err)
	}

	return conf
}
