package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"strconv"
	"fmt"
)

func List(c *gin.Context) {
	db:= models.DBConn()
	perPage := 5
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)// the err is necessary
	var skip int
	if page >0 {
		skip = (page - 1)* perPage
	}else {
		skip = 0
	}
	stmt, err := db.Prepare("select * from users limit ?,?")
	rows, err := stmt.Query(skip,perPage + 1)
	if err != nil {
		c.JSON(500, gin.H {
			"message":"server error",
		})
	}
	users := make([]models.User,0)
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name,&user.Email)
		users = append(users,user)
	}
	fmt.Println(len(users))
	var nextPage bool
	usersLen := len(users)
	if usersLen > perPage {
		nextPage = true
		users = users[:len(users)-1]// remote the last item.
	} else {
		nextPage = false
	}

	if err = rows.Err(); err != nil {
	    c.JSON(500, gin.H{"error:":"server error"})
	    return
	}
	c.JSON(200,gin.H{
		"page":page,
		"result":users,
		"next": nextPage,
	})
}
