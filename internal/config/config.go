package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func init() {
	path, err := filepath.Abs("../../.env")

	if err != nil {
		log.Fatal(err)
	}

	_ = godotenv.Load(path)
}

func GetVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return value, fmt.Errorf("provided key has no value %s", key)
	}

	return value, nil
}
