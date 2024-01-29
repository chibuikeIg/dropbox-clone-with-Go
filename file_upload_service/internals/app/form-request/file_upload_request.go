package fr

type FileUploadRequest struct {
	FileName    string `json:"filename" binding:"required"`
	FolderName  string `json:"folder_name"`
	ContentType string `json:"content_type" binding:"required"`
}
