package go_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

var articles []Article

func main() {
	r := gin.Default()

	// get请求获取文章列表接口
	r.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, articles)
	})

	// post请求创建文章接口
	r.POST("/articles", func(c *gin.Context) {
		var article Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		articles = append(articles, article)
		c.JSON(http.StatusCreated, article)
	})

	r.Run(":8080")
}
