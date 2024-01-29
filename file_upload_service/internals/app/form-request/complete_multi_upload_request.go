package fr

type CompleteMultiUploadRequest struct {
	UploadId  string `json:"upload_id" binding:"required"`
	ObjectKey string `json:"object_key" binding:"required"`
}
