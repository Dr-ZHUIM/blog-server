package Utils

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 可以用 c.ShouldBindJSON(&obj) 代替
func BindJSON(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
