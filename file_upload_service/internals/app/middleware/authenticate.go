package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	UserID string `headers:"UserId" binding:"required"`
}

type Authenticate struct {
}

func NewAuthenticate() *Authenticate {
	return &Authenticate{}
}

func (a Authenticate) Auth(c *gin.Context) {

	var Headers Headers

	if err := c.ShouldBindHeader(&Headers); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

}
