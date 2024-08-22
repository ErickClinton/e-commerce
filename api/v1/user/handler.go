package user

import (
	"eccomerce/internal/v1/user/dto"
	"eccomerce/internal/v1/user/services"
	"eccomerce/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) create(c *gin.Context) {
	utils.Logger.Info().Msg("Start method create")

	var input dto.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method create %s", err.Error())
		return
	}

	if err := h.service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func (h *Handler) GetByID(c *gin.Context) {
	utils.Logger.Info().Msg("Start method GetByID")
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		utils.Logger.Error().Msgf("Error method GetByID %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) Update(c *gin.Context) {
	utils.Logger.Info().Msg("Start method UpdateUser")

	var input dto.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method UpdateUser %s", err.Error())
		return
	}

	if err := h.service.Update(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (h *Handler) UpdateById(c *gin.Context) {
	utils.Logger.Info().Msg("Start method UpdateUser")

	id, _ := strconv.Atoi(c.Param("id"))
	var input dto.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method UpdateUser %s", err.Error())
		return
	}

	if err := h.service.UpdateById(&input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (h *Handler) Delete(c *gin.Context) {
	utils.Logger.Info().Msg("Start method UpdateUser")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method UpdateUser %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}

func (h *Handler) Login(c *gin.Context) {
	utils.Logger.Info().Msg("Start method login")

	var input dto.LoginUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method create %s", err.Error())
		return
	}

	token, err := h.service.Login(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
