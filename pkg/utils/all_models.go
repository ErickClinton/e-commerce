package utils

import (
	"eccomerce/internal/v1/entity"
)

func AllModels() []interface{} {
	return []interface{}{
		&entity.CartProduct{},
		&entity.Cart{},
		&entity.User{},
		&entity.Product{},
		&entity.Wallet{},
	}
}
