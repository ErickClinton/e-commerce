package user

import (
	"eccomerce/internal/v1/user"
	"eccomerce/internal/v1/user/dto"
	"eccomerce/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
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
	utils.Logger.Info().Msg("Start method UpdateById")

	id, _ := strconv.Atoi(c.Param("id"))
	var input dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method UpdateById %s", err.Error())
		return
	}

	if err := h.service.UpdateById(&input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (h *Handler) Delete(c *gin.Context) {
	utils.Logger.Info().Msg("Start method Delete")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.Logger.Error().Msgf("Error method Delete %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}
