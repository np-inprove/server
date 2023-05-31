// Code generated by mockery v2.28.1. DO NOT EDIT.

package ent

import (
	ent "github.com/np-inprove/server/internal/ent"
	mock "github.com/stretchr/testify/mock"
)

// MockuserOption is an autogenerated mock type for the userOption type
type MockuserOption struct {
	mock.Mock
}

type MockuserOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockuserOption) EXPECT() *MockuserOption_Expecter {
	return &MockuserOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockuserOption) Execute(_a0 *ent.UserMutation) {
	_m.Called(_a0)
}

// MockuserOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockuserOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *ent.UserMutation
func (_e *MockuserOption_Expecter) Execute(_a0 interface{}) *MockuserOption_Execute_Call {
	return &MockuserOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockuserOption_Execute_Call) Run(run func(_a0 *ent.UserMutation)) *MockuserOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ent.UserMutation))
	})
	return _c
}

func (_c *MockuserOption_Execute_Call) Return() *MockuserOption_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockuserOption_Execute_Call) RunAndReturn(run func(*ent.UserMutation)) *MockuserOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockuserOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockuserOption creates a new instance of MockuserOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockuserOption(t mockConstructorTestingTNewMockuserOption) *MockuserOption {
	mock := &MockuserOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
