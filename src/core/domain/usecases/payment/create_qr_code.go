package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type CreateQrCodeUseCase struct {
	PaymentInterface  gateways.PaymentInterface
	PaymentRepository gateways.PaymentRepository
}

func (r *CreateQrCodeUseCase) Execute(order entities.Order) (string, error) {
	var err error

	generatedQrCode, err := r.PaymentInterface.CreatePayment(order)
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
