package env

import "os"

func GetEnv(envVar, defaultVal string) string {
	if value, ok := os.LookupEnv(envVar); ok && value != "" {
		return value
	}
	return defaultVal
}
