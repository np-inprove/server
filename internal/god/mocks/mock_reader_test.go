// Code generated by mockery v2.28.1. DO NOT EDIT.

package god

import (
	context "context"

	ent "github.com/np-inprove/server/internal/ent"

	mock "github.com/stretchr/testify/mock"
)

// MockReader is an autogenerated mock type for the Reader type
type MockReader struct {
	mock.Mock
}

type MockReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReader) EXPECT() *MockReader_Expecter {
	return &MockReader_Expecter{mock: &_m.Mock}
}

// FindInstitutions provides a mock function with given fields: ctx
func (_m *MockReader) FindInstitutions(ctx context.Context) ([]*ent.Institution, error) {
	ret := _m.Called(ctx)

	var r0 []*ent.Institution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*ent.Institution, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*ent.Institution); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Institution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindInstitutions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitutions'
type MockReader_FindInstitutions_Call struct {
	*mock.Call
}

// FindInstitutions is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockReader_Expecter) FindInstitutions(ctx interface{}) *MockReader_FindInstitutions_Call {
	return &MockReader_FindInstitutions_Call{Call: _e.mock.On("FindInstitutions", ctx)}
}

func (_c *MockReader_FindInstitutions_Call) Run(run func(ctx context.Context)) *MockReader_FindInstitutions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockReader_FindInstitutions_Call) Return(_a0 []*ent.Institution, _a1 error) *MockReader_FindInstitutions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindInstitutions_Call) RunAndReturn(run func(context.Context) ([]*ent.Institution, error)) *MockReader_FindInstitutions_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockReader creates a new instance of MockReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockReader(t mockConstructorTestingTNewMockReader) *MockReader {
	mock := &MockReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
