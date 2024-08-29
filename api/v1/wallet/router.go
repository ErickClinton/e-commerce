package wallet

import (
	"eccomerce/internal/v1/middleware"
	"eccomerce/internal/v1/wallet"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := wallet.NewWalletRepository(db)
	service := wallet.NewWalletService(repo)
	handler := NewHandler(service)

	protectedRoutes := r.Group("/api/v1/wallet")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.GET("/:id", handler.GetByID)
		protectedRoutes.PUT("/:id", handler.UpdateById)
	}
}
