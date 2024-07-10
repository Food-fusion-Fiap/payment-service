package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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

	prepare := func(t *testing.T, pi *mocks.PaymentInterface, pr *mocks.PaymentRepository, oi *mocks.OrderInterface) {
		t.Helper()
		oi.On("GetOrder", mockOrderId).Return(mockOrder, nil)
		pi.On("CreatePayment", mockOrder).Return(mockQrCode, nil)
		pr.On("Create", mockPayment).Return(mock.Anything, nil)
	}

	t.Run("create qr code sucess case", func(t *testing.T) {
		paymentInterfaceMock := &mocks.PaymentInterface{}
		paymentRepositoryMock := &mocks.PaymentRepository{}
		orderInterfaceMock := &mocks.OrderInterface{}

		prepare(t, paymentInterfaceMock, paymentRepositoryMock, orderInterfaceMock)

		usecase := CreateQrCodeUseCase{
			PaymentInterface:  paymentInterfaceMock,
			PaymentRepository: paymentRepositoryMock,
			OrderInterface:    orderInterfaceMock,
		}

		output, err := usecase.Execute(mockOrderId)

		assert.Equal(t, mockQrCode, output)
		assert.Equal(t, nil, err)
	})
}
