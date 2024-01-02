package Service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// TODO
	content, err := c.MultipartForm()
	if err == nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "upload file success", "status": 200})
		fmt.Println(content)
		return
	}
	c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "upload file error", "status": "upload_file_error"})
}
