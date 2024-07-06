package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type PaymentRepository interface {
	Create(order *entities.Payment) (*entities.Payment, error)
}
