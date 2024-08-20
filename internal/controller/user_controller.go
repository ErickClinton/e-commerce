package controllers

import (
	"eccomerce/internal/services"

	"github.com/gin-gonic/gin"
)

type UserRoutesController struct {
	userService *services.UserService
}

func ConfigureUserRoutes(router *gin.Engine, userService *services.UserService) {
	controller := &UserRoutesController{
		userService: userService,
	}

	router.POST("/users", controller.CreateUser)
	router.GET("/users/:id", controller.GetUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}

func (c *UserRoutesController) CreateUser(ctx *gin.Context) {

}

func (c *UserRoutesController) GetUser(ctx *gin.Context) {

}

func (c *UserRoutesController) UpdateUser(ctx *gin.Context) {

}

func (c *UserRoutesController) DeleteUser(ctx *gin.Context) {

}
