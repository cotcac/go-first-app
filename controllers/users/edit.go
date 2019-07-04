package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

func Edit(c *gin.Context) {
	id := c.Param("id")
	db:= models.DBConn()
	type User struct {
		Name string `form:"name" json:"name" binding:"required"`
	}
	var json User

	if err:= c.ShouldBindJSON(&json); err == nil {
		edit, err := db.Prepare("update users set name=? where id=" + id)
		if err != nil {
			c.JSON(500, gin.H{
				"message":err,
			})
			return
		}
		if json.Name == "" {
			c.JSON(400, gin.H{
				"message":"404",
			})
			return
		}
		edit.Exec(json.Name)
		c.JSON(200, gin.H{
			"Edit": "Edited",
		})

	}else {
		c.JSON(200, gin.H{"error": err.Error()})
	}

	
}
