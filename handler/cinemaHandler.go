package handler

import (
	"net/http"

	"github.com/piipiets/go-cinema-restapi/model"
	"github.com/piipiets/go-cinema-restapi/service"

	"github.com/gin-gonic/gin"
)

type CinemaHandler struct {
	service *service.CinemaService
}

func NewCinemaHandler(service *service.CinemaService) *CinemaHandler {
	return &CinemaHandler{
		service: service,
	}
}

func (h *CinemaHandler) CreateCinema(c *gin.Context) {

	var cinema model.Cinema

	// Parse JSON
	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	cinema, err := h.service.CreateCinema(cinema)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, cinema)
}
