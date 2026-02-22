package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	SSLMode string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file was found in project")
	}
	return &Config{
		Port:    getEnv("PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASS", "fdsfdsfds"),
		DBName:  getEnv("DB_NAME", "todo"),
		SSLMode: getEnv("SSL_MODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
