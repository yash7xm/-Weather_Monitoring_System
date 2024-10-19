package api

import (
	"encoding/json"
	"net/http"
	// "weather-monitoring/pkg/storage"
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


// GetWeatherSummary handles the request to get the daily weather summary.
func GetWeatherSummary(w http.ResponseWriter, r *http.Request) {
	// Set content-type to JSON
	w.Header().Set("Content-Type", "application/json")

	// You can get the city from query parameters or a dynamic segment
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	// Fetch the daily weather summary from the database
	// summary, err := storage.GetDailyWeatherSummary(city)
	// if err != nil {
	// 	http.Error(w, "Error fetching weather summary: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Return the weather summary as JSON
	// if err := json.NewEncoder(w).Encode(summary); err != nil {
	// 	http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

// CreateWeatherData handles the request to manually insert weather data (useful for testing).
func CreateWeatherData(w http.ResponseWriter, r *http.Request) {
	// Set content-type to JSON
	w.Header().Set("Content-Type", "application/json")

	var weatherData weather.WeatherResponse
	if err := json.NewDecoder(r.Body).Decode(&weatherData); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Process and save the weather data (This function should save the data to the database)
	// if err := storage.SaveWeatherData(weatherData); err != nil {
	// 	http.Error(w, "Error saving weather data: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Return success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Weather data created successfully"}`))
}

// GetWeatherHistory handles requests to fetch weather history over a specific time range.
// func GetWeatherHistory(w http.ResponseWriter, r *http.Request) {
// 	// Set content-type to JSON
// 	w.Header().Set("Content-Type", "application/json")

// 	// Get parameters from query string
// 	city := r.URL.Query().Get("city")
// 	startDate := r.URL.Query().Get("start")
// 	endDate := r.URL.Query().Get("end")

// 	if city == "" || startDate == "" || endDate == "" {
// 		http.Error(w, "City, start, and end parameters are required", http.StatusBadRequest)
// 		return
// 	}

// 	// Parse the date parameters
// 	startTime, err := time.Parse("2006-01-02", startDate)
// 	if err != nil {
// 		http.Error(w, "Invalid start date format", http.StatusBadRequest)
// 		return
// 	}

// 	endTime, err := time.Parse("2006-01-02", endDate)
// 	if err != nil {
// 		http.Error(w, "Invalid end date format", http.StatusBadRequest)
// 		return
// 	}

// 	// Fetch historical weather data from the database
// 	history, err := storage.GetWeatherHistory(city, startTime, endTime)
// 	if err != nil {
// 		http.Error(w, "Error fetching weather history: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the historical data as JSON
// 	if err := json.NewEncoder(w).Encode(history); err != nil {
// 		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
