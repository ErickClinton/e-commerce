package wallet

import (
	"eccomerce/internal/v1/wallet/dto"
	"eccomerce/internal/v1/wallet/services"
	"eccomerce/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service services.WalletService
}

func NewHandler(service services.WalletService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetByID(c *gin.Context) {
	utils.Logger.Info().Msg("Start method GetByID")
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		utils.Logger.Error().Msgf("Error method GetByID %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) UpdateById(c *gin.Context) {
	utils.Logger.Info().Msg("Start method UpdateById")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		utils.Logger.Error().Msgf("Error converting ID %s", err.Error())
		return
	}

	var input dto.CreateWalletRequest
	input.UserId = uint(id)
	utils.Logger.Info().Msgf("Input before setting UserId: %+v", input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error binding JSON %s", err.Error())
		return
	}

	// Log para verificar o valor do ID e do input
	utils.Logger.Info().Msgf("ID from URL: %d", id)
	utils.Logger.Info().Msgf("Input before setting UserId: %+v", input)

	input.UserId = uint(id)

	// Log para verificar o valor do input após setar o UserId
	utils.Logger.Info().Msgf("Input after setting UserId: %+v", input)

	if err := h.service.Update(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error updating wallet %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
