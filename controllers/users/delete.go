package users

import (
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"Delete": id,
	})
}
