package product

import (
	"eccomerce/internal/v1/product/dto"
	"eccomerce/internal/v1/product/services"
	"eccomerce/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service services.ProductService
}

func NewHandler(service services.ProductService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) create(c *gin.Context) {
	utils.Logger.Info().Msg("Start method create")
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	userIdPtr, ok := userId.(*uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	var input dto.CreateProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method create %s", err.Error())
		return
	}

	input.UserId = *userIdPtr
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
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		utils.Logger.Error().Msgf("Error method GetByID %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) UpdateById(c *gin.Context) {
	utils.Logger.Info().Msg("Start method UpdateUser")
	id, _ := strconv.Atoi(c.Param("id"))

	var input dto.UpdateProductRequest
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
