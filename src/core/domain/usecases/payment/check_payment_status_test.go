package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPaymentStatus_PaidStatus(t *testing.T) {
	mockOrderId := uint(407)
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	paymentRepositoryMock := &mocks.PaymentRepository{}

	usecase := CheckPaymentStatusUsecase{
		PaymentRepository: paymentRepositoryMock,
	}

	t.Run("status paid", func(t *testing.T) {
		prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
			t.Helper()
			pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		}
		prepare(t, paymentRepositoryMock)

		output, err := usecase.Execute(mockOrderId)

		assert.Equal(t, "Pedido pago", output)
		assert.Equal(t, nil, err)
	})
}

func TestCheckPaymentStatus_StatusAwaitingPayment(t *testing.T) {
	mockOrderId := uint(407)
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	paymentRepositoryMock := &mocks.PaymentRepository{}

	usecase := CheckPaymentStatusUsecase{
		PaymentRepository: paymentRepositoryMock,
	}

	t.Run("awaiting payment", func(t *testing.T) {
		mockPayment.PaymentStatus = enums.AwaitingPayment
		prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
			t.Helper()
			pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		}
		prepare(t, paymentRepositoryMock)

		output, err := usecase.Execute(mockOrderId)

		assert.Equal(t, "Pedido aguardando pagamento", output)
		assert.Equal(t, nil, err)
	})
}

func TestCheckPaymentStatus_UnknownStatus(t *testing.T) {
	mockOrderId := uint(407)
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	paymentRepositoryMock := &mocks.PaymentRepository{}

	usecase := CheckPaymentStatusUsecase{
		PaymentRepository: paymentRepositoryMock,
	}

	t.Run("unknown status", func(t *testing.T) {
		mockPayment.PaymentStatus = "errado"
		prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
			t.Helper()
			pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		}
		prepare(t, paymentRepositoryMock)

		output, err := usecase.Execute(mockOrderId)

		assert.Equal(t, "Status desconhecido", output)
		assert.Equal(t, nil, err)
	})
}
