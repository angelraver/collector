package config

import (
	"net/http"
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

func SetCookie(w http.ResponseWriter, name, value string, maxAge int) {
	cookie := &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: maxAge, // MaxAge is in seconds
		// Other cookie options if needed, e.g., Domain, Path, etc.
	}
	http.SetCookie(w, cookie)
}
