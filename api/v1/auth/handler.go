package auth

import (
	"eccomerce/internal/v1/auth/dto"
	"eccomerce/internal/v1/auth/services"
	"eccomerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerAuth struct {
	serviceAuth services.ServiceAuth
}

func NewHandlerAuth(service services.ServiceAuth) *HandlerAuth {
	return &HandlerAuth{serviceAuth: service}
}

func (h *HandlerAuth) Login(c *gin.Context) {
	utils.Logger.Info().Msg("Start method login")
	var input dto.LoginUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method create %s", err.Error())
		return
	}

	token, err := h.serviceAuth.Login(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
