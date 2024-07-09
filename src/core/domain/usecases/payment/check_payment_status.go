package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type CheckPaymentStatusUsecase struct {
	PaymentRepository gateways.PaymentRepository
}

func (r *CheckPaymentStatusUsecase) Execute(orderId uint) (string, error) {
	payment, err := r.PaymentRepository.FindByOrderId(orderId)
	if payment.ID == 0 || err != nil {
		return "Pagamento não encontrado", errors.New("pagamento não encontrado")
	}

	if payment.PaymentStatus == enums.Paid {
		return "Pedido pago", nil
	} else if payment.PaymentStatus == enums.AwaitingPayment {
		return "Pedido aguardando pagamento", nil
	} else {
		return "Status desconhecido", nil
	}
}
