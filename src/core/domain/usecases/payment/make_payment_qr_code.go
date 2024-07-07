package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type MakePaymentWithQrCodeUseCase struct {
	PaymentRepository gateways.PaymentRepository
}

func (r *MakePaymentWithQrCodeUseCase) Execute(qrCode string) (string, error) {
	var err error

	payment, err := r.PaymentRepository.FindByQrCode(qrCode)
	if err != nil {
		return "", err
	}

	if payment.PaymentStatus == enums.AwaitingPayment {
		payment.PaymentStatus = enums.Paid
		_, err = r.PaymentRepository.Create(payment)
		if err != nil {
			return "", errors.New("Não foi possível efeatuar o pagamento")
		}
	} else {
		return "", errors.New("Não foi possível efeatuar o pagamento")
	}

	return "Pago", nil
}
