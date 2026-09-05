package handlers

import (
	"fmt"
	"golang_load_test/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	TicketID  int64  `json:"ticket_id" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	ticket, err := models.GetTicketByID(req.TicketID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid ticket id"})
		return
	}

	participant, err := models.CreateParticipant(req.Name, req.Email, req.Phone)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "email already exists or database error"})
		return
	}

	orderID := fmt.Sprintf("REG-%d-%d", participant.ID, time.Now().Unix())

	payment, err := models.CreatePayment(participant.ID, ticket.ID, orderID, ticket.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"participant": participant,
			"payment":     payment,
		},
	})
}
