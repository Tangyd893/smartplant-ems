package config

import (
	"os"
	"strconv"
)

func GetDBUser() string {
	return getEnv("DB_USER", "spuser")
}

func GetDBPass() string {
	return getEnv("DB_PASS", "sp123456")
}

func GetDBHost() string {
	return getEnv("DB_HOST", "127.0.0.1:3306")
}

func GetDBName() string {
	return getEnv("DB_NAME", "smartplant_ems")
}

func GetJWTSecret() string {
	return getEnv("JWT_SECRET", "smartplant-ems-secret-key-2026")
}

func GetMaxIdleConns() int {
	return getEnvInt("MAX_IDLE_CONNS", 10)
}

func GetMaxOpenConns() int {
	return getEnvInt("MAX_OPEN_CONNS", 100)
}

func GetCORSAllowedOrigin() string {
	return getEnv("CORS_ALLOWED_ORIGIN", "*")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}