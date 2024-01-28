package Service

import (
	CustomMysql "article-server/internal/app/db"
	DTOs "article-server/internal/app/dtos"
	"fmt"
	"log"
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
	var article DTOs.Article
	if err := c.BindJSON(&article); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	CustomMysql.DB.Where("1 = 1")
	query := CustomMysql.DB.Model(&DTOs.Article{})
	if article.Category != "" {
		query = query.Where("category = ?", article.Category)
	}
	if article.Layer != "" {
		fmt.Println("layer: ", article.Layer)
		query = query.Where("layer = ?", article.Layer)
	}
	query.Order("createdAt DESC").Find(&articleList)
	if len(articleList) > 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "status": 200, "data": articleList})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no article", "status": 404, "data": articleList})
}

func GetArticleListReview(c *gin.Context) {
	type RequestDto struct {
		Layer    string `json:"layer"`
		Category string `json:"category"`
	}
	var articleList []DTOs.ArticleReview
	var request RequestDto
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// CustomMysql.DB.Where("1 = 1")
	query := CustomMysql.DB.Model(&DTOs.Article{})
	if request.Layer != "" {
		query = query.Where("layer = ?", request.Layer)
	}
	if request.Category != "" {
		query = query.Where("category = ?", request.Category)
	}
	result := query.Order("createdat DESC").Find(&articleList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	if len(articleList) > 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "status": 200, "data": articleList})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no article", "status": 404, "data": articleList})
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article DTOs.Article
	CustomMysql.DB.Preload("Components").Take(&article, id)
	if article.ID != 0 {
		c.IndentedJSON(http.StatusOK, article)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "article not found", "code": 404, "data": nil})
}

func GetArticleByTitle(c *gin.Context) {
	title := c.Param("title")
	var article DTOs.Article
	CustomMysql.DB.Preload("Components").Where("title = ?", title).Take(&article)
	if article.ID != 0 {
		c.IndentedJSON(http.StatusOK, article)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "article not found", "code": 404, "data": nil})
}

func SearchArticle(c *gin.Context) {
	var articleList []DTOs.Article
	var keyword DTOs.KeyWord
	if err := c.BindJSON(&keyword); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	query := CustomMysql.DB.Model(&DTOs.Article{})
	if keyword.Keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword.Keyword+"%")
	}
	query.Order("createdAt DESC").Find(&articleList)
	if len(articleList) > 0 {
		c.IndentedJSON(http.StatusOK, articleList)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "there is no article"})
}
