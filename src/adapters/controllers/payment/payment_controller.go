package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/check_payment_status"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/create_qr_code"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/get_all_payments"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/make_payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/mercado_pago"
)

type PaymentController struct {
	PaymentInterface  gateways.PaymentInterface
	PaymentRepository gateways.PaymentRepository
	OrderInterface    gateways.OrderInterface
}

func RequestQrCode(c *gin.Context, useCase create_qr_code.CreateQrCodeInterface) {
	value, _ := c.GetQuery("orderId")
	orderId, _ := strconv.Atoi(value)

	response, err := useCase.ExecuteCreateQrCode(uint(orderId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func Pay(c *gin.Context, useCase make_payment.MakePaymentInterface) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	response, err := useCase.ExecuteApprovedPaymentWithOrderId(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func PayQrCode(c *gin.Context, useCase make_payment.MakePaymentInterface) {
	qrCode := c.Params.ByName("qr")

	response, err := useCase.ExecuteApprovedPaymentWithQrCode(qrCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func CheckOrderPaymentStatus(c *gin.Context, useCase check_payment_status.CheckPaymentStatusUseCaseInterface) {
	value, _ := c.GetQuery("orderId")
	orderId, _ := strconv.Atoi(value)

	response, err := useCase.ExecuteCheckPaymentStatus(uint(orderId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func MercadoPagoPayment(c *gin.Context, useCase make_payment.MakePaymentInterface) {
	var inputDto mercado_pago.PostPayment
	var err error
	var response string

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//value, _ := c.GetQuery("data.id")
	//orderId, _ := strconv.Atoi(value)
	//Explicação: para funcionar o teste do mercado livre, precisa pegar do ID,
	//já que o external reference não é mandado. Mas o id de dentro da aplicação estará em external reference
	var orderId, _ = strconv.Atoi(inputDto.AdditionalInfo.ExternalReference)

	if inputDto.State == mercado_pago.Finished {
		response, err = useCase.ExecuteApprovedPaymentWithOrderId(uint(orderId))
	} else if inputDto.State == mercado_pago.Error || inputDto.State == mercado_pago.Canceled {
		response, err = useCase.ExecuteErrorPaymentWithOrderId(uint(orderId))
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetPaymentsQuantity(c *gin.Context, useCase get_all_payments.GetAllPaymentsInterface) {
	response, err := useCase.ExecuteGetAllPayments()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
