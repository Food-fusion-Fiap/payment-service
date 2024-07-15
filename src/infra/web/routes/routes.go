package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	controller "github.com/CAVAh/api-tech-challenge/src/adapters/controllers/payment"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
)

func HandleRequests() {
	router := gin.Default()

	getAllPaymentsUseCase := usecases.GetAllPaymentsUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	orderRoutes := router.Group("/payments")
	{
		orderRoutes.GET("/qr-code", controller.RequestQrCode)
		orderRoutes.GET("/status", controller.CheckOrderPaymentStatus)
		orderRoutes.POST("", controller.MercadoPagoPayment)
		orderRoutes.GET("/quantity", func(c *gin.Context) {
			controller.GetPaymentsQuantity(c, getAllPaymentsUseCase)
		})

		//Routes that mock MercadoPago webhook payment:
		orderRoutes.POST("/alternative-pay/:id", controller.Pay)
		orderRoutes.POST("/alternative-pay-with-qr-code/:qr", controller.PayQrCode)
	}

	err := router.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
