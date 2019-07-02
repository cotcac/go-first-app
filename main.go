package main

import (
	"./controllers/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// test api

	router.GET("/users/ping", users.Ping)
	router.GET("/users/", users.List)
	router.POST("/users/", users.Insert)
	router.PATCH("/users/edit/:id", users.Edit)
	router.GET("/users/single/:id", users.Single)
	router.DELETE("/users/delete/:id", users.Delete)

	router.Run() // listen and serve on 0.0.0.0:8080
}
