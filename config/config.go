package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Returns the environment value for the provided key
func Get(key string) string {
	envMap, mapErr := godotenv.Read(".env")
	if mapErr != nil {
		os.Exit(1)
	}
	return envMap[key]
}
