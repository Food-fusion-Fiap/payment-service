package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type OrderInterface interface {
	GetOrder(orderId string) (entities.Order, error)
}
