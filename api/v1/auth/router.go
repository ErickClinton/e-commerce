package auth

import (
	"eccomerce/internal/v1/auth"
	"eccomerce/internal/v1/cart"
	"eccomerce/internal/v1/user"
	"eccomerce/internal/v1/wallet"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := user.NewUserRepository(db)

	walletRepo := wallet.NewWalletRepository(db)
	walletService := wallet.NewWalletService(walletRepo)

	cartRepository := cart.NewCartRepository(db)
	cartService := cart.NewCartService(cartRepository)
	userService := user.NewService(userRepo, walletService, cartService)

	authService := auth.NewServiceAuth(userService)
	authHandler := NewHandlerAuth(authService)

	v1 := r.Group("/api/v1/auth")
	{
		v1.POST("/login", authHandler.Login)
	}
}
