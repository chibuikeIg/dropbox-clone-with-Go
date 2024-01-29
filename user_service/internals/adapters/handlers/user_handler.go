package handlers

import (
	"log"
	"net/http"
	fr "user-service/internals/app/form-request"
	"user-service/internals/core/domain"
	"user-service/internals/core/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	us *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (uh *UserHandler) Store(c *gin.Context) {
	var userReq fr.UserRequest

	// validate request

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// store new registered user
	user, err := uh.us.Store(&domain.User{
		ID:       userReq.ID,
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to store user information, please try again",
		})

		log.Printf("unable to store new user %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
