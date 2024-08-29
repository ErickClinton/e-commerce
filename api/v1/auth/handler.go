package auth

import (
	"eccomerce/internal/v1/auth"
	"eccomerce/internal/v1/auth/dto"
	"eccomerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerAuth struct {
	services auth.ServiceAuth
}

func NewHandlerAuth(service auth.ServiceAuth) *HandlerAuth {
	return &HandlerAuth{services: service}
}

func (h *HandlerAuth) Login(c *gin.Context) {
	utils.Logger.Info().Msg("Start method login")
	var input dto.LoginUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method create %s", err.Error())
		return
	}

	token, err := h.services.Login(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
