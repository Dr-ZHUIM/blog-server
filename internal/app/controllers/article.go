package articleController

import (
	articleService "article-server/internal/app/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ArticleServices() {
	fmt.Println("---------------------------------------------------")
	router := gin.Default()
	router.POST("/getArticleList", articleService.GetArticleList)
	router.POST("/addArticle", articleService.AddArticle)
	router.POST("/getArticle/:id", articleService.GetArticle)
	router.Run("localhost:8080")
	fmt.Println("server article opened")
	fmt.Println("---------------------------------------------------")
}
