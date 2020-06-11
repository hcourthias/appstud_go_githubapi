package controllers

import (
	"appstud.com/github-core/src/models"

	"github.com/gin-gonic/gin"
)

func getMcFly(c *gin.Context) {
	c.JSON(200, models.McFlyResponse{
		{Name: "My mom is in love with me",
			Version: "1.0",
			Time:    -446723100},
		{Name: "I go to the future and my mom end up with the wrong guy",
			Version: "2.0",
			Time:    1445470140},
		{Name: "I go to the past and you will not believe what happens next",
			Version: "3.0",
			Time:    -9223372036854775808},
	})
}

// McFlyController - Route controller
func McFlyController(engine *gin.Engine) {
	engine.GET("/api/timemachine/logs/mcfly", getMcFly)
}
