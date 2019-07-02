package users

import (
	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	// id := c.PostForm("id")
	name := c.PostForm("name")
	c.JSON(200, gin.H{
		"Insert": name,
	})
}
