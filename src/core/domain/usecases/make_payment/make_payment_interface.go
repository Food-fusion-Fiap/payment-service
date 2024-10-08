package make_payment

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type MakePaymentInterface interface {
	ExecuteApprovedPaymentWithOrderId(orderId string) (string, error)
	UpdateToStatusApproved(payment entities.Payment) (string, error)
	ExecuteErrorPaymentWithOrderId(orderId string) (string, error)
}
