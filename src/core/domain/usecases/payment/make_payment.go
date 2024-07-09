package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type MakePaymentUseCase struct {
	PaymentRepository gateways.PaymentRepository
}

func (r *MakePaymentUseCase) Execute(orderId uint) (string, error) {
	var err error

	payment, err := r.PaymentRepository.FindByOrderId(orderId)
	if err != nil {
		return "", errors.New("Pagamento não encontrado")
	}

	if payment.PaymentStatus == enums.AwaitingPayment {
		payment.PaymentStatus = enums.Paid
		_, err = r.PaymentRepository.Create(payment)
		if err != nil {
			return "", errors.New("não foi possível efeatuar o pagamento")
		}
	} else {
		return "", errors.New("não foi possível efeatuar o pagamento")
	}

	return "Pago", nil
}
