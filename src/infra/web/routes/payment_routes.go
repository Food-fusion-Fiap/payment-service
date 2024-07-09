package routes

import (
	paymentController "github.com/CAVAh/api-tech-challenge/src/adapters/controllers/payment"
	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/payments")
	{
		orderRoutes.GET("/qr-code", paymentController.RequestQrCode)
		orderRoutes.GET("/status", paymentController.CheckOrderPaymentStatus)
		orderRoutes.POST("", paymentController.MercadoPagoPayment)

		//Routes that mock MercadoPago webhook payment:
		orderRoutes.POST("/alternative-pay/:id", paymentController.Pay)
		orderRoutes.POST("/alternative-pay-with-qr-code/:qr", paymentController.PayQrCode)
	}
}
