package fr

import "mime/multipart"

type FilePartUploadRequest struct {
	File      *multipart.FileHeader `json:"file" binding:"required"`
	UploadId  string                `json:"upload_id" binding:"required"`
	ObjectKey string                `json:"object_key" binding:"required"`
}
