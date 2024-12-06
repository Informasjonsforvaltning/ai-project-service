package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     env.CorsOriginPatterns(),
		AllowMethods:     []string{"OPTIONS", "GET"},
		AllowHeaders:     []string{"*"},
		AllowWildcard:    true,
		AllowAllOrigins:  false,
		AllowCredentials: false,
		AllowFiles:       false,
		MaxAge:           1 * time.Hour,
	}))
	InitializeRoutes(router)

	return router
}
