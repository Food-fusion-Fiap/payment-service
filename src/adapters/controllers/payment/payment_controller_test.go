package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/payment"
)

func TestRequestQrCode(t *testing.T) {
	//Mock gin context
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	r := gin.Default()
	r.GET("/payments/qr-code", func(c *gin.Context) {
		RequestQrCode(c)
	})
}

func TestGetPaymentsQuantity(t *testing.T) {
	// Configurar o gin em modo de teste
	gin.SetMode(gin.TestMode)

	// Criar o mock do repositório de clientes
	// Substituir o repositório real pelo mock no usecase
	useCase := usecases.GetAllPaymentsUseCase{
		PaymentRepository: &mocks.PaymentRepository{},
	}

	// Configurar o controlador com o mock do usecase
	r := gin.Default()
	r.GET("/quantity", func(c *gin.Context) {
		GetPaymentsQuantity(c, useCase)
	})
}
