package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

func Insert(c *gin.Context) {
	db:= models.DBConn()
	type CreateUser struct {
		Name string `form:"name" json:"name" binding:"required"`
	}
	var json CreateUser
	if err :=c.ShouldBindJSON(&json); err == nil {
		insertUser, err := db.Prepare("insert into users(name) values (?)",)
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
		insertUser.Exec(json.Name)

		c.JSON(200, gin.H{
			"message": "inserted",
		})
	} else {
		c.JSON(500, gin.H{
			"error":err.Error(),
		})
	}
	defer db.Close()
}
