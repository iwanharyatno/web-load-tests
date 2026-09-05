package handlers

import (
	"golang_load_test/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTickets(c *gin.Context) {
	tickets, err := models.GetAllTickets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tickets})
}
