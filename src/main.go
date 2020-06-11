package main

import (
	"appstud.com/github-core/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.HelloWorldController(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
