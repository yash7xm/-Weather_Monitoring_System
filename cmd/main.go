package main

import (
	"log"
	"net/http"

	"github.com/yash7xm/Weather_Monitoring_System/config"
	"github.com/yash7xm/Weather_Monitoring_System/pkg/api"
	db "github.com/yash7xm/Weather_Monitoring_System/pkg/storage"
	"github.com/yash7xm/Weather_Monitoring_System/pkg/weather"
)

func main() {
	// Initialize configuration
	config.Init()

	// Initialize the database connection
	db.InitDB()
	defer db.DB.Close()

	// Run database migrations
	err := db.RunMigrations()
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Start periodic weather fetching (every 5 minutes)
	go weather.StartWeatherMonitoring()

	// Set up API routes
	router := api.SetupRoutes()

	// Start server
	port := config.Config.PORT
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
