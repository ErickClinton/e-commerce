package dto

type UpdateUserRequest struct {
	Username *string `json:"username" `
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password" validate:"omitempty,min=6"`
	Role     *string `json:"role" `
}
