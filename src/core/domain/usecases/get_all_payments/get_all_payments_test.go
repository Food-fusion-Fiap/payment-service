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

	t.Run("create qr code sucess case", func(t *testing.T) {
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
	prepare := func(t *testing.T, pr *mocks.PaymentRepository) {
		t.Helper()
		pr.On("FindPaymentsQuantity").Return("", errors.New("test error"))
	}

	t.Run("fails", func(t *testing.T) {
		paymentRepositoryMock := &mocks.PaymentRepository{}

		prepare(t, paymentRepositoryMock)

		usecase := GetAllPaymentsUseCase{
			PaymentRepository: paymentRepositoryMock,
		}

		output, _ := usecase.ExecuteGetAllPayments()

		assert.Equal(t, "", output)
	})
}
