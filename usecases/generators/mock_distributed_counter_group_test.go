// Code generated by mockery v1.0.0. DO NOT EDIT.

package generators

import counters "github.com/thewizardplusplus/go-link-shortener-backend/usecases/generators/counters"
import mock "github.com/stretchr/testify/mock"

// MockDistributedCounterGroup is an autogenerated mock type for the DistributedCounterGroup type
type MockDistributedCounterGroup struct {
	mock.Mock
}

// SelectCounter provides a mock function with given fields:
func (_m *MockDistributedCounterGroup) SelectCounter() counters.DistributedCounter {
	ret := _m.Called()

	var r0 counters.DistributedCounter
	if rf, ok := ret.Get(0).(func() counters.DistributedCounter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(counters.DistributedCounter)
		}
	}

	return r0
}