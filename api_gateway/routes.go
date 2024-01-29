package main

import (
	"api-gateway/internals/adapters/handlers"
	"api-gateway/internals/adapters/repository"
	"api-gateway/internals/app/middleware"
	"api-gateway/internals/core/services"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {

	// start dynamodb client
	dynamoDB := repository.NewDynamoDbRepository()

	// start services
	{
		reg_service := services.NewRegistrationService(dynamoDB)
		login_service := services.NewLoginService(dynamoDB)
		rh := handlers.NewRegistrationHandler(reg_service)
		lh := handlers.NewLoginHandler(login_service)
		api := r.Group("/api")
		api.POST("/register", rh.Store)
		api.POST("/login", lh.Login)
	}

	// route requests to appropriate services
	{
		srh := handlers.NewServiceRouterHandler()
		authorized := r.Group("/api", middleware.Auth)
		authorized.Any("/core/:service/*req_path", srh.RouteRequests)
	}

}
