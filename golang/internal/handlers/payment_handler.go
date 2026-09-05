package handlers

import (
	"golang_load_test/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	orderID := c.Param("orderId")

	payment, err := models.GetPaymentByOrderID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": payment})
}
