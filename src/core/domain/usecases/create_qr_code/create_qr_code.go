package create_qr_code

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"log"
)

type CreateQrCodeUseCase struct {
	PaymentInterface  gateways.PaymentInterface
	PaymentRepository gateways.PaymentRepository
	OrderInterface    gateways.OrderInterface
}

func (r *CreateQrCodeUseCase) ExecuteCreateQrCode(orderId string) (string, error) {
	var err error

	order, err := r.OrderInterface.GetOrder(orderId)
	if err != nil || order.ID == "" {
		log.Print(order)
		log.Print(err)
		return "", errors.New("erro ao recuperar o pedido do order-service")
	}

	generatedQrCode, err := r.PaymentInterface.CreatePayment(order)
	if err != nil {
		log.Println(err)
		return "", err
	}

	payment := entities.Payment{
		OrderID:       order.ID,
		QrCode:        generatedQrCode,
		PaymentStatus: enums.AwaitingPayment,
	}

	_, err = r.PaymentRepository.Create(payment)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return generatedQrCode, err
}
