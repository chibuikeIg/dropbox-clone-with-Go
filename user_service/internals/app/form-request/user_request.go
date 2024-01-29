package fr

type UserRequest struct {
	ID       string `json:"id" binding:"required,uuid"`
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
