package categories

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

func Insert(c *gin.Context) {
	db:= models.DBConn()
	type Create struct {
		Name string `json:"title" binding:"required"`
	}
	var json Create
	if err :=c.ShouldBindJSON(&json); err == nil {
		insert, err := db.Prepare("insert into categories(title) values (?)",)
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
		insert.Exec(json.Name)

		c.JSON(200, gin.H{
			"message": "inserted",
		})
	} else {
		c.JSON(500, gin.H{
			"error":err.Error(),
		})
	}
}
