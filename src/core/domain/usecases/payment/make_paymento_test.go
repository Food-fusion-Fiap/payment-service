package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMakePaymentUseCase(t *testing.T) {
	mockQrCode := "mockQrCode"
	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindByQrCode", mockQrCode).Return(mock.Anything)
	}

	t.Run("when sucess findbyQRCode", func(t *testing.T) {
		paymentRepositoryMock := mocks.PaymentRepository{}
		prepare(t, &paymentRepositoryMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: &paymentRepositoryMock,
		}

		_, err := usecase.ExecuteWithQrCode(mockQrCode)
		if err != nil {
			return
		}
	})
}
