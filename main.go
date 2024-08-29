package main

import (
	"eccomerce/api/v1/auth"
	"eccomerce/api/v1/cart"
	"eccomerce/api/v1/product"
	"eccomerce/api/v1/user"
	"eccomerce/api/v1/wallet"
	"eccomerce/configs"
	"eccomerce/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	db := configs.SetupDatabase()
	user.RegisterRoutes(r, db)
	godotenv.Load()
	product.RegisterRoutes(r, db)
	cart.RegisterRoutes(r, db)
	auth.RegisterRoutes(r, db)
	wallet.RegisterRoutes(r, db)
	utils.ConfigBasicLogger()
	r.Run(":8080")
}
