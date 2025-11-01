package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func findEnvFile() string {
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)

	for i := 0; i < 5; i++ { // search up to 5 levels up
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}
		dir = filepath.Dir(dir)
	}
	return ""
}

func init() {
	envPath := findEnvFile()
	if envPath == "" {
		log.Fatal(".env file not found")
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Printf(".env loaded from %s", envPath)
}

func GetVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return value, fmt.Errorf("provided key has no value %s", key)
	}

	return value, nil
}
