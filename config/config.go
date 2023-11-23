package config

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"strconv"
)

// Returns the environment value for the provided key
func Get(key string) string {
	envMap, mapErr := godotenv.Read(".env")
	if mapErr != nil {
		os.Exit(1)
	}
	return envMap[key]
}

func StrToIntOrNil(s string) *int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &i
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
