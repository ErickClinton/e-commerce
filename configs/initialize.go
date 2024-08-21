package configs

import (
	controllers "eccomerce/internal/controller"
	"eccomerce/internal/repository"
	"eccomerce/internal/services"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {
	// Configurar o banco de dados
	db := SetupDatabase()

	// Inicializar o repositório e o serviço
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Configurar o roteador
	router := gin.Default()

	// Configurar as rotas de usuário
	controllers.ConfigureUserRoutes(router, userService)

	return router
}
