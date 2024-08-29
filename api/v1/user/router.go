package user

import (
	"eccomerce/internal/v1/cart"
	"eccomerce/internal/v1/middleware"
	"eccomerce/internal/v1/user/repository"
	"eccomerce/internal/v1/user/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	cartRepository := cart.NewCartRepository(db)
	cartService := cart.NewCartService(cartRepository)
	service := services.NewService(repo, cartService)
	handler := NewHandler(service)

	publicRoutes := r.Group("/api/v1/register")
	{
		publicRoutes.POST("/", handler.create)
	}

	protectedRoutes := r.Group("/api/v1/users")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.POST("/", handler.create)
		protectedRoutes.GET("/:id", handler.GetByID)
		protectedRoutes.PUT("/:id", handler.UpdateById)
		protectedRoutes.DELETE("/:id", middleware.AuthRoleMiddleware("admin"), handler.Delete)
	}
}
