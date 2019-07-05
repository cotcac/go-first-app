package categories

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
	stmt, err := db.Prepare("select * from categories limit ?,?")
	rows, err := stmt.Query(skip,perPage + 1)
	if err != nil {
		c.JSON(500, gin.H {
			"message":"server error",
		})
	}
	categories := make([]models.Category,0)
	for rows.Next() {
		var category models.Category
		rows.Scan(&category.Id, &category.Title)
		categories = append(categories,category)
	}
	fmt.Println(len(categories))
	var nextPage bool
	categoriesLen := len(categories)
	if categoriesLen > perPage {
		nextPage = true
		categories = categories[:len(categories)-1]// remote the last item.
	} else {
		nextPage = false
	}

	if err = rows.Err(); err != nil {
	    c.JSON(500, gin.H{"error:":"server error"})
	    return
	}
	c.JSON(200,gin.H{
		"page":page,
		"result":categories,
		"next": nextPage,
	})
}
