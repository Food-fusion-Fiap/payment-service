package controllers

import (
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/mercado_pago"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/order_service_mock"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RequestQrCode(c *gin.Context) {
	value, _ := c.GetQuery("orderId")
	orderId, _ := strconv.Atoi(value)

	usecase := usecases.CreateQrCodeUseCase{
		PaymentInterface:  &mercado_pago.MercadoPagoIntegration{},
		PaymentRepository: &repositories.PaymentRepository{},
		OrderInterface:    &order_service_mock.OrderInterface{},
	}

	response, err := usecase.Execute(uint(orderId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func Pay(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	usecase := usecases.MakePaymentUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	response, err := usecase.ExecuteWithOrderId(uint(id))
	//TODO: avisar pro microserviço de order que foi pago

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func PayQrCode(c *gin.Context) {
	qrCode := c.Params.ByName("qr")

	usecase := usecases.MakePaymentUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	response, err := usecase.ExecuteWithQrCode(qrCode)
	//TODO: avisar pro microserviço de order que foi pago

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func CheckOrderPaymentStatus(c *gin.Context) {
	value, _ := c.GetQuery("orderId")
	orderId, _ := strconv.Atoi(value)

	usecase := usecases.CheckPaymentStatusUsecase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	response, err := usecase.Execute(uint(orderId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func MercadoPagoPayment(c *gin.Context) {
	var inputDto mercado_pago.PostPayment

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

	usecase := usecases.MakePaymentUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	if inputDto.State == mercado_pago.Finished {
		response, err := usecase.ExecuteWithOrderId(uint(orderId))
		//TODO: avisar pro microserviço de order que foi pago
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	} else if inputDto.State == mercado_pago.Error || inputDto.State == mercado_pago.Canceled {
		//TODO: avisar pro microserviço de order que foi cancelado
		c.Status(http.StatusNoContent)
	}
}

func GetPaymentsQuantity(c *gin.Context) {
	usecase := usecases.GetAllPaymentsUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	response, err := usecase.Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
