package check_payment_status

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type CheckPaymentStatusUsecase struct {
	PaymentRepository gateways.PaymentRepository
}

func (r *CheckPaymentStatusUsecase) ExecuteCheckPaymentStatus(orderId uint) (string, error) {
	payment, err := r.PaymentRepository.FindByOrderId(orderId)
	if err != nil {
		return "", err
	}

	if payment.PaymentStatus == enums.Paid {
		return "Pedido pago", nil
	} else if payment.PaymentStatus == enums.AwaitingPayment {
		return "Pedido aguardando pagamento", nil
	} else {
		return "Status desconhecido", nil
	}
}
