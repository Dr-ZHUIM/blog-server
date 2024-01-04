package Service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	// file, err := c.MultipartForm()
	if err == nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "upload file success", "status": 200})
		fmt.Println(file)
		c.SaveUploadedFile(file, "./static/upload/"+file.Filename)
		return
	}
	c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "upload file error", "status": "upload_file_error"})
}

func DownloadFile(c *gin.Context) {
	fileName := c.Param("filename")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.File("./static/upload/" + fileName)
}
