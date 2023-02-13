package config

import (
	"github.com/gin-gonic/gin"

	"github.com/Informasjonsforvaltning/ai-project-service/config/env"
	"github.com/Informasjonsforvaltning/ai-project-service/handlers"
)

func InitializeRoutes(e *gin.Engine) {
	e.SetTrustedProxies(nil)
	e.GET(env.PathValues.Ping, handlers.PingHandler())
	e.GET(env.PathValues.Ready, handlers.ReadyHandler())
	e.GET(env.PathValues.Data, handlers.GetCSVFile())
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	InitializeRoutes(router)
	return router
}
