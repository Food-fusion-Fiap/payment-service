package order_service_mock

import (
	"fmt"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"math/rand"
)

type OrderInterface struct {
}

func (r OrderInterface) GetOrder(orderId string) (entities.Order, error) {
	var water = entities.Product{
		ID:          1,
		Name:        "Água",
		Price:       1,
		Description: "Água sem gás",
	}

	var sandwich = entities.Product{
		ID:          2,
		Name:        "Sanduíche",
		Price:       2,
		Description: "Sanduíche com carne, queijo e alface",
	}

	var waterInsideOrder = entities.ProductInsideOrder{
		Product:  water,
		Quantity: 2,
	}

	var sandwichInsideOrder = entities.ProductInsideOrder{
		Product:     sandwich,
		Quantity:    1,
		Observation: "Sem queijo",
	}

	var order = entities.Order{
		ID:       fmt.Sprintf("aaaa%d", rand.Int()),
		Products: []entities.ProductInsideOrder{waterInsideOrder, sandwichInsideOrder},
	}

	return order, nil
}
