package dto

type UpdateUserRequest struct {
	Username string `json:"username" `
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"min=6"`
	Role     string `json:"role" `
}
