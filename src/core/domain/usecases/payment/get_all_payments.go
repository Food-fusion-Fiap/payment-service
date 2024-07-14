package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
)

type GetAllPaymentsUseCase struct {
	PaymentRepository gateways.PaymentRepository
}

func (r *GetAllPaymentsUseCase) Execute() (uint, error) {
	var err error
	quantity, err := r.PaymentRepository.FindPaymentsQuantity()
	if err != nil {
		return 0, err
	}
	return quantity, nil
}
