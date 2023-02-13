package main

import (
	"github.com/Informasjonsforvaltning/ai-project-service/config"
)

func main() {
	config.LoggerSetup()

	router := config.SetupRouter()
	router.Run(":8080")
	router.Run(":9091")
}
