package Log

import (
	"io"

	"github.com/gin-gonic/gin"

	"os"
)

func LogTest() {
	file, _ := os.Create("gin.log")
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(file)
}
