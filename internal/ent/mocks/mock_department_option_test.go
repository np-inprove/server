// Code generated by mockery v2.28.1. DO NOT EDIT.

package ent

import (
	ent "github.com/np-inprove/server/internal/ent"
	mock "github.com/stretchr/testify/mock"
)

// MockdepartmentOption is an autogenerated mock type for the departmentOption type
type MockdepartmentOption struct {
	mock.Mock
}

type MockdepartmentOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockdepartmentOption) EXPECT() *MockdepartmentOption_Expecter {
	return &MockdepartmentOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockdepartmentOption) Execute(_a0 *ent.DepartmentMutation) {
	_m.Called(_a0)
}

// MockdepartmentOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockdepartmentOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *ent.DepartmentMutation
func (_e *MockdepartmentOption_Expecter) Execute(_a0 interface{}) *MockdepartmentOption_Execute_Call {
	return &MockdepartmentOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockdepartmentOption_Execute_Call) Run(run func(_a0 *ent.DepartmentMutation)) *MockdepartmentOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ent.DepartmentMutation))
	})
	return _c
}

func (_c *MockdepartmentOption_Execute_Call) Return() *MockdepartmentOption_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockdepartmentOption_Execute_Call) RunAndReturn(run func(*ent.DepartmentMutation)) *MockdepartmentOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockdepartmentOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockdepartmentOption creates a new instance of MockdepartmentOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockdepartmentOption(t mockConstructorTestingTNewMockdepartmentOption) *MockdepartmentOption {
	mock := &MockdepartmentOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
