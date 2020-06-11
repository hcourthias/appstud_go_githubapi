package controllers

import (
	"appstud.com/github-core/src/models"

	"github.com/gin-gonic/gin"
)

func handle(c *gin.Context) {
	c.JSON(200, models.HelloWorldResponse{
		Hello: "world!",
	})
}

// HelloWorldController - Route controller
func HelloWorldController(engine *gin.Engine) {
	engine.GET("/api/hello", handle)
}
