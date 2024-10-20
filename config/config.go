package config

import (
	// "log"
	"os"
	"strconv"
	// "github.com/joho/godotenv"
)

type ENV struct {
	API_KEY             string
	DATABASE_URL        string
	PORT                string
	MAX_TEMPERATURE     float64
	CONSECUTIVE_UPDATES int
	SMTP_PASS           string
}

var Config ENV

func Init() {
	// Load environment variables from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	maxTempStr := os.Getenv("MAX_TEMPERATURE")
	consecUpdatesStr := os.Getenv("CONSECUTIVE_UPDATES")

	maxTemp, err := strconv.ParseFloat(maxTempStr, 64)
	if err != nil {
		maxTemp = 35.0 // Default max temperature threshold
	}

	consecUpdates, err := strconv.Atoi(consecUpdatesStr)
	if err != nil {
		consecUpdates = 2 // Default consecutive updates threshold
	}

	// Assign loaded environment variables to the global Config variable
	Config = ENV{
		API_KEY:             os.Getenv("API_KEY"),
		DATABASE_URL:        os.Getenv("DATABASE_URL"),
		PORT:                os.Getenv("PORT"),
		MAX_TEMPERATURE:     maxTemp,
		CONSECUTIVE_UPDATES: consecUpdates,
		SMTP_PASS:           os.Getenv("SMTP_PASS"),
	}
}
