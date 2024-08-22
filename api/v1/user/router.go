package user

import (
	"eccomerce/internal/v1/user/repository"
	"eccomerce/internal/v1/user/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	service := services.NewService(repo)
	handler := NewHandler(service)

	v1 := r.Group("/api/v1/users")
	{
		v1.POST("/", handler.create)
		v1.GET("/:id", handler.GetByID)
		v1.PUT("/:id", handler.UpdateById)
		v1.DELETE("/:id", handler.Delete)
		v1.POST("/login", handler.Login)
	}
}
