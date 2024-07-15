// Code generated by mockery v2.43.2. DO NOT EDIT.

package make_payment

import (
	entities "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// MakePaymentInterface is an autogenerated mock type for the MakePaymentInterface type
type MakePaymentInterface struct {
	mock.Mock
}

// ExecuteWithOrderId provides a mock function with given fields: orderId
func (_m *MakePaymentInterface) ExecuteWithOrderId(orderId uint) (string, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteWithOrderId")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (string, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteWithQrCode provides a mock function with given fields: qrCode
func (_m *MakePaymentInterface) ExecuteWithQrCode(qrCode string) (string, error) {
	ret := _m.Called(qrCode)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteWithQrCode")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(qrCode)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(qrCode)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(qrCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: payment
func (_m *MakePaymentInterface) UpdateStatus(payment entities.Payment) (string, error) {
	ret := _m.Called(payment)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Payment) (string, error)); ok {
		return rf(payment)
	}
	if rf, ok := ret.Get(0).(func(entities.Payment) string); ok {
		r0 = rf(payment)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(entities.Payment) error); ok {
		r1 = rf(payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMakePaymentInterface creates a new instance of MakePaymentInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMakePaymentInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MakePaymentInterface {
	mock := &MakePaymentInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
