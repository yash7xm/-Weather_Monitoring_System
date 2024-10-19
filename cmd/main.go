package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yash7xm/Weather_Monitoring_System/pkg/api"
	db "github.com/yash7xm/Weather_Monitoring_System/pkg/storage"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database connection
	db.InitDB()
	defer db.DB.Close() // Close the DB connection when the application shuts down

	// Run database migrations
	err = db.RunMigrations()
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Start periodic weather fetching (every 5 minutes)
	// go weather.StartWeatherMonitoring(db)

	// Set up API routes
	router := api.SetupRoutes()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if PORT is not set (for local development)
	}

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
