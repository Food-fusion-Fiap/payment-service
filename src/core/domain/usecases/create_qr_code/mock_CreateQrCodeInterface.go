// Code generated by mockery v2.43.2. DO NOT EDIT.

package create_qr_code

import mock "github.com/stretchr/testify/mock"

// MockCreateQrCodeInterface is an autogenerated mock type for the CreateQrCodeInterface type
type MockCreateQrCodeInterface struct {
	mock.Mock
}

// ExecuteCreateQrCode provides a mock function with given fields: orderId
func (_m *MockCreateQrCodeInterface) ExecuteCreateQrCode(orderId string) (string, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteCreateQrCode")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockCreateQrCodeInterface creates a new instance of MockCreateQrCodeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreateQrCodeInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreateQrCodeInterface {
	mock := &MockCreateQrCodeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
