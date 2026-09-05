package handlers

import (
	"golang_load_test/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParticipant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid participant id"})
		return
	}

	participant, err := models.GetParticipantByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "participant not found"})
		return
	}

	payments, err := models.GetPaymentsByParticipantID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"participant": participant,
			"payments":    payments,
		},
	})
}
