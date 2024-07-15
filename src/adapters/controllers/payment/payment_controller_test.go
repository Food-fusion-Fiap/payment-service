package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

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
	gin.SetMode(gin.TestMode)

	paymentRepositoryMock := &mocks.PaymentRepository{}

	useCase := usecases.GetAllPaymentsUseCase{
		PaymentRepository: paymentRepositoryMock,
	}

	paymentRepositoryMock.On("FindPaymentsQuantity").Return(uint(3), nil)

	r := gin.Default()
	r.GET("/payments/quantity", func(c *gin.Context) {
		GetPaymentsQuantity(c, useCase)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/quantity", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
