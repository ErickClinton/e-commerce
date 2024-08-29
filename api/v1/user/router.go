package user

import (
	"eccomerce/internal/v1/cart"
	"eccomerce/internal/v1/middleware"
	"eccomerce/internal/v1/user/repository"
	"eccomerce/internal/v1/user/services"
	walletRepo "eccomerce/internal/v1/wallet/repository"
	walletServices "eccomerce/internal/v1/wallet/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)

	walletRepo := walletRepo.NewWalletRepository(db)
	walletService := walletServices.NewWalletService(walletRepo)

	cartRepository := cart.NewCartRepository(db)
	cartService := cart.NewCartService(cartRepository)
	service := services.NewService(userRepo, walletService, cartService)

	handler := NewHandler(service)

	publicRoutes := r.Group("/api/v1/register")
	{
		publicRoutes.POST("/", handler.Create)
	}

	protectedRoutes := r.Group("/api/v1/users")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.GET("/:id", handler.GetByID)
		protectedRoutes.PUT("/:id", handler.UpdateById)
		protectedRoutes.DELETE("/:id", middleware.AuthRoleMiddleware("admin"), handler.Delete)
	}
}
