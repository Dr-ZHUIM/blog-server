package Service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 该方法不会在博客项目中使用，在升级过后由 next.js 控制
func RedirectTo404(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/404")
}
