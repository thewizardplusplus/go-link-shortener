// Code generated by mockery v1.0.0. DO NOT EDIT.

package httputils

import mock "github.com/stretchr/testify/mock"

// MockPrinter is an autogenerated mock type for the Printer type
type MockPrinter struct {
	mock.Mock
}

// Printf provides a mock function with given fields: template, arguments
func (_m *MockPrinter) Printf(template string, arguments ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, arguments...)
	_m.Called(_ca...)
}
