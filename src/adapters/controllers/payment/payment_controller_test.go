package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/create_qr_code"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/get_all_payments"
)

func TestRequestQrCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &create_qr_code.CreateQrCodeInterfaceMock{}
	useCaseMock.On("ExecuteCreateQrCode", mock.Anything).Return("mockQrCode", nil)

	r := gin.Default()
	r.GET("/payments/qr-code", func(c *gin.Context) {
		RequestQrCode(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/qr-code", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPaymentsQuantity(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &get_all_payments.GetAllPaymentsInterfaceMock{}
	useCaseMock.On("ExecuteGetAllPayments").Return("mockResponse", nil)

	r := gin.Default()
	r.GET("/payments/quantity", func(c *gin.Context) {
		GetPaymentsQuantity(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/quantity", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
