package dto

type UpdateProductRequest struct {
	Title       string  `json:"title" `
	Description string  `json:"description"`
	Price       float64 `json:"price" `
}
