# Real-Time Weather Monitoring System

### Overview

This project focuses on monitoring weather conditions in real-time, aggregating data, and providing daily weather summaries based on rollups and aggregates from the **OpenWeatherMap API**.

### Features

1. Real-time weather data retrieval every 5 minutes.
2. Daily weather summaries including average, maximum, and minimum temperatures.
3. Threshold alerting for specific weather conditions with email notifications.
4. WebSocket support for real-time alerting (planned feature).
5. Visual representation of weather data, including historical trends.

## Steps to Run Locally

1. Clone the repository:

```
git clone https://github.com/yash7xm/Weather_Monitoring_System.git
cd Weather_Monitoring_System
```

2. Set up PostgreSQL:
   `CREATE DATABASE weather_db;`

3. Create environment file:

```
    API_KEY             string
	DATABASE_URL        string
	PORT                string
	MAX_TEMPERATURE     float64
	CONSECUTIVE_UPDATES int
	SMTP_PASS           string
```

4. Install Go dependencies:
   `go mod tidy`

5. Run the server:
   `cd cmd && go run main.go`

### Deployment Notes

1. The backend and PostgreSQL database are deployed on Render, with the frontend hosted on Vercel.
2. Due to free-tier hosting, the backend service may take 1-2 minutes to start upon initial load.
