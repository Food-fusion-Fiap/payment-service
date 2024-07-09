package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type CreateQrCodeUseCase struct {
	PaymentInterface  gateways.PaymentInterface
	PaymentRepository gateways.PaymentRepository
	OrderInterface    gateways.OrderInterface
}

func (r *CreateQrCodeUseCase) Execute(orderId uint) (string, error) {
	var err error

	order, err := r.OrderInterface.GetOrder(orderId)
	if err != nil {
		return "", err
	}

	generatedQrCode, err := r.PaymentInterface.CreatePayment(*order)
	if err != nil {
		return "", err
	}

	payment := entities.Payment{
		OrderID:       order.ID,
		QrCode:        generatedQrCode,
		PaymentStatus: enums.AwaitingPayment,
	}

	_, err = r.PaymentRepository.Create(&payment)
	if err != nil {
		return "", err
	}

	return generatedQrCode, err
}
