package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Single(c *gin.Context) {
	type User struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}
	id := c.Param("id")
	var user User
	db:= models.DBConn()
	stmt, err := db.Prepare("select * from users where id =?")
	err = stmt.QueryRow(id).Scan(
		&user.Id,
		&user.Name,
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H {
			"message":"Not Found!",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": user,
	})
}
