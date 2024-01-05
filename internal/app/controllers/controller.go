package Controller

import (
	Service "article-server/internal/app/services"

	"github.com/gin-gonic/gin"
)

func GetServices() {
	router := gin.Default()
	aritcleRouters(router)
	staticRouters(router)
	router.POST("/login", Service.Login)
	router.Run("localhost:8080")
}
