package api

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define your routes
	r.HandleFunc("/api/weather/current", FetchCurrentWeather).Methods("GET")
	r.HandleFunc("/api/weather/summary", GetDailyWeatherSummary).Methods("GET")

	return r
}
