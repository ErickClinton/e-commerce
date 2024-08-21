package main

import (
	"eccomerce/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Inicia o controller que já contém as rotas
	router = configs.InitializeApp()

	router.Run(":8080")
}
