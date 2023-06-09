// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

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

// FindInstitutionInviteLinkWithInstitution provides a mock function with given fields: ctx, code
func (_m *MockReader) FindInstitutionInviteLinkWithInstitution(ctx context.Context, code string) (*ent.InstitutionInviteLink, error) {
	ret := _m.Called(ctx, code)

	var r0 *ent.InstitutionInviteLink
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.InstitutionInviteLink, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.InstitutionInviteLink); ok {
		r0 = rf(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.InstitutionInviteLink)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindInstitutionInviteLinkWithInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitutionInviteLinkWithInstitution'
type MockReader_FindInstitutionInviteLinkWithInstitution_Call struct {
	*mock.Call
}

// FindInstitutionInviteLinkWithInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - code string
func (_e *MockReader_Expecter) FindInstitutionInviteLinkWithInstitution(ctx interface{}, code interface{}) *MockReader_FindInstitutionInviteLinkWithInstitution_Call {
	return &MockReader_FindInstitutionInviteLinkWithInstitution_Call{Call: _e.mock.On("FindInstitutionInviteLinkWithInstitution", ctx, code)}
}

func (_c *MockReader_FindInstitutionInviteLinkWithInstitution_Call) Run(run func(ctx context.Context, code string)) *MockReader_FindInstitutionInviteLinkWithInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReader_FindInstitutionInviteLinkWithInstitution_Call) Return(_a0 *ent.InstitutionInviteLink, _a1 error) *MockReader_FindInstitutionInviteLinkWithInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindInstitutionInviteLinkWithInstitution_Call) RunAndReturn(run func(context.Context, string) (*ent.InstitutionInviteLink, error)) *MockReader_FindInstitutionInviteLinkWithInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// FindJWTRevocation provides a mock function with given fields: ctx, jti
func (_m *MockReader) FindJWTRevocation(ctx context.Context, jti string) (*ent.JWTRevocation, error) {
	ret := _m.Called(ctx, jti)

	var r0 *ent.JWTRevocation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.JWTRevocation, error)); ok {
		return rf(ctx, jti)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.JWTRevocation); ok {
		r0 = rf(ctx, jti)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.JWTRevocation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, jti)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindJWTRevocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindJWTRevocation'
type MockReader_FindJWTRevocation_Call struct {
	*mock.Call
}

// FindJWTRevocation is a helper method to define mock.On call
//   - ctx context.Context
//   - jti string
func (_e *MockReader_Expecter) FindJWTRevocation(ctx interface{}, jti interface{}) *MockReader_FindJWTRevocation_Call {
	return &MockReader_FindJWTRevocation_Call{Call: _e.mock.On("FindJWTRevocation", ctx, jti)}
}

func (_c *MockReader_FindJWTRevocation_Call) Run(run func(ctx context.Context, jti string)) *MockReader_FindJWTRevocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReader_FindJWTRevocation_Call) Return(_a0 *ent.JWTRevocation, _a1 error) *MockReader_FindJWTRevocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindJWTRevocation_Call) RunAndReturn(run func(context.Context, string) (*ent.JWTRevocation, error)) *MockReader_FindJWTRevocation_Call {
	_c.Call.Return(run)
	return _c
}

// FindUserByEmailWithInstitution provides a mock function with given fields: ctx, email
func (_m *MockReader) FindUserByEmailWithInstitution(ctx context.Context, email string) (*ent.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindUserByEmailWithInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUserByEmailWithInstitution'
type MockReader_FindUserByEmailWithInstitution_Call struct {
	*mock.Call
}

// FindUserByEmailWithInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockReader_Expecter) FindUserByEmailWithInstitution(ctx interface{}, email interface{}) *MockReader_FindUserByEmailWithInstitution_Call {
	return &MockReader_FindUserByEmailWithInstitution_Call{Call: _e.mock.On("FindUserByEmailWithInstitution", ctx, email)}
}

func (_c *MockReader_FindUserByEmailWithInstitution_Call) Run(run func(ctx context.Context, email string)) *MockReader_FindUserByEmailWithInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReader_FindUserByEmailWithInstitution_Call) Return(_a0 *ent.User, _a1 error) *MockReader_FindUserByEmailWithInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindUserByEmailWithInstitution_Call) RunAndReturn(run func(context.Context, string) (*ent.User, error)) *MockReader_FindUserByEmailWithInstitution_Call {
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
