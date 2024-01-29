package handlers

import (
	fr "file-upload-service/internals/app/form-request"
	rv "file-upload-service/internals/app/request-validator"
	"file-upload-service/internals/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FileUploadHandler struct {
	fus *services.FileUploadService
}

func NewFileUploadHandler(fus *services.FileUploadService) *FileUploadHandler {
	return &FileUploadHandler{fus: fus}
}

func (fuh FileUploadHandler) Store(c *gin.Context) {

	var request fr.FileUploadRequest

	// validate request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": rv.Validator{Err: err.(validator.ValidationErrors)}.Errors(),
		})
		return
	}

	/// process request
	uploadID, object_key, err := fuh.fus.InitiateMultipartUpload(request, c.GetHeader("UserId"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "upload initiated",
		"uploadID":   uploadID,
		"object_key": object_key,
	})

}

func (fuh FileUploadHandler) Update(c *gin.Context) {

	var request fr.FilePartUploadRequest

	// validate request
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	/// process request
	etag, err := fuh.fus.SavePartUpload(request, c.GetHeader("UserId"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "uploaded part",
		"upload_part_etag": etag,
	})

}

func (fuh FileUploadHandler) CompleteMultiUpload(c *gin.Context) {

	var request fr.CompleteMultiUploadRequest

	// validate request
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	/// process complete multi upload request
	_, err := fuh.fus.CompleteMultiUpload(request, c.GetHeader("UserId"))
	if err != nil && err.Error() == "records not found" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "error completing multipart uploads. Reason: could not retreive uploaded parts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload completed",
	})
}

func (fuh FileUploadHandler) AbortMultiUpload(c *gin.Context) {

	if c.Query("upload_id") == "" || c.Query("object_key") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "object key and upload id is required",
		})
		return
	}

	err := fuh.fus.AbortMultiUpload(c.Query("upload_id"), c.Query("object_key"), c.GetHeader("UserId"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload aborted",
	})
}
