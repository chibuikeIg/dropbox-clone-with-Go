package main

import (
	"file-upload-service/internals/adapters/handlers"
	"file-upload-service/internals/adapters/repository"
	"file-upload-service/internals/app/middleware"
	"file-upload-service/internals/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {

	// start aws s3 client
	bucketName := "myawsbucket899" // get from env
	awsS3 := repository.NewAwsS3(&bucketName)
	dynamoDB := repository.NewDynamoDbRepository()

	// load services
	middleware := middleware.NewAuthenticate()

	// service routes
	{
		fus := services.NewFileUploadService(awsS3, dynamoDB)
		fuh := handlers.NewFileUploadHandler(fus)
		api := r.Group("/api/v1", middleware.Auth)
		api.POST("/uploads/init", fuh.Store)
		api.PUT("/uploads/update", fuh.Update)
		api.POST("/uploads/complete", fuh.CompleteMultiUpload)
		api.DELETE("/uploads/abort", fuh.AbortMultiUpload)
	}

	/// service ping
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
