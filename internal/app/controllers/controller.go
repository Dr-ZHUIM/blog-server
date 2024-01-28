package Controller

import (
	"article-server/internal/app/middleware"
	Service "article-server/internal/app/services"

	"github.com/gin-gonic/gin"
)

func GetServices() {
	router := gin.Default()
	router.Use(middleware.Cors())
	aritcleRouters(router)
	staticRouters(router)
	componentRouters(router)
	router.POST("/login", Service.Login)
	router.Run("localhost:8080")
}
