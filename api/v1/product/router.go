package product

import (
	"eccomerce/internal/v1/middleware"
	"eccomerce/internal/v1/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := product.NewProductRepository(db)
	service := product.NewProductService(repo)
	handler := NewHandler(service)

	v1 := r.Group("/api/v1/product")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.POST("/", handler.create)
		v1.GET("/:id", handler.GetByID)
		v1.PUT("/:id", handler.UpdateById)
		v1.DELETE("/:id", handler.Delete)
	}
}
