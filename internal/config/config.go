package config

import (
	"fmt"
	"os"
)

func GetVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return value, fmt.Errorf("provided key has no value %s", key)
	}

	return value, nil
}
