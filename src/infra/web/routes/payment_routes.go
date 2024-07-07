package routes

import (
	paymentController "github.com/CAVAh/api-tech-challenge/src/adapters/controllers/payment"
	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/payments")
	{
		orderRoutes.GET("/request-qr-code", paymentController.RequestQrCode)
		orderRoutes.PATCH("/alternative-pay/:id", paymentController.Pay)
		orderRoutes.PATCH("/alternative-pay-with-qr-code/:qr", paymentController.PayQrCode)
	}
}
