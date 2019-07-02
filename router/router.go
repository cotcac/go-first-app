package router

import (
	"../controllers/users"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r:= gin.Default()
	r.Static("/public","./public")
	router := r.Group("/api") 
	{
		router.GET("/users/ping", users.Ping)
		router.GET("/users/", users.List)
		router.POST("/users/", users.Insert)
		router.PATCH("/users/edit/:id", users.Edit)
		router.GET("/users/single/:id", users.Single)
		router.DELETE("/users/delete/:id", users.Delete)

	}
	return r
}

func Router() *gin.Engine {
	return setupRouter()
}