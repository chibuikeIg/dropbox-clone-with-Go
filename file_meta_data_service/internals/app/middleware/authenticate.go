package middleware

import (
	"filemetadata-service/internals/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	UserID string `headers:"UserId" binding:"required"`
}

type Authenticate struct {
	db ports.FileMetaDataDBRepository
}

func NewAuthenticate(db ports.FileMetaDataDBRepository) *Authenticate {
	return &Authenticate{db: db}
}

func (a Authenticate) Auth(c *gin.Context) {

	var Headers Headers

	if err := c.ShouldBindHeader(&Headers); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

}
