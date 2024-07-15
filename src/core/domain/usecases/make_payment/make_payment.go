package make_payment

import (
	"errors"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type MakePaymentUseCase struct {
	PaymentRepository gateways.PaymentRepository
}

func (r *MakePaymentUseCase) ExecuteWithQrCode(qrCode string) (string, error) {
	var err error

	payment, err := r.PaymentRepository.FindByQrCode(qrCode)
	if err != nil {
		return "", err
	}

	return r.UpdateStatus(payment)
}

func (r *MakePaymentUseCase) ExecuteWithOrderId(orderId uint) (string, error) {
	var err error

	payment, err := r.PaymentRepository.FindByOrderId(orderId)
	if err != nil {
		return "", errors.New("pagamento não encontrado")
	}

	return r.UpdateStatus(payment)
}

func (r *MakePaymentUseCase) UpdateStatus(payment entities.Payment) (string, error) {
	if payment.PaymentStatus == enums.AwaitingPayment {
		payment.PaymentStatus = enums.Paid
		r.PaymentRepository.UpdateToPaid(payment.ID)
		return "Pago", nil
	} else {
		return "", errors.New("não foi possível efetuar o pagamento: o pagamento já foi pago")
	}
}
