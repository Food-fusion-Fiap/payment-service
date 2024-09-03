package order_service

import (
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

	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	log.Println(resp)
	log.Println(resp.Body)
	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err, "Erro ao ler body do order-service")
			log.Panic(err, "Erro ao ler body do order-service")
		}
	}(resp.Body)

	return entities.Order{ID: "aaa"}, nil
}
