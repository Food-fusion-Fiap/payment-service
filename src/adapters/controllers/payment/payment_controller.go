package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/mercado_pago"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RequestQrCode(c *gin.Context) {
	var inputDto entities.Order

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usecase := usecases.CreateQrCodeUseCase{
		PaymentInterface:  &mercado_pago.MercadoPagoIntegration{},
		PaymentRepository: &repositories.PaymentRepository{},
	}

	response, err := usecase.Execute(inputDto)

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

	response, err := usecase.Execute(uint(id))

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

	usecase := usecases.MakePaymentWithQrCodeUseCase{
		PaymentRepository: &repositories.PaymentRepository{},
	}

	response, err := usecase.Execute(qrCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

//func ChangeOrderStatus(c *gin.Context) {
//	var inputDto dtos.ChangeOrderStatusDto
//	id, _ := strconv.Atoi(c.Params.ByName("id"))
//
//	if err := c.ShouldBindJSON(&inputDto); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//
//	if err := validator.Validate(inputDto); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err,
//		})
//		return
//	}
//
//	orderRepository := &repositories.OrderRepository{}
//
//	usecase := order.ChangeOrderStatusUsecase{
//		OrderRepository: orderRepository,
//	}
//
//	orderResult, err := usecase.Execute(uint(id), inputDto.ChangeToStatus)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"orderId":   orderResult.ID,
//		"status":    orderResult.Status,
//		"updatedAt": orderResult.UpdatedAt,
//	})
//}
