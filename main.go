package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/")

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
