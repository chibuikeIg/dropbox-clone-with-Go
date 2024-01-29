package handlers

import (
	fr "api-gateway/internals/app/form-request"
	"api-gateway/internals/app/generator"
	rv "api-gateway/internals/app/request-validator"
	"api-gateway/internals/core/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginHandler struct {
	ls *services.LoginService
}

func NewLoginHandler(ls *services.LoginService) *LoginHandler {
	return &LoginHandler{
		ls: ls,
	}
}

func (lh LoginHandler) Login(c *gin.Context) {

	var loginReq fr.LoginRequest

	// validate request
	if err := c.ShouldBindJSON(&loginReq); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": rv.Validator{Err: err.(validator.ValidationErrors)}.Errors(),
		})
		return
	}

	user, err := lh.ls.Authenticate(loginReq.Email, loginReq.Password)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// generate access token
	token, err := generator.CreateAccessToken(user.ID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to generate access token",
		})

		log.Printf("unable to generate access token %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":         user,
		"access_token": token,
	})

}
