package util

import "os"

func GetConfig(key string, defaultValue string) string {
	keyValue := os.Getenv(key)
	if keyValue != "" {
		return keyValue
	}
	return defaultValue
}
