package main

import (
	"eccomerce/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router = configs.InitializeApp()

	router.Run(":8080")
}
