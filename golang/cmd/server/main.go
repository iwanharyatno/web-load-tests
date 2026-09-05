package main

import (
	"golang_load_test/internal/config"
	"golang_load_test/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/tickets", handlers.ListTickets)
		api.POST("/register", handlers.Register)
		api.GET("/participants/:id", handlers.GetParticipant)
		api.GET("/payments/:orderId", handlers.GetPayment)
		api.POST("/webhook/payment", handlers.PaymentWebhook)
	}

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
