// Code generated by mockery v2.28.1. DO NOT EDIT.

package ent

import (
	ent "github.com/np-inprove/server/internal/ent"
	mock "github.com/stretchr/testify/mock"
)

// MockuserpetOption is an autogenerated mock type for the userpetOption type
type MockuserpetOption struct {
	mock.Mock
}

type MockuserpetOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockuserpetOption) EXPECT() *MockuserpetOption_Expecter {
	return &MockuserpetOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockuserpetOption) Execute(_a0 *ent.UserPetMutation) {
	_m.Called(_a0)
}

// MockuserpetOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockuserpetOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *ent.UserPetMutation
func (_e *MockuserpetOption_Expecter) Execute(_a0 interface{}) *MockuserpetOption_Execute_Call {
	return &MockuserpetOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockuserpetOption_Execute_Call) Run(run func(_a0 *ent.UserPetMutation)) *MockuserpetOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ent.UserPetMutation))
	})
	return _c
}

func (_c *MockuserpetOption_Execute_Call) Return() *MockuserpetOption_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockuserpetOption_Execute_Call) RunAndReturn(run func(*ent.UserPetMutation)) *MockuserpetOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockuserpetOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockuserpetOption creates a new instance of MockuserpetOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockuserpetOption(t mockConstructorTestingTNewMockuserpetOption) *MockuserpetOption {
	mock := &MockuserpetOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
