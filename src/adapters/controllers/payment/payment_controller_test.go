package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/check_payment_status"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/create_qr_code"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/get_all_payments"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/make_payment"
)

func TestRequestQrCode_Success(t *testing.T) {
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

func TestRequestQrCode_Fails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &create_qr_code.CreateQrCodeInterfaceMock{}
	useCaseMock.On("ExecuteCreateQrCode", mock.Anything).Return("", errors.New("some error"))

	r := gin.Default()
	r.GET("/payments/qr-code", func(c *gin.Context) {
		RequestQrCode(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/qr-code", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPay_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &make_payment.MakePaymentInterfaceMock{}
	useCaseMock.On("ExecuteWithOrderId", mock.Anything).Return("success", nil)

	r := gin.Default()
	r.POST("/payments/alternative-pay", func(c *gin.Context) {
		Pay(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodPost, "/payments/alternative-pay", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPay_Fails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &make_payment.MakePaymentInterfaceMock{}
	useCaseMock.On("ExecuteWithOrderId", mock.Anything).Return("", errors.New("some error"))

	r := gin.Default()
	r.POST("/payments/alternative-pay", func(c *gin.Context) {
		Pay(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodPost, "/payments/alternative-pay", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPayQrCode_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &make_payment.MakePaymentInterfaceMock{}
	useCaseMock.On("ExecuteWithQrCode", mock.Anything).Return("success", nil)

	r := gin.Default()
	r.POST("/payments/alternative-pay-with-qr-code", func(c *gin.Context) {
		PayQrCode(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodPost, "/payments/alternative-pay-with-qr-code", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPayQrCode_Fails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &make_payment.MakePaymentInterfaceMock{}
	useCaseMock.On("ExecuteWithQrCode", mock.Anything).Return("", errors.New("some error"))

	r := gin.Default()
	r.POST("/payments/alternative-pay-with-qr-code", func(c *gin.Context) {
		PayQrCode(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodPost, "/payments/alternative-pay-with-qr-code", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCheckOrderPaymentStatus_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &check_payment_status.CheckPaymentStatusUseCaseInterfaceMock{}
	useCaseMock.On("ExecuteCheckPaymentStatus", mock.Anything).Return("success", nil)

	r := gin.Default()
	r.GET("/payments/status", func(c *gin.Context) {
		CheckOrderPaymentStatus(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/status", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckOrderPaymentStatus_Fails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &check_payment_status.CheckPaymentStatusUseCaseInterfaceMock{}
	useCaseMock.On("ExecuteCheckPaymentStatus", mock.Anything).Return("", errors.New("some error"))

	r := gin.Default()
	r.GET("/payments/status", func(c *gin.Context) {
		CheckOrderPaymentStatus(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/status", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPaymentsQuantity_Success(t *testing.T) {
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

func TestGetPaymentsQuantity_Fails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	useCaseMock := &get_all_payments.GetAllPaymentsInterfaceMock{}
	useCaseMock.On("ExecuteGetAllPayments").Return("", errors.New("some error"))

	r := gin.Default()
	r.GET("/payments/quantity", func(c *gin.Context) {
		GetPaymentsQuantity(c, useCaseMock)
	})

	req, _ := http.NewRequest(http.MethodGet, "/payments/quantity", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
