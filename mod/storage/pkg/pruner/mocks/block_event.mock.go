// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	context "context"

	async "github.com/berachain/beacon-kit/mod/primitives/pkg/async"

	mock "github.com/stretchr/testify/mock"

	pruner "github.com/berachain/beacon-kit/mod/storage/pkg/pruner"
)

// BlockEvent is an autogenerated mock type for the BlockEvent type
type BlockEvent[BeaconBlockT pruner.BeaconBlock] struct {
	mock.Mock
}

type BlockEvent_Expecter[BeaconBlockT pruner.BeaconBlock] struct {
	mock *mock.Mock
}

func (_m *BlockEvent[BeaconBlockT]) EXPECT() *BlockEvent_Expecter[BeaconBlockT] {
	return &BlockEvent_Expecter[BeaconBlockT]{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *BlockEvent[BeaconBlockT]) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// BlockEvent_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type BlockEvent_Context_Call[BeaconBlockT pruner.BeaconBlock] struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *BlockEvent_Expecter[BeaconBlockT]) Context() *BlockEvent_Context_Call[BeaconBlockT] {
	return &BlockEvent_Context_Call[BeaconBlockT]{Call: _e.mock.On("Context")}
}

func (_c *BlockEvent_Context_Call[BeaconBlockT]) Run(run func()) *BlockEvent_Context_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockEvent_Context_Call[BeaconBlockT]) Return(_a0 context.Context) *BlockEvent_Context_Call[BeaconBlockT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Context_Call[BeaconBlockT]) RunAndReturn(run func() context.Context) *BlockEvent_Context_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// Data provides a mock function with given fields:
func (_m *BlockEvent[BeaconBlockT]) Data() BeaconBlockT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 BeaconBlockT
	if rf, ok := ret.Get(0).(func() BeaconBlockT); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(BeaconBlockT)
		}
	}

	return r0
}

// BlockEvent_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type BlockEvent_Data_Call[BeaconBlockT pruner.BeaconBlock] struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *BlockEvent_Expecter[BeaconBlockT]) Data() *BlockEvent_Data_Call[BeaconBlockT] {
	return &BlockEvent_Data_Call[BeaconBlockT]{Call: _e.mock.On("Data")}
}

func (_c *BlockEvent_Data_Call[BeaconBlockT]) Run(run func()) *BlockEvent_Data_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockEvent_Data_Call[BeaconBlockT]) Return(_a0 BeaconBlockT) *BlockEvent_Data_Call[BeaconBlockT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Data_Call[BeaconBlockT]) RunAndReturn(run func() BeaconBlockT) *BlockEvent_Data_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with given fields:
func (_m *BlockEvent[BeaconBlockT]) ID() async.EventID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 async.EventID
	if rf, ok := ret.Get(0).(func() async.EventID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(async.EventID)
	}

	return r0
}

// BlockEvent_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type BlockEvent_ID_Call[BeaconBlockT pruner.BeaconBlock] struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *BlockEvent_Expecter[BeaconBlockT]) ID() *BlockEvent_ID_Call[BeaconBlockT] {
	return &BlockEvent_ID_Call[BeaconBlockT]{Call: _e.mock.On("ID")}
}

func (_c *BlockEvent_ID_Call[BeaconBlockT]) Run(run func()) *BlockEvent_ID_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockEvent_ID_Call[BeaconBlockT]) Return(_a0 async.EventID) *BlockEvent_ID_Call[BeaconBlockT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_ID_Call[BeaconBlockT]) RunAndReturn(run func() async.EventID) *BlockEvent_ID_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// Is provides a mock function with given fields: _a0
func (_m *BlockEvent[BeaconBlockT]) Is(_a0 async.EventID) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Is")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(async.EventID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BlockEvent_Is_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Is'
type BlockEvent_Is_Call[BeaconBlockT pruner.BeaconBlock] struct {
	*mock.Call
}

// Is is a helper method to define mock.On call
//   - _a0 async.EventID
func (_e *BlockEvent_Expecter[BeaconBlockT]) Is(_a0 interface{}) *BlockEvent_Is_Call[BeaconBlockT] {
	return &BlockEvent_Is_Call[BeaconBlockT]{Call: _e.mock.On("Is", _a0)}
}

func (_c *BlockEvent_Is_Call[BeaconBlockT]) Run(run func(_a0 async.EventID)) *BlockEvent_Is_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(async.EventID))
	})
	return _c
}

func (_c *BlockEvent_Is_Call[BeaconBlockT]) Return(_a0 bool) *BlockEvent_Is_Call[BeaconBlockT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Is_Call[BeaconBlockT]) RunAndReturn(run func(async.EventID) bool) *BlockEvent_Is_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// NewBlockEvent creates a new instance of BlockEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockEvent[BeaconBlockT pruner.BeaconBlock](t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockEvent[BeaconBlockT] {
	mock := &BlockEvent[BeaconBlockT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
