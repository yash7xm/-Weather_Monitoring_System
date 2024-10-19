package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/yash7xm/Weather_Monitoring_System/config"
)

var apiKey = config.Config.API_KEY

// const cities = "Delhi,Mumbai,Chennai,Bangalore,Kolkata,Hyderabad"

type WeatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Dt int64 `json:"dt"`
}

// FetchWeather fetches current weather data for the given city from OpenWeatherMap API
func FetchWeather(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		// Read the response body for debugging
		var errorResponse map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		return nil, fmt.Errorf("failed to fetch weather: %s", errorResponse["message"])
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &weatherResp, nil
}

// func StartWeatherMonitoring(db *DB) {
// 	ticker := time.NewTicker(5 * time.Minute) // Fetch every 5 minutes
// 	for range ticker.C {
// 		for _, city := range []string{"Delhi", "Mumbai", "Chennai", "Bangalore", "Kolkata", "Hyderabad"} {
// 			weather, err := FetchWeather(city)
// 			if err != nil {
// 				fmt.Println("Error fetching weather data:", err)
// 				continue
// 			}
// 			fmt.Println("Weather in", city, ":", weather)
// 			// Process and store the data in the database
// 		}
// 	}
// }
