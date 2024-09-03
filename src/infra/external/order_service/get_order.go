package order_service

import (
	"encoding/json"
	"fmt"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"io"
	"log"
	"net/http"
)

type OrderInterface struct {
}

func (r OrderInterface) GetOrder(orderId string) (entities.Order, error) {
	orderServiceAddress := fmt.Sprintf("http://ada859e8692e840ea9a41533c9f5c8d6-315966617.us-east-1.elb.amazonaws.com/orders/%s", orderId)
	resp, err := http.Get(orderServiceAddress)
	if err != nil {
		fmt.Println(err, "Erro ao conectar com order-service")
		log.Panic(err, "Erro ao conectar com order-service")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	var targetOrder entities.Order
	err = json.Unmarshal(bodyBytes, &targetOrder)
	if err != nil {
		log.Println(err, "Erro deserializar o pedido")
	}

	return targetOrder, nil
}
