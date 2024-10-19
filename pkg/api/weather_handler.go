package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yash7xm/Weather_Monitoring_System/pkg/weather"
)

// FetchCurrentWeather handles the request to get current weather data from the OpenWeatherMap API.
func FetchCurrentWeather(w http.ResponseWriter, r *http.Request) {
	// Set content-type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the city parameter from the query
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	// Fetch the current weather data from the OpenWeatherMap API
	weatherData, err := weather.FetchWeather(city)
	if err != nil {
		http.Error(w, "Error fetching weather data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the fetched weather data as JSON
	if err := json.NewEncoder(w).Encode(weatherData); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler function to get the daily weather summary
func GetDailyWeatherSummary(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	date := r.URL.Query().Get("date")

	if date == "" {
		// If no date is provided, use the current date
		date = time.Now().Format("2006-01-02")
	}

	summary, err := weather.RollUpDailySummaries(city, date)
	if err != nil {
		http.Error(w, "Error retrieving daily summary", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
