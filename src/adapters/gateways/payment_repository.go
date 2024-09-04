package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type PaymentRepository interface {
	Create(order entities.Payment) (string, error)
	FindByOrderId(orderId string) (entities.Payment, error)
	UpdateToPaid(paymentID uint)
	FindPaymentsQuantity() (uint, error)
}
