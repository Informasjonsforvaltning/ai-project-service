package handlers

import (
	"net/http"

	"github.com/Informasjonsforvaltning/ai-project-service/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetCSVFile() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		data, err := service.ReadCSVData("ai_projects_norwegian_state - Oversatt_v1.csv")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			logrus.Info("Error reading CSV file: %s", err)
			return
		}
		c.JSON(http.StatusOK, data)
	}
}
