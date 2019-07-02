package users

import (
	"github.com/gin-gonic/gin"
)

func Single(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}
