package create_qr_code

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

func TestCreateQRCodeUseCase(t *testing.T) {
	mockOrderId := uint(905)
	mockOrder := entities.Order{ID: mockOrderId}
	mockQrCode := "mockqrcode"
	mockPayment := entities.Payment{
		OrderID:       mockOrderId,
		QrCode:        mockQrCode,
		PaymentStatus: enums.AwaitingPayment,
	}

	prepare := func(t *testing.T, pi *gateways.MockPaymentInterface, pr *gateways.MockPaymentRepository, oi *gateways.MockOrderInterface) {
		t.Helper()
		oi.On("GetOrder", mockOrderId).Return(mockOrder, nil)
		pi.On("CreatePayment", mockOrder).Return(mockQrCode, nil)
		pr.On("Create", mockPayment).Return(mock.Anything, nil)
	}

	t.Run("create qr code sucess case", func(t *testing.T) {
		paymentInterfaceMock := &gateways.MockPaymentInterface{}
		paymentRepositoryMock := &gateways.MockPaymentRepository{}
		orderInterfaceMock := &gateways.MockOrderInterface{}

		prepare(t, paymentInterfaceMock, paymentRepositoryMock, orderInterfaceMock)

		usecase := CreateQrCodeUseCase{
			PaymentInterface:  paymentInterfaceMock,
			PaymentRepository: paymentRepositoryMock,
			OrderInterface:    orderInterfaceMock,
		}

		output, err := usecase.ExecuteCreateQrCode(mockOrderId)

		assert.Equal(t, mockQrCode, output)
		assert.Equal(t, nil, err)
	})
}
