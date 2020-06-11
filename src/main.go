package main

import (
	"appstud.com/github-core/src/controllers"
	"appstud.com/github-core/src/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()
	controllers.HelloWorldController(r)  // controller for /api/hello
	controllers.HealthCheckController(r) // controller for /api/healthcheck
	controllers.McFlyController(r)       // controller for /api/timemachine/logs/mcfly (easter egg)
	controllers.AuthentificationController(r)
	controllers.GithubController(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
