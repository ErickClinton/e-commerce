package main

import (
	"eccomerce/api/v1/product"
	"eccomerce/api/v1/user"
	"eccomerce/configs"
	"eccomerce/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := configs.SetupDatabase()
	user.RegisterRoutes(r, db)
	product.RegisterRoutes(r, db)
	utils.ConfigBasicLogger()
	r.Run(":8080")
}
