package main

import (
	"net/http"
	"user-service/internals/adapters/handlers"
	"user-service/internals/adapters/repository"
	"user-service/internals/app/middleware"
	"user-service/internals/core/services"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {

	// start dynamodb client
	dynamoDB := repository.NewDynamoDbRepository()

	// load services &
	us := services.NewUserService(dynamoDB)
	middleware := middleware.NewAuthenticate(dynamoDB)

	// service routes
	{
		uh := handlers.NewUserHandler(us)
		api := r.Group("/api/v1", middleware.Auth)
		api.POST("/user", uh.Store)
	}

	/// service ping
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
