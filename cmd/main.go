package main

import (
	"log"
	"net/http"

	"github.com/yash7xm/Weather_Monitoring_System/pkg/api"
)

func main() {
	// Load config (API keys, DB connection)
	// cfg := config.Load()

	// Connect to database
	// db, err := storage.InitDB(cfg.DBConfig)
	// if err != nil {
	//     log.Fatal("Failed to connect to the database:", err)
	// }

	// Start periodic weather fetching (every 5 minutes)
	// go weather.StartWeatherMonitoring(db)

	// Set up API routes
	router := api.SetupRoutes()

	// Start server
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
