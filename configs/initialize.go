package configs

import (
	controllers "eccomerce/internal/controller"
	"eccomerce/internal/repository"
	"eccomerce/internal/services"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {

	db := SetupDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	router := gin.Default()

	controllers.ConfigureUserRoutes(router, userService)

	return router
}
