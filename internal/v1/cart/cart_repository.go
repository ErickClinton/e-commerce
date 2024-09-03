package cart

import (
	"eccomerce/internal/v1/entity"
	"gorm.io/gorm"
)

type CartRepository interface {
	AddProduct(product *entity.CartProduct) (*entity.CartProduct, error)
	UpdateProduct(product *entity.CartProduct) (*entity.CartProduct, error)
	getCartWithProductByUserId(cartId uint) (*entity.Cart, error)
	getCartByUserId(cartId uint) (*entity.Cart, error)
	GetCartProductByCartIdAndProductId(cartId uint, productId uint) (*entity.CartProduct, error)
	Create(cart *entity.Cart) (*entity.Cart, error)
}
type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (c *cartRepository) Create(cart *entity.Cart) (*entity.Cart, error) {
	if err := c.db.Create(cart).Error; err != nil {
	}
	return cart, nil
}
func (c *cartRepository) AddProduct(product *entity.CartProduct) (*entity.CartProduct, error) {
	if err := c.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (c *cartRepository) UpdateProduct(product *entity.CartProduct) (*entity.CartProduct, error) {
	if err := c.db.Where("cart_id = ?", product.CartId).Updates(product).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil

}

func (c *cartRepository) GetCartProductByCartIdAndProductId(cartId uint, productId uint) (*entity.CartProduct, error) {
	var cartProduct entity.CartProduct

	if err := c.db.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&cartProduct).Error; err != nil {
		return nil, err
	}

	return &cartProduct, nil
}

func (c *cartRepository) getCartByUserId(userId uint) (*entity.Cart, error) {
	var cart entity.Cart
	if err := c.db.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}
func (c *cartRepository) getCartWithProductByUserId(userId uint) (*entity.Cart, error) {
	var cart entity.Cart

	if err := c.db.Preload("Products.Product").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}
