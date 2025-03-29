package app

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var (
	config *Config
	once   sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("error loading .env file: %s", err)
			return
		}

		config = &Config{
			Port:       getEnv("PORT"),
			DBHost:     getEnv("DB_HOST"),
			DBPort:     getEnv("DB_PORT"),
			DBUser:     getEnv("DB_USER"),
			DBPassword: getEnv("DB_PASSWORD"),
			DBName:     getEnv("DB_NAME"),
		}

	})

	return config
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("environment variable %s is not set", key))
}
