package articles

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Single(c *gin.Context) {

	id := c.Param("id")
	var article models.Article
	db:= models.DBConn()
	stmt, err := db.Prepare("select articles.*, users.* from articles inner join users on articles.user_id = users.id where articles.id=?")
	err = stmt.QueryRow(id).Scan(
		&article.Id,
		&article.Title,
		&article.UserId,
		&article.User.Id,
		&article.User.Name,
		&article.User.Email,

	)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H {
			"message":"Not Found!",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": article,
	})
}
