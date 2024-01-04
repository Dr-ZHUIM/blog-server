package Controller

import (
	Service "article-server/internal/app/services"

	"github.com/gin-gonic/gin"
)

func aritcleRouters(router *gin.Engine) {
	article := router.Group("article")
	article.POST("/getArticleList", Service.GetArticleList)
	article.POST("/addArticle", Service.AddArticle)
	article.POST("/getArticle/:id", Service.GetArticle)
}

func staticRouters(router *gin.Engine) {
	static := router.Group("static")
	static.Static("/static", "./static")
	static.POST("/uploadFile", Service.UploadFile)
}
