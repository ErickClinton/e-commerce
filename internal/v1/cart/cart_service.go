package cart

import (
	"eccomerce/internal/v1/cart/dto"
	"eccomerce/internal/v1/entity"
)

type CartService interface {
	AddProductById(addProductDto *dto.AddProductInCart, userId uint) *entity.CartProduct
	GetCartWithProductByUserId(cartId uint) (*entity.Cart, error)
}

type cartService struct {
	repository CartRepository
}

func NewCartService(repository CartRepository) CartService {
	return &cartService{repository: repository}
}

func (service cartService) AddProductById(productCartDto *dto.AddProductInCart, userId uint) *entity.CartProduct {
	cartId, _ := service.repository.getCartByUserId(userId)
	existProduct, _ := service.repository.GetCartProductByCartIdAndProductId(cartId.Id, productCartDto.ProductId)
	if existProduct != nil {
		cartProduct := &entity.CartProduct{
			CartId:    cartId.Id,
			ProductId: existProduct.ProductId,
			Quantity:  productCartDto.Quantity,
		}
		product, _ := service.repository.UpdateProduct(cartProduct)
		return product
	} else {
		cartProduct := &entity.CartProduct{
			CartId:    cartId.Id,
			ProductId: productCartDto.ProductId,
			Quantity:  productCartDto.Quantity,
		}
		product, _ := service.repository.AddProduct(cartProduct)
		return product
	}

}

func (service cartService) GetCartWithProductByUserId(userId uint) (*entity.Cart, error) {
	return service.repository.getCartWithProductByUserId(userId)
}
