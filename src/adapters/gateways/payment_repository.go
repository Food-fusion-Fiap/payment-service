package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type PaymentRepository interface {
	Create(order entities.Payment) (entities.Payment, error)
	FindByOrderId(orderId uint) (entities.Payment, error)
	FindByQrCode(qrCode string) (entities.Payment, error)
	UpdateToPaid(paymentID uint)
}
