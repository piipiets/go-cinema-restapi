package handler

import (
	"net/http"
	"strconv"

	"github.com/piipiets/go-cinema-restapi/model"
	"github.com/piipiets/go-cinema-restapi/service"

	"database/sql"
	"errors"

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

func (h *CinemaHandler) GetAllCinema(c *gin.Context) {
	cinemas, err := h.service.GetAllCinema()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get cinema data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cinemas)
}

func (h *CinemaHandler) GetCinemaById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid cinema ID",
		})
		return
	}

	cinema, err := h.service.GetCinemaById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Cinema not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get cinema",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cinema)
}

func (h *CinemaHandler) UpdateCinema(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid cinema ID",
		})
		return
	}

	var cinema model.Cinema

	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	updatedCinema, err := h.service.UpdateCinema(id, cinema)

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cinema not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update cinema",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedCinema)
}

func (h *CinemaHandler) DeleteCinema(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid cinema ID",
		})
		return
	}

	err = h.service.DeleteCinema(id)

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cinema not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete cinema",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cinema successfully deleted",
	})
}
