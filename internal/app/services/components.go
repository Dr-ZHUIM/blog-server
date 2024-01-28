package Service

import (
	CustomMysql "article-server/internal/app/db"
	DTOs "article-server/internal/app/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComponent(c *gin.Context) {
	var newComponent DTOs.Component
	if err := c.BindJSON(&newComponent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	CustomMysql.DB.Create(&newComponent)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "article created"})
}

func BindComponent(c *gin.Context) {
	type RequestType struct {
		ArticleId  int   `json:"articleId"`
		Components []int `json:"components"`
	}
	var request RequestType
	var targetArticle DTOs.Article
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	CustomMysql.DB.Take(&targetArticle, request.ArticleId)
	if targetArticle.ID != 0 {
		for _, componentId := range request.Components {
			var targetComponent DTOs.Component
			CustomMysql.DB.Take(&targetComponent, componentId)
			if targetComponent.ID != 0 {
				targetArticle.Components = append(targetArticle.Components, targetComponent)
			}
		}
		CustomMysql.DB.Save(&targetArticle)
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "article binded"})
	}
}

func GetComponent(c *gin.Context) {
	var target DTOs.Component
	id, _ := strconv.Atoi(c.Param("id"))
	CustomMysql.DB.Take(&target, id)
	if target.ID != 0 {
		c.IndentedJSON(http.StatusOK, target)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "article not found", "code": 404, "data": nil})
}
