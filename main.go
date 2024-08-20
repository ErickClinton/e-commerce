package main

import (
	"eccomerce/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Inicia o controller que já contém as rotas
	controller.NewUserController(router)

	router.Run(":8080")
}
