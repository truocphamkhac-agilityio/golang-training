package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB 		*DBConfig
	Port 	string
}

type DBConfig struct {
	DbDialect 	string
	DbUsername 	string
	DbPassword 	string
	DbPort 		string
	DbHost 		string
	DbName 		string
}

func GetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Println("We are getting the env values")
	}

	return &Config{
		DB: &DBConfig{
			DbDialect: 	os.Getenv("DB_DRIVER"),
			DbUsername: os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbPort: 	os.Getenv("DB_PORT"),
			DbHost: 	os.Getenv("DB_HOST"),
			DbName: 	os.Getenv("DB_NAME"),
		},
		Port: os.Getenv("PORT"),
	}
}