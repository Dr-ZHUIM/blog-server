package Controller

import (
	Log "article-server/internal/app/log"
	Service "article-server/internal/app/services"

	"github.com/gin-gonic/gin"
)

func GetServices() {
	Log.LogTest()
	router := gin.Default()
	aritcleRouters(router)
	staticRouters(router)
	router.POST("/login", Service.Login)
	router.Run("localhost:8080")
}
