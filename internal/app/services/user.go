package Service

import (
	DTOs "article-server/internal/app/dtos"
	jwtFunc "article-server/pkg/Jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []DTOs.User{
	{Username: "admin", Password: "123456", Id: 1},
	{Username: "root", Password: "123456", Id: 2},
}

func Login(c *gin.Context) {
	var json LoginData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "struct error", "status": "bad_request"})
		return
	}
	for _, user := range users {
		if user.Username == json.Username {
			if user.Password == json.Password {
				Token, err := jwtFunc.GenerateJWTToken()
				if err != nil {
					c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "generate token error", "status": "generate_token_error"})
					return
				}
				c.IndentedJSON(http.StatusOK, gin.H{"message": "login success", "status": "success", "token": Token})
				return
			}
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "found wrong password", "status": "wrong_password"})
			return
		}
	}
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "user not found", "status": "not_found"})
}
