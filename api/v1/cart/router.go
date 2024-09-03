package cart

import (
	"eccomerce/internal/v1/cart"
	"eccomerce/internal/v1/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := cart.NewCartRepository(db)

	service := cart.NewCartService(repo)
	handler := NewHandler(service)
	protectedRoutes := r.Group("/api/v1/cart")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.POST("/", handler.addProductInCart)
		protectedRoutes.GET("/", handler.get)
		protectedRoutes.GET("/total-value", handler.totalValue)
	}
}
