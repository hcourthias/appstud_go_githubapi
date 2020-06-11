package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGitHubFeed(c *gin.Context) {
	var data []interface{}
	resp, err := http.Get("https://api.github.com/events")

	if err != nil {
		c.JSON(500, "Failed to retrieve Github API.")
		panic(err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}
	err = json.Unmarshal([]byte(responseBody), &data)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}
	c.JSON(200, data)

}

func getGithubUserInfo(c *gin.Context) {
	var data interface{}
	resp, err := http.Get("https://api.github.com/users/" + c.Param("login"))

	if err != nil {
		c.JSON(500, "Failed to retrieve Github API.")
		panic(err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}
	err = json.Unmarshal([]byte(responseBody), &data)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}
	c.JSON(200, data)
}

// GithubController - Route controller
func GithubController(engine *gin.Engine) {
	engine.GET("/api/github/feed", getGitHubFeed)
	engine.GET("/api/github/users/:login", getGithubUserInfo)

}
