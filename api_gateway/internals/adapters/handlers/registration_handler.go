package handlers

import (
	"api-gateway/internals/app"
	consul_api "api-gateway/internals/app/consul"
	fr "api-gateway/internals/app/form-request"
	"api-gateway/internals/app/generator"
	rv "api-gateway/internals/app/request-validator"
	"api-gateway/internals/core/domain"
	"api-gateway/internals/core/services"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegistrationHandler struct {
	rs *services.RegistrationService
}

func NewRegistrationHandler(rs *services.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		rs: rs,
	}
}

// store user data
func (rh *RegistrationHandler) Store(c *gin.Context) {

	var userReq fr.RegistrationRequest

	// validate request
	if err := c.ShouldBindJSON(&userReq); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": rv.Validator{Err: err.(validator.ValidationErrors)}.Errors(),
		})
		return
	}

	// check if user exists
	if rh.rs.CheckUserExists(userReq.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("This email address %v, already exists in our database", userReq.Email),
		})
		return
	}

	// store new registered user
	user, err := rh.rs.Store(&domain.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to register user, please try again",
		})

		log.Printf("unable to create new user %v", err)
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

	// send created user record to user service
	url, _ := consul_api.ServiceDiscovery("user", "/user")

	headers := c.Request.Header
	headers.Add("UserId", user.ID)

	/// the reason for converting this user model (struct)
	/// to map is so the payload can contain the password
	/// as the password was omitted from been marshalled

	json, _ := json.Marshal(map[string]string{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	})

	body := bytes.NewReader(json)

	// send request
	app.SendRequest(c.Request.Method, url, body, headers)

	c.JSON(http.StatusOK, gin.H{
		"user":         user,
		"access_token": token,
	})
}
