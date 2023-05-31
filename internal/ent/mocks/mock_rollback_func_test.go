// Code generated by mockery v2.28.1. DO NOT EDIT.

package ent

import (
	context "context"

	ent "github.com/np-inprove/server/internal/ent"
	mock "github.com/stretchr/testify/mock"
)

// MockRollbackFunc is an autogenerated mock type for the RollbackFunc type
type MockRollbackFunc struct {
	mock.Mock
}

type MockRollbackFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRollbackFunc) EXPECT() *MockRollbackFunc_Expecter {
	return &MockRollbackFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockRollbackFunc) Execute(_a0 context.Context, _a1 *ent.Tx) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRollbackFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockRollbackFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *ent.Tx
func (_e *MockRollbackFunc_Expecter) Execute(_a0 interface{}, _a1 interface{}) *MockRollbackFunc_Execute_Call {
	return &MockRollbackFunc_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1)}
}

func (_c *MockRollbackFunc_Execute_Call) Run(run func(_a0 context.Context, _a1 *ent.Tx)) *MockRollbackFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*ent.Tx))
	})
	return _c
}

func (_c *MockRollbackFunc_Execute_Call) Return(_a0 error) *MockRollbackFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollbackFunc_Execute_Call) RunAndReturn(run func(context.Context, *ent.Tx) error) *MockRollbackFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRollbackFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRollbackFunc creates a new instance of MockRollbackFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRollbackFunc(t mockConstructorTestingTNewMockRollbackFunc) *MockRollbackFunc {
	mock := &MockRollbackFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
