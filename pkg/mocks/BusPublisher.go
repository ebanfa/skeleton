// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ebanfa/skeleton/pkg/common"
	mock "github.com/stretchr/testify/mock"
)

// BusPublisher is an autogenerated mock type for the BusPublisher type
type BusPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: event
func (_m *BusPublisher) Publish(event common.Event) {
	_m.Called(event)
}

// NewBusPublisher creates a new instance of BusPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBusPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *BusPublisher {
	mock := &BusPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
