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

func (r OrderInterface) GetOrder(orderId uint) (entities.Order, error) {
	//TODO: terminar
	resp, err := http.Get("http://order-service.svc.cluster.local:30202/")
	if err != nil {
		fmt.Println(err, "Erro ao conectar com order-service")
		log.Panic(err, "Erro ao conectar com order-service")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err, "Erro ao ler body do order-service")
			log.Panic(err, "Erro ao ler body do order-service")
		}
	}(resp.Body)

	return entities.Order{ID: 1}, nil
}
