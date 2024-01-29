package fr

type FileRequest struct {
	Name     string `json:"filename" binding:"required"`
	FolderId string `json:"folder_id" binding:"omitempty,uuid"`
}
