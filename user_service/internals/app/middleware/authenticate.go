package middleware

import (
	"net/http"
	"user-service/internals/core/ports"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	UserID string `headers:"UserId" binding:"required"`
}

type Authenticate struct {
	db ports.UserDBRepository
}

func NewAuthenticate(db ports.UserDBRepository) *Authenticate {
	return &Authenticate{db: db}
}

func (a Authenticate) Auth(c *gin.Context) {

	var Headers Headers

	if err := c.ShouldBindHeader(&Headers); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

}
