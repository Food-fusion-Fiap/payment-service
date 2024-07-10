package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakePaymentUseCase_FindByQrCodeAndUpdatePayment_Success(t *testing.T) {
	mockQrCode := "mockQrCode"
	mockPayment := entities.Payment{PaymentStatus: enums.AwaitingPayment, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindByQrCode", mockQrCode).Return(mockPayment, nil)
		pr.On("UpdateToPaid", mockPayment.ID).Return()

	}

	t.Run("when success findbyQRCode", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}
		prepare(t, paymentRepositoryMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		output, err := usecase.ExecuteWithQrCode(mockQrCode)

		assert.Equal(t, "Pago", output)
		assert.Equal(t, nil, err)
	})
}

func TestMakePaymentUseCase_FindByQrCodeAndUpdatePayment_Fail(t *testing.T) {
	mockQrCode := "mockQrCode"
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindByQrCode", mockQrCode).Return(mockPayment, nil)
		pr.On("UpdateToPaid", mockPayment.ID).Return()

	}

	t.Run("when fails findbyQRCode", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}
		prepare(t, paymentRepositoryMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		output, err := usecase.ExecuteWithQrCode(mockQrCode)

		assert.Equal(t, output, "")
		assert.Equal(t, errors.New("não foi possível efetuar o pagamento: o pagamento já foi pago"), err)
	})
}

func TestMakePaymentUseCase_FindByOrderIdAndUpdatePayment_Success(t *testing.T) {
	mockOrderId := uint(407)
	mockPayment := entities.Payment{PaymentStatus: enums.AwaitingPayment, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		pr.On("UpdateToPaid", mockPayment.ID).Return()

	}

	t.Run("when success FindByOrderId", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}
		prepare(t, paymentRepositoryMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		output, err := usecase.ExecuteWithOrderId(mockOrderId)

		assert.Equal(t, "Pago", output)
		assert.Equal(t, nil, err)
	})
}

func TestMakePaymentUseCase_FindByOrderIdAndUpdatePayment_Fail(t *testing.T) {
	mockOrderId := uint(407)
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		pr.On("UpdateToPaid", mockPayment.ID).Return()

	}

	t.Run("when fails findbyQROrderId", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}
		prepare(t, paymentRepositoryMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		output, err := usecase.ExecuteWithOrderId(mockOrderId)

		assert.Equal(t, output, "")
		assert.Equal(t, errors.New("não foi possível efetuar o pagamento: o pagamento já foi pago"), err)
	})
}
