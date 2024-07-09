package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/models"
	"strings"
)

type PaymentRepository struct {
}

func (r PaymentRepository) Create(e *entities.Payment) (*entities.Payment, error) {
	payment := models.Payment{
		OrderID:       e.OrderID,
		QrCode:        e.QrCode,
		PaymentStatus: e.PaymentStatus,
	}

	if err := gorm.DB.Create(&payment).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("pagamento já existe no sistema")
		} else {
			return nil, errors.New("ocorreu um erro desconhecido ao criar o pagamento")
		}
	}

	result := payment.ToDomain()

	return &result, nil
}

func (r PaymentRepository) FindByOrderId(orderId uint) (*entities.Payment, error) {
	var payment models.Payment
	gorm.DB.Where("order_id = ?", orderId).Find(&payment)
	if payment.ID != 0 {
		return nil, errors.New("pagamento associado ao id do pedido não encontado")
	}
	result := payment.ToDomain()
	return &result, nil
}

func (r PaymentRepository) FindByQrCode(qrCode string) (*entities.Payment, error) {
	var payment models.Payment

	gorm.DB.Where("payment_status = ? and qr_code = ?", enums.Paid, qrCode).Find(&payment)
	if payment.ID != 0 {
		return nil, errors.New("já está pago")
	}

	gorm.DB.Where("payment_status = ? and qr_code = ?", enums.AwaitingPayment, qrCode).Find(&payment)
	result := payment.ToDomain()
	return &result, nil
}