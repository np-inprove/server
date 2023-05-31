// Code generated by mockery v2.28.1. DO NOT EDIT.

package transactor

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTransactor is an autogenerated mock type for the Transactor type
type MockTransactor struct {
	mock.Mock
}

type MockTransactor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransactor) EXPECT() *MockTransactor_Expecter {
	return &MockTransactor_Expecter{mock: &_m.Mock}
}

// WithTx provides a mock function with given fields: _a0, _a1
func (_m *MockTransactor) WithTx(_a0 context.Context, _a1 func(context.Context) error) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransactor_WithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTx'
type MockTransactor_WithTx_Call struct {
	*mock.Call
}

// WithTx is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 func(context.Context) error
func (_e *MockTransactor_Expecter) WithTx(_a0 interface{}, _a1 interface{}) *MockTransactor_WithTx_Call {
	return &MockTransactor_WithTx_Call{Call: _e.mock.On("WithTx", _a0, _a1)}
}

func (_c *MockTransactor_WithTx_Call) Run(run func(_a0 context.Context, _a1 func(context.Context) error)) *MockTransactor_WithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) error))
	})
	return _c
}

func (_c *MockTransactor_WithTx_Call) Return(_a0 error) *MockTransactor_WithTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransactor_WithTx_Call) RunAndReturn(run func(context.Context, func(context.Context) error) error) *MockTransactor_WithTx_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockTransactor interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTransactor creates a new instance of MockTransactor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTransactor(t mockConstructorTestingTNewMockTransactor) *MockTransactor {
	mock := &MockTransactor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
