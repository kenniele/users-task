package config

import (
	"os"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func NewConfig() *Config {
	return &Config{
		DBHost: getEnv("DBHost", "localhost"),
		DBPort: getEnv("DBPort", "5432"),
		DBUser: getEnv("DBUser", "postgres"),
		DBPass: getEnv("DBPass", "postgres"),
		DBName: getEnv("DBName", "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
