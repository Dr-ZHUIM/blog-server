package Controller

import (
	Service "article-server/internal/app/services"

	"github.com/gin-gonic/gin"
)

func GetServices() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.POST("/getArticleList", Service.GetArticleList)
	router.POST("/addArticle", Service.AddArticle)
	router.POST("/getArticle/:id", Service.GetArticle)
	router.POST("/login", Service.Login)
	router.Run("localhost:8080")
}
