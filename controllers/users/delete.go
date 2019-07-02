package users

import (
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Users",
	})
}
