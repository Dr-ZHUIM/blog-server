package Service

import (
	CustomMysql "article-server/internal/app/db"
	DTOs "article-server/internal/app/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddArticle(c *gin.Context) {
	var newArticle DTOs.Article
	if err := c.BindJSON(&newArticle); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	CustomMysql.DB.Create(&newArticle)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "article created"})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article DTOs.Article
	CustomMysql.DB.Take(&article, id)
	if article.ID != 0 {
		CustomMysql.DB.Delete(&article, id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "article deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}

func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article DTOs.Article
	CustomMysql.DB.Take(&article, id)
	if article.ID != 0 {
		if err := c.BindJSON(&article); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		CustomMysql.DB.Save(&article)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "article updated"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}

func GetArticleList(c *gin.Context) {
	var articleList []DTOs.Article
	CustomMysql.DB.Find(&articleList)
	if len(articleList) > 0 {
		c.IndentedJSON(http.StatusOK, articleList)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "there is no article"})
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article DTOs.Article
	CustomMysql.DB.Take(&article, id)
	if article.ID != 0 {
		c.IndentedJSON(http.StatusOK, article)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}
