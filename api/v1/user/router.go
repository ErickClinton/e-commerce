package user

import (
	"eccomerce/internal/v1/cart"
	"eccomerce/internal/v1/middleware"
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
	service := user.NewService(userRepo, walletService, cartService)

	handler := NewHandler(service)

	publicRoutes := r.Group("/api/v1/register")
	{
		publicRoutes.POST("/", handler.Create)
	}

	protectedRoutes := r.Group("/api/v1/users")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.GET("/:id", middleware.AuthRoleMiddleware("admin"), handler.GetByID)
		protectedRoutes.GET("/current-user", handler.GetCurrentUser)
		protectedRoutes.PUT("/:id", handler.UpdateById)
		protectedRoutes.DELETE("/:id", middleware.AuthRoleMiddleware("admin"), handler.Delete)
	}
}
