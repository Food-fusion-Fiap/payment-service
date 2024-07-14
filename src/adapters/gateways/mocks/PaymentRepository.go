// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	entities "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// PaymentRepository is an autogenerated mock type for the PaymentRepository type
type PaymentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: order
func (_m *PaymentRepository) Create(order entities.Payment) (string, error) {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Payment) (string, error)); ok {
		return rf(order)
	}
	if rf, ok := ret.Get(0).(func(entities.Payment) string); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(entities.Payment) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByOrderId provides a mock function with given fields: orderId
func (_m *PaymentRepository) FindByOrderId(orderId uint) (entities.Payment, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for FindByOrderId")
	}

	var r0 entities.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entities.Payment, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(uint) entities.Payment); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Get(0).(entities.Payment)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByQrCode provides a mock function with given fields: qrCode
func (_m *PaymentRepository) FindByQrCode(qrCode string) (entities.Payment, error) {
	ret := _m.Called(qrCode)

	if len(ret) == 0 {
		panic("no return value specified for FindByQrCode")
	}

	var r0 entities.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entities.Payment, error)); ok {
		return rf(qrCode)
	}
	if rf, ok := ret.Get(0).(func(string) entities.Payment); ok {
		r0 = rf(qrCode)
	} else {
		r0 = ret.Get(0).(entities.Payment)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(qrCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPaymentsQuantity provides a mock function with given fields:
func (_m *PaymentRepository) FindPaymentsQuantity() (uint, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindPaymentsQuantity")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateToPaid provides a mock function with given fields: paymentID
func (_m *PaymentRepository) UpdateToPaid(paymentID uint) {
	_m.Called(paymentID)
}

// NewPaymentRepository creates a new instance of PaymentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentRepository {
	mock := &PaymentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
