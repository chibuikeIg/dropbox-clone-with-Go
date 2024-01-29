package fr

type FolderRequest struct {
	Name string `json:"name" binding:"required"`
}
