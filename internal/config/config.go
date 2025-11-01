package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return value, errors.New("provided key has no value")
	}

	return value, nil
}
