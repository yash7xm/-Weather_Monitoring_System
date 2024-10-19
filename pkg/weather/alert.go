package weather

import (
	"fmt"
	"log"

	"github.com/yash7xm/Weather_Monitoring_System/config"
	"gopkg.in/gomail.v2"
)

var breachCounter map[string]int = make(map[string]int)

func CheckThresholds(city string, tempC float64) {
	threshold := config.Config.MAX_TEMPERATURE
	consecLimit := config.Config.CONSECUTIVE_UPDATES

	// Check if the temperature exceeds the threshold
	if tempC > threshold {
		breachCounter[city]++
		log.Printf("Threshold exceeded in %s! Current temperature: %.2f°C", city, tempC)

		// Check if the number of consecutive updates exceeds the limit
		if breachCounter[city] >= consecLimit {
			TriggerAlert(city, tempC)
			breachCounter[city] = 0
		}
	} else {
		breachCounter[city] = 0
	}
}

func TriggerAlert(city string, tempC float64) {
	log.Printf("ALERT! Temperature in %s has exceeded %.2f°C for consecutive updates! Current temperature: %.2f°C", city, config.Config.MAX_TEMPERATURE, tempC)
	sendEmailAlert(city, tempC)
}

func sendEmailAlert(city string, tempC float64) {
	m := gomail.NewMessage()
	m.SetHeader("From", "yash.7xm@gmail.com")
	m.SetHeader("To", "yash.7xm@gmail.com")
	m.SetHeader("Subject", fmt.Sprintf("Weather Alert for %s", city))
	m.SetBody("text/plain", fmt.Sprintf("Temperature in %s has exceeded %.2f°C. Current temperature: %.2f°C", city, config.Config.MAX_TEMPERATURE, tempC))

	d := gomail.NewDialer("smtp.gmail.com", 587, "yash.7xm@gmail.com", config.Config.SMTP_PASS)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send alert email: %v", err)
	}
}
