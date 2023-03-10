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

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(Cors())

	InitializeRoutes(router)
	

	return router
}
