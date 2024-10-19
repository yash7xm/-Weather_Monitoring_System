package weather

import (
	"database/sql"

	db "github.com/yash7xm/Weather_Monitoring_System/pkg/storage"
)

// RollUpDailySummaries aggregates the weather data for a specific day
func RollUpDailySummaries(city string, date string) (map[string]interface{}, error) {
	query := `
		SELECT
			AVG(temperature) AS avg_temp,
			MAX(temperature) AS max_temp,
			MIN(temperature) AS min_temp,
			(SELECT main_condition FROM weather_data WHERE city = $1 AND DATE(timestamp) = $2 GROUP BY main_condition ORDER BY COUNT(*) DESC LIMIT 1) AS dominant_condition
		FROM
			weather_data
		WHERE
			city = $1 AND DATE(timestamp) = $2
	`

	var avgTemp, maxTemp, minTemp sql.NullFloat64
	var dominantCondition sql.NullString

	err := db.DB.QueryRow(query, city, date).Scan(&avgTemp, &maxTemp, &minTemp, &dominantCondition)
	if err != nil {
		return nil, err
	}

	summary := map[string]interface{}{
		"average_temperature": avgTemp,
		"maximum_temperature": maxTemp,
		"minimum_temperature": minTemp,
		"dominant_condition":  dominantCondition,
	}

	return summary, nil
}
