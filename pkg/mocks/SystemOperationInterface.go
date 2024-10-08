// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ebanfa/skeleton/pkg/common"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ebanfa/skeleton/pkg/types"
)

// SystemOperationInterface is an autogenerated mock type for the SystemOperationInterface type
type SystemOperationInterface struct {
	mock.Mock
}

// Description provides a mock function with given fields:
func (_m *SystemOperationInterface) Description() string {
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

// Execute provides a mock function with given fields: ctx, input
func (_m *SystemOperationInterface) Execute(ctx *common.Context, input *types.SystemOperationInput) (*types.SystemOperationOutput, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *types.SystemOperationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(*common.Context, *types.SystemOperationInput) (*types.SystemOperationOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(*common.Context, *types.SystemOperationInput) *types.SystemOperationOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.SystemOperationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(*common.Context, *types.SystemOperationInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ID provides a mock function with given fields:
func (_m *SystemOperationInterface) ID() string {
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

// Initialize provides a mock function with given fields: ctx, system
func (_m *SystemOperationInterface) Initialize(ctx *common.Context, system types.SystemInterface) error {
	ret := _m.Called(ctx, system)

	if len(ret) == 0 {
		panic("no return value specified for Initialize")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Context, types.SystemInterface) error); ok {
		r0 = rf(ctx, system)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *SystemOperationInterface) Name() string {
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
func (_m *SystemOperationInterface) Type() types.ComponentType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 types.ComponentType
	if rf, ok := ret.Get(0).(func() types.ComponentType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.ComponentType)
	}

	return r0
}

// NewSystemOperationInterface creates a new instance of SystemOperationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSystemOperationInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SystemOperationInterface {
	mock := &SystemOperationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
