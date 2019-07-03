package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func List(c *gin.Context) {
	db:= models.DBConn()
	rows, err := db.Query("select * from users")
	if err != nil {
		c.JSON(500, gin.H {
			"message":"server error",
		})
	}

	users := make([]User,0)
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		users = append(users,user)
	}
	if err = rows.Err(); err != nil {
	    c.JSON(500, gin.H{"error:":"server error"})
	    return
	}
	c.JSON(200,users)
	defer db.Close()
}
