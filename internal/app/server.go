package articleServer

import (
	Controller "article-server/internal/app/controllers"
	Log "article-server/internal/app/log"
	"io"

	"github.com/gin-gonic/gin"
)

func Launch() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(Log.GinLogFile)
	Controller.GetServices()
}
