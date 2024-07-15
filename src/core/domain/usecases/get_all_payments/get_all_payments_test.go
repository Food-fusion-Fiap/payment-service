package get_all_payments

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways/mocks"
)

func TestGetAllPaymentsUseCase(t *testing.T) {
	mockedQuantity := uint(2)

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindPaymentsQuantity").Return(mockedQuantity, nil)
	}

	t.Run("get all payments", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}

		prepare(t, paymentRepositoryMock)

		usecase := GetAllPaymentsUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		output, err := usecase.ExecuteGetAllPayments()

		assert.Equal(t, strconv.Itoa(int(mockedQuantity)), output)
		assert.Equal(t, nil, err)
	})
}

func TestGetAllPaymentsUseCase_Fails(t *testing.T) {
	mockedQuantity := uint(2)

	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindPaymentsQuantity").Return(mockedQuantity, errors.New("fail"))
	}

	t.Run("get all payments - fails", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}

		prepare(t, paymentRepositoryMock)

		usecase := GetAllPaymentsUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		output, _ := usecase.ExecuteGetAllPayments()

		assert.Equal(t, "", output)
	})
}
