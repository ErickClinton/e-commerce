package main

import (
	"eccomerce/api/v1/user"
	"eccomerce/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := configs.SetupDatabase()
	user.RegisterRoutes(r, db)

	r.Run(":8080")
}
