package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// test api
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// get params
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// post request
	router.POST("post", func(c *gin.Context) {
		// id := c.PostForm("id")
		name := c.PostForm("name")
		c.String(http.StatusOK, "Hello %s", name)

	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
