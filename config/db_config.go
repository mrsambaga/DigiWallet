package config

import (
	"os"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var (
	ENV      = getENV("ENV", "testing") // testing as default to skip auth middleware during unit test
	AppName  = "sea-labs-library"
	DBConfig = dbConfig{
		Host:     getENV("HOST", "localhost"),
		User:     getENV("DB_USER", "radjasa.dzar"),
		Password: getENV("DB_PASS", "radjasa.dzar"),
		DBName:   getENV("DB_NAME", "wallet_starter_db"),
		Port:     getENV("DB_PORT", "5432"),
	}
)
