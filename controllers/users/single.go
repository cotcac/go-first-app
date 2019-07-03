package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

func Single(c *gin.Context) {
	type User struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}
	id := c.Param("id")
	var user User
	db:= models.DBConn()
	err := db.QueryRow("select * from users where id =?", id).Scan(
		&user.Id,
		&user.Name,
	)
	if err != nil {
		c.JSON(500, gin.H {
			"message":"server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": user,
	})
}
