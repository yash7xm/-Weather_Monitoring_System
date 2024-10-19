package api

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define your routes
	r.HandleFunc("/api/weather/current", FetchCurrentWeather).Methods("GET")
	r.HandleFunc("/api/weather", GetWeatherSummary).Methods("GET")
	// r.HandleFunc("/api/alerts", CreateAlert).Methods("POST")

	return r
}
