package config

import "os"

// Config ...
type Config struct {
	AppPort           string
	AppHost           string
	AppEnv            string
	PgHost            string
	PgPort            string
	PgName            string
	PgUser            string
	PgPassword        string
	BasicAuthUsername string
	BasicAuthPassword string
}

// New set env value
func New() *Config {
	return &Config{
		AppPort:           getEnv("APP_PORT", "8080"),
		AppHost:           getEnv("APP_HOST", "localhost"),
		AppEnv:            getEnv("APP_ENV", "local"),
		PgHost:            getEnv("PG_HOST", "127.0.0.1"),
		PgPort:            getEnv("PG_PORT", "5432"),
		PgName:            getEnv("PG_NAME", "postgres"),
		PgUser:            getEnv("PG_USER", "root"),
		PgPassword:        getEnv("PG_PASSWORD", ""),
		BasicAuthUsername: getEnv("BASIC_AUTH_USERNAME", ""),
		BasicAuthPassword: getEnv("BASIC_AUTH_PASSWORD", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
