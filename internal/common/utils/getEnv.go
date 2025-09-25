package utils

import (
	"os"
)

// Get env variables
func GetEnv(key, defaultValue string) string {
	if str := os.Getenv(key); str != "" {
		return str
	}
	return defaultValue

}
