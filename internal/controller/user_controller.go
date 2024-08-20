package controller

import (
	"eccomerce/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(router *gin.Engine) {
	service := services.NewUserService()
	controller := &UserController{
		userService: service,
	}

	// Define a rota GET para listar usuários
	router.GET("/users", controller.ListUsers)

}

func (ctrl *UserController) ListUsers(c *gin.Context) {
	users := ctrl.userService.ListUsers()
	c.JSON(200, users)
}
