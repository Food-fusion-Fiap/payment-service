package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type PaymentRepository interface {
	Create(order entities.Payment) (string, error)
	FindByOrderId(orderId uint) (entities.Payment, error)
	FindByQrCode(qrCode string) (entities.Payment, error)
	UpdateToPaid(paymentID uint)
}
