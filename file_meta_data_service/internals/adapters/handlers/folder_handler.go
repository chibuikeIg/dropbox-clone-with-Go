package handlers

import (
	fr "filemetadata-service/internals/app/form-request"
	rv "filemetadata-service/internals/app/request-validator"
	"filemetadata-service/internals/core/domain"
	"filemetadata-service/internals/core/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FolderHandler struct {
	fs *services.FolderService
}

func NewFolderHandler(fs *services.FolderService) *FolderHandler {
	return &FolderHandler{fs: fs}
}

func (fh FolderHandler) Index(c *gin.Context) {

	folders := fh.fs.Folders(c.GetHeader("UserId"))

	if len(folders) == 0 {
		folders = []domain.Folder{}
	}

	c.JSON(http.StatusOK, gin.H{
		"folders": folders,
	})
}

func (fh FolderHandler) Store(c *gin.Context) {

	var request fr.FolderRequest

	// validate request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": rv.Validator{Err: err.(validator.ValidationErrors)}.Errors(),
		})
		return
	}

	// if folder already exists rename here
	// then
	// store records
	folder, err := fh.fs.Store(request, c.GetHeader("UserId"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurred",
		})
		log.Printf("unable to create new folder %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"folder": folder,
	})
}

func (fh FolderHandler) LatestSnapshot(c *gin.Context) {

	files, err := fh.fs.GetFolderFiles(c.Param("folder"), c.GetHeader("UserId"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "record not found",
		})
		return
	}

	if len(files) == 0 {
		files = []domain.File{}
	}

	c.JSON(http.StatusOK, gin.H{
		"folderId": c.Param("folder"),
		"fileList": files,
	})
}

func (fh FolderHandler) Update(c *gin.Context) {

	var request fr.FolderRequest

	// validate request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": rv.Validator{Err: err.(validator.ValidationErrors)}.Errors(),
		})
		return
	}

	// update record
	folder, err := fh.fs.Update(request, c.GetHeader("UserId"), c.Param("folder"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurred",
		})
		log.Printf("unable to update folder %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"folder": folder,
	})

}

func (fh FolderHandler) Delete(c *gin.Context) {

	err := fh.fs.Delete(c.GetHeader("UserId"), c.Param("folder"))

	if err != nil && err.Error() != "record not found" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "folder not found",
		})
		return
	}

	if err != nil && err.Error() != "internal server error" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "folder deleted",
	})
}
