package controllers

import (
	"eccomerce/internal/models"
	"eccomerce/internal/services"
	"fmt"
	"net/http"

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
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Chegou!")
	err := c.userService.CreateUser(req.Username, req.Email, req.Password, req.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (c *UserRoutesController) GetUser(ctx *gin.Context) {

}

func (c *UserRoutesController) UpdateUser(ctx *gin.Context) {

}

func (c *UserRoutesController) DeleteUser(ctx *gin.Context) {

}
