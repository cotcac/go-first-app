package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	db:= models.DBConn()
	delete, err := db.Prepare("DELETE FROM users WHERE id=" + id)
    if err != nil {
        c.JSON(500, gin.H{
			"message":err,
		})
		return
	}
	delete.Exec()
	c.JSON(200, gin.H{
		"Delete": "deleted",
	})
	defer db.Close()
}
