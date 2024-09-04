package routes

import (
	"github.com/CAVAh/api-tech-challenge/src/infra/external/aws_sns"
	"log"

	"github.com/gin-gonic/gin"

	controller "github.com/CAVAh/api-tech-challenge/src/adapters/controllers/payment"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/check_payment_status"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/create_qr_code"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/get_all_payments"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/make_payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/mercado_pago"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/order_service"
)

func HandleRequests() {
	router := gin.Default()

	getAllPaymentsUseCase := get_all_payments.GetAllPaymentsUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	createQrCodeUseCase := create_qr_code.CreateQrCodeUseCase{
		PaymentInterface:  &mercado_pago.MercadoPagoIntegration{},
		PaymentRepository: &repositories.PaymentRepository{},
		OrderInterface:    &order_service.OrderInterface{},
	}

	makePaymentUseCase := usecases.MakePaymentUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
		OrderInterface:    &order_service.OrderInterface{},
		PubSubInterface:   &aws_sns.PubSubInterface{},
	}

	checkPaymentStatusUsecase := check_payment_status.CheckPaymentStatusUsecase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	paymentRoutes := router.Group("/payments")
	{
		paymentRoutes.GET("/qr-code", func(c *gin.Context) {
			controller.RequestQrCode(c, &createQrCodeUseCase)
		})
		paymentRoutes.GET("/status", func(c *gin.Context) {
			controller.CheckOrderPaymentStatus(c, &checkPaymentStatusUsecase)
		})
		paymentRoutes.POST("", func(c *gin.Context) {
			controller.MercadoPagoPayment(c, &makePaymentUseCase)
		})
		paymentRoutes.GET("/quantity", func(c *gin.Context) {
			controller.GetPaymentsQuantity(c, &getAllPaymentsUseCase)
		})

		//Routes that mock MercadoPago webhook payment:
		paymentRoutes.POST("/alternative-pay-success/:id", func(c *gin.Context) {
			controller.PayAlternativeSuccess(c, &makePaymentUseCase)
		})
		paymentRoutes.POST("/alternative-pay-failure/:id", func(c *gin.Context) {
			controller.PayAlternativeFailure(c, &makePaymentUseCase)
		})
	}

	err := router.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
