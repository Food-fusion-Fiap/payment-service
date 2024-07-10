package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"testing"
)

func TestMakePaymentUseCase_FindByQrCode_Success(t *testing.T) {
	mockQrCode := "mockQrCode"
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindByQrCode", mockQrCode).Return(mockPayment, nil)
	}

	t.Run("when success findbyQRCode", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}
		prepare(t, paymentRepositoryMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		_, err := usecase.ExecuteWithQrCode(mockQrCode)
		if err != nil {
			return
		}

	})
}

func TestMakePaymentUseCase(t *testing.T) {
	mockOrderId := uint(407)
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
	}

	t.Run("when sucess findbyQRCode", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}
		prepare(t, paymentRepositoryMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		_, err := usecase.ExecuteWithOrderId(mockOrderId)
		if err != nil {
			return
		}
	})
}
