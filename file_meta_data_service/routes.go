package main

import (
	"filemetadata-service/internals/adapters/handlers"
	"filemetadata-service/internals/adapters/repository"
	"filemetadata-service/internals/app/middleware"
	"filemetadata-service/internals/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {

	// start dynamodb & redis client
	dynamoDB := repository.NewDynamoDbRepository()
	redisDB := repository.NewRedisDBRepository()

	// middleware
	middleware := middleware.NewAuthenticate(dynamoDB)

	// load services & handler
	// folder service & handler
	fs := services.NewFolderService(dynamoDB, redisDB)
	fh := handlers.NewFolderHandler(fs)

	// file service & handler
	fileService := services.NewFileService(dynamoDB, redisDB)
	fileHandler := handlers.NewFileHandler(fileService)

	// service routes
	{
		api := r.Group("/api/v1", middleware.Auth)

		// folder endpoints
		folder := api.Group("/folders")
		folder.GET("/", fh.Index)
		folder.POST("/", fh.Store)
		folder.GET("/:folder", fh.LatestSnapshot)
		folder.POST("/:folder", fh.Update)
		folder.DELETE("/:folder", fh.Delete)

		// file endpoints
		file := api.Group("/files")
		file.GET("/", fileHandler.Index)
		file.POST("/", fileHandler.Store)
		file.GET("/:file", fileHandler.Show)
		file.DELETE("/:file", fileHandler.Delete)
	}

	/// service ping
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
