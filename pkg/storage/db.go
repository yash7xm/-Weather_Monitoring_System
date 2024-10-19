package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/yash7xm/Weather_Monitoring_System/config"
)

var DB *sql.DB

// Initialize the database connection
func InitDB() {
	connStr := config.Config.DATABASE_URL
	if connStr == "" {
		log.Fatalf("DATABASE_URL is not set")
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Test the database connection
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Connected to the database successfully!")
}

// Run database migrations
func RunMigrations() error {
	migrationQuery := `
	CREATE TABLE IF NOT EXISTS weather_data (
		id SERIAL PRIMARY KEY,
		city VARCHAR(100),
		temperature DECIMAL,
		feels_like DECIMAL,
		main_condition VARCHAR(100),
		timestamp TIMESTAMP
	);
	`

	_, err := DB.Exec(migrationQuery)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
		return err
	}

	fmt.Println("Database migrations completed successfully!")
	return nil
}
