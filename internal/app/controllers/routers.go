package Controller

import (
	Service "article-server/internal/app/services"

	"github.com/gin-gonic/gin"
)

func aritcleRouters(router *gin.Engine) {
	article := router.Group("article")
	article.POST("/getArticleList", Service.GetArticleList)
	article.POST("/addArticle", Service.AddArticle)
	article.POST("/searchArticle", Service.SearchArticle)
	article.POST("/getArticle/:id", Service.GetArticle)
	article.POST("/getArticleByTitle/:title", Service.GetArticleByTitle)
	article.POST("/deleteArticle/:id", Service.DeleteArticle)
	article.POST("/getArticleListByReview", Service.GetArticleListReview)
}

func staticRouters(router *gin.Engine) {
	static := router.Group("static")
	static.Static("/static", "./static")
	static.POST("/uploadFile", Service.UploadFile)
}

func componentRouters(router *gin.Engine) {
	component := router.Group("component")
	component.POST("/addComponent", Service.AddComponent)
	component.POST("/bindComponent", Service.BindComponent)
	component.POST("/getComponent/:id", Service.GetComponent)
}
