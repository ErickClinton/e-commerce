package cart

import (
	"eccomerce/internal/v1/cart"
	"eccomerce/internal/v1/cart/dto"
	"eccomerce/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service cart.CartService
}

func NewHandler(service cart.CartService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) addProductInCart(c *gin.Context) {
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
	var input dto.AddProductInCart
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method create %s", err.Error())
		return
	}

	if err := h.service.AddProductById(&input, *userIdPtr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func (h *Handler) get(c *gin.Context) {
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

	cart, err := h.service.GetCartWithProductByUserId(*userIdPtr)
	if err != nil {
	}

	c.JSON(http.StatusCreated, gin.H{"data": cart})
}
