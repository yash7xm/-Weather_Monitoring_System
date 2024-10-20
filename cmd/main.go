package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
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

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://weather-ui-r7xp.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS handler
	handler := corsHandler.Handler(router)

	// Start server
	port := config.Config.PORT
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
