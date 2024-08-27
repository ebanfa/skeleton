// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	system "github.com/ebanfa/skeleton/pkg/system"
	mock "github.com/stretchr/testify/mock"
)

// ComponentInterface is an autogenerated mock type for the ComponentInterface type
type ComponentInterface struct {
	mock.Mock
}

// Description provides a mock function with given fields:
func (_m *ComponentInterface) Description() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Description")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *ComponentInterface) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *ComponentInterface) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *ComponentInterface) Type() system.ComponentType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 system.ComponentType
	if rf, ok := ret.Get(0).(func() system.ComponentType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(system.ComponentType)
	}

	return r0
}

// NewComponentInterface creates a new instance of ComponentInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComponentInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ComponentInterface {
	mock := &ComponentInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
