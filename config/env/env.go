package env

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
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
