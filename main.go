package main

import (
	"net/http"
	"./controllers"
	"./controllers/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// test api
	router.GET("/ping", controllers.Ping)
	router.GET("/pong", controllers.Pong)
	router.GET("/pung", controllers.Pung)
	router.GET("/users/ping", users.Ping)
	router.GET("/users/", users.List)
	router.POST("/users/", users.Insert)
	router.PATCH("/users/edit/:id", users.Edit)
	router.GET("/users/single/:id", users.Single)
	router.DELETE("/users/delete/:id", users.Delete)

	// post request
	router.POST("post", func(c *gin.Context) {
		// id := c.PostForm("id")
		name := c.PostForm("name")
		c.String(http.StatusOK, "Hello %s", name)

	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
