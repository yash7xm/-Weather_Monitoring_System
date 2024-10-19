package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	API_KEY      string
	DATABASE_URL string
	PORT         string
}

var Config ENV

func Init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Assign loaded environment variables to the global Config variable
	Config = ENV{
		API_KEY:      os.Getenv("API_KEY"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		PORT:         os.Getenv("PORT"),
	}
}
