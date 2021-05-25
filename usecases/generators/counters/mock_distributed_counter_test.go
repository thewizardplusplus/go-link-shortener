// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package counters

import mock "github.com/stretchr/testify/mock"

// MockDistributedCounter is an autogenerated mock type for the DistributedCounter type
type MockDistributedCounter struct {
	mock.Mock
}

// NextCountChunk provides a mock function with given fields:
func (_m *MockDistributedCounter) NextCountChunk() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
