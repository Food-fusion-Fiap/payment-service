package bdd

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/create_qr_code"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func GetMockedPayment(clientOderId uint, qrCode string) entities.Payment {
	return entities.Payment{
		OrderID:       clientOderId,
		QrCode:        qrCode,
		PaymentStatus: enums.AwaitingPayment,
	}
}

func TestRequestQrCode(t *testing.T) {
	//Dado que o cliente já realizou um pedido e possui o orderId
	clientOderId := uint(905)

	clientOrder := entities.Order{ID: clientOderId}
	qrCode := "qrcodedocliente"
	clientPayment := GetMockedPayment(clientOderId, qrCode)

	prepare := func(t *testing.T, pi *gateways.MockPaymentInterface, pr *gateways.MockPaymentRepository, oi *gateways.MockOrderInterface) {
		t.Helper()
		oi.On("GetOrder", clientOderId).Return(clientOrder, nil)
		pi.On("CreatePayment", clientOrder).Return(qrCode, nil)
		pr.On("Create", clientPayment).Return(mock.Anything, nil)
	}

	t.Run("create qr code sucess case", func(t *testing.T) {
		paymentInterfaceMock := &gateways.MockPaymentInterface{}
		paymentRepositoryMock := &gateways.MockPaymentRepository{}
		orderInterfaceMock := &gateways.MockOrderInterface{}

		prepare(t, paymentInterfaceMock, paymentRepositoryMock, orderInterfaceMock)

		usecase := create_qr_code.CreateQrCodeUseCase{
			PaymentInterface:  paymentInterfaceMock,
			PaymentRepository: paymentRepositoryMock,
			OrderInterface:    orderInterfaceMock,
		}

		//Quando ele mandar o orderId para o microserviço de pagamento
		output, _ := usecase.ExecuteCreateQrCode(clientOderId)

		//Então ele deverá recber um código QR
		assert.Equal(t, qrCode, output)
	})
}
