package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/create_qr_code"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/get_all_payments"
)

func TestRequestQrCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	paymentInterfaceMock := &mocks.PaymentInterface{}
	paymentRepositoryMock := &mocks.PaymentRepository{}
	orderInterfaceMock := &mocks.OrderInterface{}
	paymentInterfaceMock.On("GetOrder", mockOrderId).Return(mockOrder, nil)
	paymentInterfaceMock.On("CreatePayment", mockOrder).Return(mockQrCode, nil)
	paymentRepositoryMock.On("Create", mockPayment).Return(mock.Anything, nil)
	useCase := create_qr_code.CreateQrCodeUseCase{
		PaymentInterface:  paymentInterfaceMock,
		PaymentRepository: paymentRepositoryMock,
		OrderInterface:    orderInterfaceMock,
	}

	paymentRepositoryMock.On("FindPaymentsQuantity").Return(uint(3), nil)

	r := gin.Default()
	r.GET("/payments/qr-code", func(c *gin.Context) {
		RequestQrCode(c, useCase)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/qr-code", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
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
