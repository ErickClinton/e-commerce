package wallet

import (
	"eccomerce/internal/v1/middleware"
	"eccomerce/internal/v1/wallet/repository"
	"eccomerce/internal/v1/wallet/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewWalletRepository(db)
	service := services.NewWalletService(repo)
	handler := NewHandler(service)

	protectedRoutes := r.Group("/api/v1/wallet")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.GET("/:id", handler.GetByID)
		protectedRoutes.PUT("/:id", handler.UpdateById)
	}
}
