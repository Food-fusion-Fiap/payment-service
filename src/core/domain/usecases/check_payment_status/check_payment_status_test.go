package check_payment_status

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

func TestCheckPaymentStatus_PaidStatus(t *testing.T) {
	mockOrderId := "randomuuid"
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: "randomuuid2"}

	paymentRepositoryMock := &gateways.MockPaymentRepository{}

	usecase := CheckPaymentStatusUsecase{
		PaymentRepository: paymentRepositoryMock,
	}

	t.Run("status paid", func(t *testing.T) {
		prepare := func(t *testing.T, pr *gateways.MockPaymentRepository) {
			t.Helper()
			pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		}
		prepare(t, paymentRepositoryMock)

		output, err := usecase.ExecuteCheckPaymentStatus(mockOrderId)

		assert.Equal(t, "Pedido pago", output)
		assert.Equal(t, nil, err)
	})
}

func TestCheckPaymentStatus_StatusAwaitingPayment(t *testing.T) {
	mockOrderId := "randomuuid"
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: "randomuuid2"}

	paymentRepositoryMock := &gateways.MockPaymentRepository{}

	usecase := CheckPaymentStatusUsecase{
		PaymentRepository: paymentRepositoryMock,
	}

	t.Run("awaiting payment", func(t *testing.T) {
		mockPayment.PaymentStatus = enums.AwaitingPayment
		prepare := func(t *testing.T, pr *gateways.MockPaymentRepository) {
			t.Helper()
			pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		}
		prepare(t, paymentRepositoryMock)

		output, err := usecase.ExecuteCheckPaymentStatus(mockOrderId)

		assert.Equal(t, "Pedido aguardando pagamento", output)
		assert.Equal(t, nil, err)
	})
}

func TestCheckPaymentStatus_UnknownStatus(t *testing.T) {
	mockOrderId := "randomuuid"
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: "randomuuid2"}

	paymentRepositoryMock := &gateways.MockPaymentRepository{}

	usecase := CheckPaymentStatusUsecase{
		PaymentRepository: paymentRepositoryMock,
	}

	t.Run("unknown status", func(t *testing.T) {
		mockPayment.PaymentStatus = "errado"
		prepare := func(t *testing.T, pr *gateways.MockPaymentRepository) {
			t.Helper()
			pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		}
		prepare(t, paymentRepositoryMock)

		output, err := usecase.ExecuteCheckPaymentStatus(mockOrderId)

		assert.Equal(t, "Status desconhecido", output)
		assert.Equal(t, nil, err)
	})
}

func TestCheckPaymentStatus_DoNotFindOrder(t *testing.T) {
	mockOrderId := "randomuuid"
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: "randomuuid2"}

	paymentRepositoryMock := &gateways.MockPaymentRepository{}

	usecase := CheckPaymentStatusUsecase{
		PaymentRepository: paymentRepositoryMock,
	}

	t.Run("order not found", func(t *testing.T) {
		mockPayment.PaymentStatus = "errado"
		prepare := func(t *testing.T, pr *gateways.MockPaymentRepository) {
			t.Helper()
			pr.On("FindByOrderId", mockOrderId).Return(mockPayment, errors.New("do not found"))
		}
		prepare(t, paymentRepositoryMock)

		output, _ := usecase.ExecuteCheckPaymentStatus(mockOrderId)

		assert.Equal(t, "", output)
	})
}
