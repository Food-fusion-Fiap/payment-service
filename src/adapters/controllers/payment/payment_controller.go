package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/infra/external/mercado_pago"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestPayment(c *gin.Context) {
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
