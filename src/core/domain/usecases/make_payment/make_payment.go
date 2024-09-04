package make_payment

import (
	"errors"
	"log"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type MakePaymentUseCase struct {
	PaymentRepository gateways.PaymentRepository
	OrderInterface    gateways.OrderInterface
	PubSubInterface   gateways.PubSubInterface
}

func (r *MakePaymentUseCase) ExecuteApprovedPaymentWithOrderId(orderId string) (string, error) {
	var err error

	payment, err := r.PaymentRepository.FindByOrderId(orderId)
	if err != nil {
		log.Println(err)
		return "", errors.New("pagamento não encontrado")
	}

	out, err := r.UpdateToStatusApproved(payment)

	return out, err
}

func (r *MakePaymentUseCase) UpdateToStatusApproved(payment entities.Payment) (string, error) {
	if payment.PaymentStatus == enums.AwaitingPayment {
		payment.PaymentStatus = enums.Paid
		log.Print(payment)
		r.PaymentRepository.UpdateToPaid(payment.ID)
		err := r.PubSubInterface.NotifyPaymentApproved(payment.OrderID)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return "Pago", nil
	} else {
		return "", errors.New("não foi possível efetuar o pagamento: o pagamento já foi pago")
	}
}

func (r *MakePaymentUseCase) ExecuteErrorPaymentWithOrderId(orderId string) (string, error) {
	//TODO: this service can be improved and update payment status to error in DB. For now, we only notify a topic
	err := r.PubSubInterface.NotifyPaymentError(orderId)
	if err != nil {
		return "", err
	}
	return "Erro", nil
}
