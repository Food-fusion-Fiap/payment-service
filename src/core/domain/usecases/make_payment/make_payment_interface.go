package make_payment

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type MakePaymentInterface interface {
	ExecuteWithQrCode(qrCode string) (string, error)
	ExecuteWithOrderId(orderId uint) (string, error)
	UpdateStatus(payment entities.Payment) (string, error)
}
