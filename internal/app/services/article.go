package Service

import (
	DTOs "article-server/internal/app/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

var articles = []DTOs.Article{
	{ID: "1", Title: "Test article1", Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Alias quisquam blanditiis vitae consequuntur quia recusandae nobis, a quasi aliquid consequatur. Magni incidunt libero necessitatibus illo repellat error fuga, a, numquam dicta porro earum. Nulla sit minus blanditiis hic libero, et ad sunt suscipit quis saepe sed explicabo, id autem, quas eaque harum dolores! Delectus ab enim modi labore nisi dolorem voluptatum, quo temporibus a incidunt unde debitis eum quaerat culpa similique totam at odio. Non, quasi vero asperiores earum at molestiae laboriosam a culpa omnis commodi minus ipsum ipsa aspernatur, provident minima rerum ut praesentium delectus beatae fugiat quos quibusdam? Corporis dolor dolores sunt accusantium voluptatibus nulla doloribus sapiente! Facere molestiae corporis ad sed magni eveniet, mollitia exercitationem adipisci sapiente voluptate ullam eaque.", Artist: "Dr. Zhuim"},
	{ID: "2", Title: "Test article2", Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Illo ducimus, possimus architecto accusamus, beatae eos mollitia voluptates perferendis earum vero ex eligendi in nisi neque quo! Sed tempore provident exercitationem!", Artist: "Dr. Zhuim"},
}

func GetArticleList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

func AddArticle(c *gin.Context) {
	var newArticle DTOs.Article
	if err := c.BindJSON(&newArticle); err != nil {
		return
	}
	articles = append(articles, newArticle)
	c.IndentedJSON(http.StatusCreated, newArticle)
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	for _, a := range articles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}
