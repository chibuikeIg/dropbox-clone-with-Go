package middleware

import (
	"api-gateway/internals/app/generator"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	Authorization string `headers:"Authorization" binding:"required,startswith=Bearer "`
}

func Auth(c *gin.Context) {

	var Headers Headers

	if err := c.ShouldBindHeader(&Headers); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	authorizationArray := strings.Split(Headers.Authorization, " ")

	if len(authorizationArray) < 3 {

		if err := c.ShouldBindHeader(&Headers); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid authorization header provided"})
			return
		}
	}

	token := authorizationArray[len(authorizationArray)-1]

	parsedToken, err := generator.ParseToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid access token provided",
		})
		return
	}

	c.Request.Header.Add("UserId", parsedToken.UserID)

}
