package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
	APIServerPort     string
	APISecret         string
	CookieEncryptKey  string
	JWTSecretKey      string
	JWTTokenTTLHour   int
	DBMaxIdleConns    int
	DBMaxOpenConns    int
	DBConnMaxLifetime time.Duration
}

var AppConfig *Config

func LoadConfig(path string) (*Config, error) {

	if err := godotenv.Load(path + "/app.env"); err != nil {
		log.Println("No .env file found or error loading, relying on environment variables")
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "3499"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}
	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "100"))
	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "100"))
	connMaxLifetimeHours, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME_HOURS", "1"))

	AppConfig = &Config{
		DBHost:            getEnv("DB_HOST", "202.10.39.16"),
		DBPort:            dbPort,
		DBUser:            getEnv("DB_USER", "user_production"),
		DBPassword:        getEnv("DB_PASSWORD", ""),
		DBName:            getEnv("DB_NAME", "edunusa_main"),
		APIServerPort:     getEnv("API_SERVER_PORT", ""),
		APISecret:         getEnv("API_SECRET", ""), // Pastikan ini kuat di produksi!,
		DBMaxIdleConns:    maxIdleConns,
		DBMaxOpenConns:    maxOpenConns,
		DBConnMaxLifetime: time.Duration(connMaxLifetimeHours),
	}

	// // Validasi APISecret
	// if AppConfig.APISecret == "f8c1e2f3a4b5c6d7" || len(AppConfig.APISecret) < 32 {
	// 	log.Println("Warning: API_SECRET is weak or using default. Please set a strong secret for JWT.")
	// }

	return AppConfig, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
