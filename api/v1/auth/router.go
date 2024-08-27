package auth

import (
	authServices "eccomerce/internal/v1/auth/services"
	userRepo "eccomerce/internal/v1/user/repository"
	userServices "eccomerce/internal/v1/user/services"
	walletRepo "eccomerce/internal/v1/wallet/repository"
	walletServices "eccomerce/internal/v1/wallet/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repoUser := userRepo.NewUserRepository(db)
	walletRepo := walletRepo.NewWalletRepository(db)
	walletService := walletServices.NewWalletService(walletRepo)
	userService := userServices.NewService(repoUser, walletService)
	authService := authServices.NewServiceAuth(userService)
	authHandler := NewHandlerAuth(authService)

	v1 := r.Group("/api/v1/auth")
	{
		v1.POST("/login", authHandler.Login)
	}
}
