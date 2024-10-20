package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yash7xm/Weather_Monitoring_System/config"
	db "github.com/yash7xm/Weather_Monitoring_System/pkg/storage"
)

type WeatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	}
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Timestamp int64 `json:"dt"`
}

// FetchWeather fetches weather data from the OpenWeatherMap API and stores it in the database
func FetchWeather(city string) (*WeatherResponse, error) {
	apiKey := config.Config.API_KEY
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert temperature from Kelvin to Celsius
	tempC := weatherResp.Main.Temp - 273.15
	feelsLikeC := weatherResp.Main.FeelsLike - 273.15
	mainCondition := weatherResp.Weather[0].Main

	// Insert data into the database
	timestamp := time.Unix(weatherResp.Timestamp, 0)
	_, err = db.DB.Exec(`INSERT INTO weather_data (city, temperature, feels_like, main_condition, timestamp) VALUES ($1, $2, $3, $4, $5)`,
		city, tempC, feelsLikeC, mainCondition, timestamp)

	CheckThresholds(city, tempC)

	return &weatherResp, err
}

func StartWeatherMonitoring() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		for _, city := range []string{"Delhi", "Mumbai", "Chennai", "Bangalore", "Kolkata", "Hyderabad"} {
			go func(city string) {
				weather, err := FetchWeather(city)
				if err != nil {
					log.Printf("Error fetching weather data for %s: %v", city, err)
					return
				}
				log.Printf("Fetched weather data for %s: %+v", city, weather)
			}(city)
		}
	}
}
