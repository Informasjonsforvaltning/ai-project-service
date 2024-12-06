package env

import (
	"os"
	"strings"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func CorsOriginPatterns() []string {
	return strings.Split(getEnv("CORS_ORIGIN_PATTERNS", "*"), ",")
}

type Paths struct {
	Ping  string
	Ready string
	Data  string
}

var PathValues = Paths{
	Ping:  "/ping",
	Ready: "/ready",
	Data:  "/projects",
}
