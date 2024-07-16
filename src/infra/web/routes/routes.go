package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	controller "github.com/CAVAh/api-tech-challenge/src/adapters/controllers/payment"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/check_payment_status"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/create_qr_code"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/get_all_payments"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/make_payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/mercado_pago"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/order_service_mock"
)

func HandleRequests() {
	router := gin.Default()

	getAllPaymentsUseCase := get_all_payments.GetAllPaymentsUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	createQrCodeUseCase := create_qr_code.CreateQrCodeUseCase{
		PaymentInterface:  &mercado_pago.MercadoPagoIntegration{},
		PaymentRepository: &repositories.PaymentRepository{},
		OrderInterface:    &order_service_mock.OrderInterface{},
	}

	makePaymentUseCase := usecases.MakePaymentUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
		OrderInterface:    &order_service_mock.OrderInterface{},
	}

	checkPaymentStatusUsecase := check_payment_status.CheckPaymentStatusUsecase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	orderRoutes := router.Group("/payments")
	{
		orderRoutes.GET("/qr-code", func(c *gin.Context) {
			controller.RequestQrCode(c, &createQrCodeUseCase)
		})
		orderRoutes.GET("/status", func(c *gin.Context) {
			controller.CheckOrderPaymentStatus(c, &checkPaymentStatusUsecase)
		})
		orderRoutes.POST("", func(c *gin.Context) {
			controller.MercadoPagoPayment(c, &makePaymentUseCase)
		})
		orderRoutes.GET("/quantity", func(c *gin.Context) {
			controller.GetPaymentsQuantity(c, &getAllPaymentsUseCase)
		})

		//Routes that mock MercadoPago webhook payment:
		orderRoutes.POST("/alternative-pay/:id", func(c *gin.Context) {
			controller.Pay(c, &makePaymentUseCase)
		})
		orderRoutes.POST("/alternative-pay-with-qr-code/:qr", func(c *gin.Context) {
			controller.PayQrCode(c, &makePaymentUseCase)
		})
	}

	err := router.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
