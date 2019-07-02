package users

import (
	"github.com/gin-gonic/gin"
)

func Single(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single User",
	})
}
