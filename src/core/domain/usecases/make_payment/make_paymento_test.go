package make_payment

import (
	"errors"
	"testing"

	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"

	"github.com/stretchr/testify/assert"
)

//func TestMakePaymentUseCase_FindByQrCodeAndUpdatePayment_Success(t *testing.T) {
//	mockQrCode := "mockQrCode"
//	mockPayment := entities.Payment{PaymentStatus: enums.AwaitingPayment, ID: 2, OrderID: 30}
//
//	prepare := func(t *testing.T, pr *gateways.MockPaymentRepository, pb *gateways.MockPubSubInterface) {
//		t.Helper()
//		pr.On("FindByQrCode", mockQrCode).Return(mockPayment, nil)
//		pr.On("UpdateToPaid", mockPayment.ID).Return()
//		pb.On("NotifyStatusChange", mockPayment.OrderID).Return(nil)
//	}
//
//	t.Run("TestMakePaymentUseCase_FindByQrCodeAndUpdatePayment_Success", func(t *testing.T) {
//		paymentRepositoryMock := &gateways.MockPaymentRepository{}
//		orderInterfaceMock := &gateways.MockOrderInterface{}
//		pubSubInterfaceMock := &gateways.MockPubSubInterface{}
//		prepare(t, paymentRepositoryMock, pubSubInterfaceMock)
//
//		usecase := MakePaymentUseCase{
//			PaymentRepository: paymentRepositoryMock,
//			OrderInterface:    orderInterfaceMock,
//		}
//
//		output, err := usecase.ExecuteApprovedPaymentWithQrCode(mockQrCode)
//
//		assert.Equal(t, "Pago", output)
//		assert.Equal(t, nil, err)
//	})
//}

func TestMakePaymentUseCase_FindByQrCodeAndUpdatePayment_Fail(t *testing.T) {
	mockQrCode := "mockQrCode"
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *gateways.MockPaymentRepository, pb *gateways.MockPubSubInterface) {
		t.Helper()
		pr.On("FindByQrCode", mockQrCode).Return(mockPayment, nil)
		pr.On("UpdateToPaid", mockPayment.ID).Return()
	}

	t.Run("TestMakePaymentUseCase_FindByQrCodeAndUpdatePayment_Fail", func(t *testing.T) {
		paymentRepositoryMock := &gateways.MockPaymentRepository{}
		orderInterfaceMock := &gateways.MockOrderInterface{}
		pubSubInterfaceMock := &gateways.MockPubSubInterface{}
		prepare(t, paymentRepositoryMock, pubSubInterfaceMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
			OrderInterface:    orderInterfaceMock,
		}

		output, err := usecase.ExecuteApprovedPaymentWithQrCode(mockQrCode)

		assert.Equal(t, output, "")
		assert.Equal(t, errors.New("não foi possível efetuar o pagamento: o pagamento já foi pago"), err)
	})
}

//func TestMakePaymentUseCase_FindByOrderIdAndUpdatePayment_Success(t *testing.T) {
//	mockOrderId := uint(407)
//	mockPayment := entities.Payment{PaymentStatus: enums.AwaitingPayment, ID: 2, OrderID: 30}
//
//	prepare := func(t *testing.T, pr *gateways.MockPaymentRepository, pb *gateways.MockPubSubInterface) {
//		t.Helper()
//		pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
//		pr.On("UpdateToPaid", mockPayment.ID).Return()
//	}
//
//	t.Run("when success FindByOrderId", func(t *testing.T) {
//		paymentRepositoryMock := &gateways.MockPaymentRepository{}
//		orderInterfaceMock := &gateways.MockOrderInterface{}
//		pubSubInterfaceMock := &gateways.MockPubSubInterface{}
//		prepare(t, paymentRepositoryMock, pubSubInterfaceMock)
//
//		usecase := MakePaymentUseCase{
//			PaymentRepository: paymentRepositoryMock,
//			OrderInterface:    orderInterfaceMock,
//		}
//
//		output, err := usecase.ExecuteApprovedPaymentWithOrderId(mockOrderId)
//
//		assert.Equal(t, "Pago", output)
//		assert.Equal(t, nil, err)
//	})
//}

func TestMakePaymentUseCase_FindByOrderIdAndUpdatePayment_Fail(t *testing.T) {
	mockOrderId := uint(407)
	mockPayment := entities.Payment{PaymentStatus: enums.Paid, ID: 2, OrderID: 30}

	prepare := func(t *testing.T, pr *gateways.MockPaymentRepository, pb *gateways.MockPubSubInterface) {
		t.Helper()
		pr.On("FindByOrderId", mockOrderId).Return(mockPayment, nil)
		pr.On("UpdateToPaid", mockPayment.ID).Return()
	}

	t.Run("when fails findbyQROrderId", func(t *testing.T) {
		paymentRepositoryMock := &gateways.MockPaymentRepository{}
		orderInterfaceMock := &gateways.MockOrderInterface{}
		pubSubInterfaceMock := &gateways.MockPubSubInterface{}
		prepare(t, paymentRepositoryMock, pubSubInterfaceMock)

		usecase := MakePaymentUseCase{
			PaymentRepository: paymentRepositoryMock,
			OrderInterface:    orderInterfaceMock,
		}

		output, err := usecase.ExecuteApprovedPaymentWithOrderId(mockOrderId)

		assert.Equal(t, output, "")
		assert.Equal(t, errors.New("não foi possível efetuar o pagamento: o pagamento já foi pago"), err)
	})
}
