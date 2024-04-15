package configuration

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	DBName       string
	SSLMode      string
	Host         string
	DatabasePort string
	TimeZone     string
}

func GetDatabaseConfig() string {
	dbConfig := DatabaseConfig{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		Host:         os.Getenv("DB_HOST"),
		DatabasePort: os.Getenv("DB_PORT"),
		SSLMode:      os.Getenv("SSL_MODE"),
		TimeZone:     os.Getenv("TIME_ZONE"),
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBName,
	)
}