package dto

type CreateProductRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	UserId      uint    `json:"userId" `
}
