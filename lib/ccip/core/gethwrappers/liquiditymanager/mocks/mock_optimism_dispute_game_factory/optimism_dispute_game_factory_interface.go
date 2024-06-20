// Code generated by mockery v2.42.2. DO NOT EDIT.

package mock_optimism_dispute_game_factory

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	optimism_dispute_game_factory "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/optimism_dispute_game_factory"
)

// OptimismDisputeGameFactoryInterface is an autogenerated mock type for the OptimismDisputeGameFactoryInterface type
type OptimismDisputeGameFactoryInterface struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *OptimismDisputeGameFactoryInterface) Address() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// FindLatestGames provides a mock function with given fields: opts, _gameType, _start, _n
func (_m *OptimismDisputeGameFactoryInterface) FindLatestGames(opts *bind.CallOpts, _gameType uint32, _start *big.Int, _n *big.Int) ([]optimism_dispute_game_factory.IOptimismDisputeGameFactoryGameSearchResult, error) {
	ret := _m.Called(opts, _gameType, _start, _n)

	if len(ret) == 0 {
		panic("no return value specified for FindLatestGames")
	}

	var r0 []optimism_dispute_game_factory.IOptimismDisputeGameFactoryGameSearchResult
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint32, *big.Int, *big.Int) ([]optimism_dispute_game_factory.IOptimismDisputeGameFactoryGameSearchResult, error)); ok {
		return rf(opts, _gameType, _start, _n)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint32, *big.Int, *big.Int) []optimism_dispute_game_factory.IOptimismDisputeGameFactoryGameSearchResult); ok {
		r0 = rf(opts, _gameType, _start, _n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]optimism_dispute_game_factory.IOptimismDisputeGameFactoryGameSearchResult)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, uint32, *big.Int, *big.Int) error); ok {
		r1 = rf(opts, _gameType, _start, _n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameCount provides a mock function with given fields: opts
func (_m *OptimismDisputeGameFactoryInterface) GameCount(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for GameCount")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (*big.Int, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOptimismDisputeGameFactoryInterface creates a new instance of OptimismDisputeGameFactoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOptimismDisputeGameFactoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *OptimismDisputeGameFactoryInterface {
	mock := &OptimismDisputeGameFactoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
