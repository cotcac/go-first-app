package users

import (
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Edit Users",
	})
}
