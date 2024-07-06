package routes

import (
	paymentController "github.com/CAVAh/api-tech-challenge/src/adapters/controllers/payment"
	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/payments")
	{
		orderRoutes.GET("/request-payment", paymentController.RequestPayment)
	}
}
