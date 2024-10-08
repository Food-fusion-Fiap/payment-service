package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/models"
	"log"
	"strings"
)

type PaymentRepository struct {
}

func (r PaymentRepository) Create(e entities.Payment) (string, error) {
	payment := models.Payment{
		OrderID:       e.OrderID,
		QrCode:        e.QrCode,
		PaymentStatus: e.PaymentStatus,
	}

	if err := gorm.DB.Create(&payment).Error; err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return "", errors.New("pagamento já existe no sistema")
		} else {
			return "", errors.New("ocorreu um erro desconhecido ao criar o pagamento")
		}
	}

	return "Criado com sucesso", nil
}

func (r PaymentRepository) FindByOrderId(orderId string) (entities.Payment, error) {
	var payment models.Payment
	//se o orderId tiver mais que um QRCode associado, pega o último
	gorm.DB.Where("order_id = ?", orderId).Order("created_at DESC").Limit(1).Find(&payment)
	if payment.ID == 0 {
		return entities.Payment{}, errors.New("pagamento associado ao id do pedido não encontado")
	}

	result := payment.ToDomain()
	return result, nil
}

func (r PaymentRepository) UpdateToPaid(paymentID uint) {
	var payment models.Payment
	gorm.DB.First(&payment, paymentID)
	gorm.DB.Model(&payment).Updates(models.Payment{PaymentStatus: enums.Paid})
}

func (r PaymentRepository) FindPaymentsQuantity() (uint, error) {
	var payments []models.Payment
	result := gorm.DB.Find(&payments)
	return uint(result.RowsAffected), result.Error
}
