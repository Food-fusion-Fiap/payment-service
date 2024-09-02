package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/utils"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID       string
	QrCode        string
	PaymentStatus string
}

func (o Payment) ToDomain() entities.Payment {
	return entities.Payment{
		ID:            o.ID,
		OrderID:       o.OrderID,
		QrCode:        o.QrCode,
		PaymentStatus: o.PaymentStatus,
		CreatedAt:     o.CreatedAt.Format(utils.CompleteEnglishDateFormat),
	}
}
