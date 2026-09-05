package handlers

import (
	"golang_load_test/internal/models"
	"golang_load_test/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookRequest struct {
	OrderID string `json:"order_id" binding:"required"`
}

func PaymentWebhook(c *gin.Context) {
	var req WebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	payment, err := models.GetPaymentByOrderID(req.OrderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "payment not found"})
		return
	}

	if payment.Status == "paid" {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "payment already processed",
			"data":    payment,
		})
		return
	}

	ticket, err := models.GetTicketByID(payment.TicketID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "ticket not found"})
		return
	}

	bibNumber := services.GenerateBibNumber(ticket.BibPrefix, ticket.BibPadding, ticket.BibIncrement)

	if err := models.IncrementBib(ticket.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	if err := models.MarkPaymentPaid(payment.ID, bibNumber); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	payment.Status = "paid"
	payment.BibNumber = &bibNumber

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "payment confirmed",
		"data":    payment,
	})
}
