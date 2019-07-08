package articles

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
	stmt, err := db.Prepare(`
	select articles.id, articles.title, users.id, users.name 
	from articles 
	inner join users on articles.user_id = users.id 
	limit ?,?`)
	rows, err := stmt.Query(skip,perPage + 1)
	if err != nil {
		c.JSON(500, gin.H {
			"message":"server error",
		})
	}
	fmt.Print(rows)
	articles := make([]models.Articles,0)
	for rows.Next() {
		article := models.Articles{}
		err := rows.Scan(&article.Id, &article.Title,&article.User.Id,&article.User.Name)
		if err != nil {
			panic(err)
		}
		articles = append(articles,article)
	}
	fmt.Println(len(articles))
	var nextPage bool
	theLen := len(articles)
	if theLen > perPage {
		nextPage = true
		articles = articles[:len(articles)-1]// remote the last item.
	} else {
		nextPage = false
	}

	if err = rows.Err(); err != nil {
	    c.JSON(500, gin.H{"error:":"server error"})
	    return
	}
	c.JSON(200,gin.H{
		"page":page,
		"result":articles,
		"next": nextPage,
	})
}
