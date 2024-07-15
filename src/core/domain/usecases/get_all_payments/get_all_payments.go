package get_all_payments

import (
	"strconv"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
)

type GetAllPaymentsUseCase struct {
	PaymentRepository gateways.PaymentRepository
}

func (r *GetAllPaymentsUseCase) ExecuteGetAllPayments() (string, error) {
	var err error
	quantity, err := r.PaymentRepository.FindPaymentsQuantity()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(quantity)), nil
}
