// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ebanfa/skeleton/pkg/common"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ebanfa/skeleton/pkg/types"
)

// SystemInterface is an autogenerated mock type for the SystemInterface type
type SystemInterface struct {
	mock.Mock
}

// ComponentRegistry provides a mock function with given fields:
func (_m *SystemInterface) ComponentRegistry() types.ComponentRegistrarInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ComponentRegistry")
	}

	var r0 types.ComponentRegistrarInterface
	if rf, ok := ret.Get(0).(func() types.ComponentRegistrarInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ComponentRegistrarInterface)
		}
	}

	return r0
}

// Configuration provides a mock function with given fields:
func (_m *SystemInterface) Configuration() *types.Configuration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Configuration")
	}

	var r0 *types.Configuration
	if rf, ok := ret.Get(0).(func() *types.Configuration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Configuration)
		}
	}

	return r0
}

// EventBus provides a mock function with given fields:
func (_m *SystemInterface) EventBus() common.EventBusInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EventBus")
	}

	var r0 common.EventBusInterface
	if rf, ok := ret.Get(0).(func() common.EventBusInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.EventBusInterface)
		}
	}

	return r0
}

// ExecuteOperation provides a mock function with given fields: ctx, operationID, data
func (_m *SystemInterface) ExecuteOperation(ctx *common.Context, operationID string, data *types.SystemOperationInput) (*types.SystemOperationOutput, error) {
	ret := _m.Called(ctx, operationID, data)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteOperation")
	}

	var r0 *types.SystemOperationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(*common.Context, string, *types.SystemOperationInput) (*types.SystemOperationOutput, error)); ok {
		return rf(ctx, operationID, data)
	}
	if rf, ok := ret.Get(0).(func(*common.Context, string, *types.SystemOperationInput) *types.SystemOperationOutput); ok {
		r0 = rf(ctx, operationID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.SystemOperationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(*common.Context, string, *types.SystemOperationInput) error); ok {
		r1 = rf(ctx, operationID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: ctx
func (_m *SystemInterface) Initialize(ctx *common.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Initialize")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *SystemInterface) Logger() common.LoggerInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Logger")
	}

	var r0 common.LoggerInterface
	if rf, ok := ret.Get(0).(func() common.LoggerInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.LoggerInterface)
		}
	}

	return r0
}

// MultiStore provides a mock function with given fields:
func (_m *SystemInterface) MultiStore() types.MultiStore {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MultiStore")
	}

	var r0 types.MultiStore
	if rf, ok := ret.Get(0).(func() types.MultiStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.MultiStore)
		}
	}

	return r0
}

// PluginManager provides a mock function with given fields:
func (_m *SystemInterface) PluginManager() types.PluginManagerInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PluginManager")
	}

	var r0 types.PluginManagerInterface
	if rf, ok := ret.Get(0).(func() types.PluginManagerInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.PluginManagerInterface)
		}
	}

	return r0
}

// RestartService provides a mock function with given fields: ctx, serviceID
func (_m *SystemInterface) RestartService(ctx *common.Context, serviceID string) error {
	ret := _m.Called(ctx, serviceID)

	if len(ret) == 0 {
		panic("no return value specified for RestartService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Context, string) error); ok {
		r0 = rf(ctx, serviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *SystemInterface) Start(ctx *common.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartService provides a mock function with given fields: ctx, serviceID
func (_m *SystemInterface) StartService(ctx *common.Context, serviceID string) error {
	ret := _m.Called(ctx, serviceID)

	if len(ret) == 0 {
		panic("no return value specified for StartService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Context, string) error); ok {
		r0 = rf(ctx, serviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: ctx
func (_m *SystemInterface) Stop(ctx *common.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopService provides a mock function with given fields: ctx, serviceID
func (_m *SystemInterface) StopService(ctx *common.Context, serviceID string) error {
	ret := _m.Called(ctx, serviceID)

	if len(ret) == 0 {
		panic("no return value specified for StopService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Context, string) error); ok {
		r0 = rf(ctx, serviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSystemInterface creates a new instance of SystemInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSystemInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SystemInterface {
	mock := &SystemInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
