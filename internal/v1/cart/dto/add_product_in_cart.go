package dto

type AddProductInCart struct {
	ProductId uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}
