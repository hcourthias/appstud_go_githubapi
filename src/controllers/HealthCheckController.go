package controllers

import (
	"time"

	"appstud.com/github-core/src/models"

	"github.com/gin-gonic/gin"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getHealthCheck(c *gin.Context) {
	c.JSON(200, models.HealthCheckResponse{
		Name:    "github-api",
		Version: "1.0",
		Time:    makeTimestamp(),
	})
}

// HealthCheckController - Route controller
func HealthCheckController(engine *gin.Engine) {
	engine.GET("/api/healthcheck", getHealthCheck)
}
