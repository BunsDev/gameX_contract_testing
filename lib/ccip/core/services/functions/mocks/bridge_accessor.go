// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	functions "github.com/smartcontractkit/chainlink/v2/core/services/functions"
	mock "github.com/stretchr/testify/mock"
)

// BridgeAccessor is an autogenerated mock type for the BridgeAccessor type
type BridgeAccessor struct {
	mock.Mock
}

// NewExternalAdapterClient provides a mock function with given fields: _a0
func (_m *BridgeAccessor) NewExternalAdapterClient(_a0 context.Context) (functions.ExternalAdapterClient, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for NewExternalAdapterClient")
	}

	var r0 functions.ExternalAdapterClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (functions.ExternalAdapterClient, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) functions.ExternalAdapterClient); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(functions.ExternalAdapterClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBridgeAccessor creates a new instance of BridgeAccessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBridgeAccessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *BridgeAccessor {
	mock := &BridgeAccessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
