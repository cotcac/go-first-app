package articles

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Insert(c *gin.Context) {
	db:= models.DBConn()
	type CreateArticle struct {
		Title string `json:"title" binding:"required,min=3,max=15"`
		UserId int `json:"user_id" binding:"required"`
	}  
	var json CreateArticle
	if err:= c.ShouldBindJSON(&json); err!=nil {
		fmt.Println(err)
		c.JSON(422, gin.H{"error1": err.Error()})
		return
	}
	// article := models.Article{Title:json.Title, UserId:json.UserId}
	insertArticle, err := db.Prepare("insert into articles(title,user_id) values (?,?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message3":"error 1",
		})
		return
	}
	res, err := insertArticle.Exec(json.Title, json.UserId)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message4": res,
	})
		
}
