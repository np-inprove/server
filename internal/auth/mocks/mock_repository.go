// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	ent "github.com/np-inprove/server/internal/ent"
	hash "github.com/np-inprove/server/internal/hash"

	institution "github.com/np-inprove/server/internal/entity/institution"

	mock "github.com/stretchr/testify/mock"

	time "time"

	user "github.com/np-inprove/server/internal/entity/user"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// CreateJWTRevocation provides a mock function with given fields: ctx, jti, expiry
func (_m *MockRepository) CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (*ent.JWTRevocation, error) {
	ret := _m.Called(ctx, jti, expiry)

	var r0 *ent.JWTRevocation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) (*ent.JWTRevocation, error)); ok {
		return rf(ctx, jti, expiry)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) *ent.JWTRevocation); ok {
		r0 = rf(ctx, jti, expiry)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.JWTRevocation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) error); ok {
		r1 = rf(ctx, jti, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_CreateJWTRevocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJWTRevocation'
type MockRepository_CreateJWTRevocation_Call struct {
	*mock.Call
}

// CreateJWTRevocation is a helper method to define mock.On call
//   - ctx context.Context
//   - jti string
//   - expiry time.Time
func (_e *MockRepository_Expecter) CreateJWTRevocation(ctx interface{}, jti interface{}, expiry interface{}) *MockRepository_CreateJWTRevocation_Call {
	return &MockRepository_CreateJWTRevocation_Call{Call: _e.mock.On("CreateJWTRevocation", ctx, jti, expiry)}
}

func (_c *MockRepository_CreateJWTRevocation_Call) Run(run func(ctx context.Context, jti string, expiry time.Time)) *MockRepository_CreateJWTRevocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Time))
	})
	return _c
}

func (_c *MockRepository_CreateJWTRevocation_Call) Return(_a0 *ent.JWTRevocation, _a1 error) *MockRepository_CreateJWTRevocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_CreateJWTRevocation_Call) RunAndReturn(run func(context.Context, string, time.Time) (*ent.JWTRevocation, error)) *MockRepository_CreateJWTRevocation_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, instID, instRole, firstName, lastName, email, password, opts
func (_m *MockRepository) CreateUser(ctx context.Context, instID int, instRole institution.Role, firstName string, lastName string, email string, password hash.Encoded, opts ...user.Option) (*ent.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, instID, instRole, firstName, lastName, email, password)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, institution.Role, string, string, string, hash.Encoded, ...user.Option) (*ent.User, error)); ok {
		return rf(ctx, instID, instRole, firstName, lastName, email, password, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, institution.Role, string, string, string, hash.Encoded, ...user.Option) *ent.User); ok {
		r0 = rf(ctx, instID, instRole, firstName, lastName, email, password, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, institution.Role, string, string, string, hash.Encoded, ...user.Option) error); ok {
		r1 = rf(ctx, instID, instRole, firstName, lastName, email, password, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - instID int
//   - instRole institution.Role
//   - firstName string
//   - lastName string
//   - email string
//   - password hash.Encoded
//   - opts ...user.Option
func (_e *MockRepository_Expecter) CreateUser(ctx interface{}, instID interface{}, instRole interface{}, firstName interface{}, lastName interface{}, email interface{}, password interface{}, opts ...interface{}) *MockRepository_CreateUser_Call {
	return &MockRepository_CreateUser_Call{Call: _e.mock.On("CreateUser",
		append([]interface{}{ctx, instID, instRole, firstName, lastName, email, password}, opts...)...)}
}

func (_c *MockRepository_CreateUser_Call) Run(run func(ctx context.Context, instID int, instRole institution.Role, firstName string, lastName string, email string, password hash.Encoded, opts ...user.Option)) *MockRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]user.Option, len(args)-7)
		for i, a := range args[7:] {
			if a != nil {
				variadicArgs[i] = a.(user.Option)
			}
		}
		run(args[0].(context.Context), args[1].(int), args[2].(institution.Role), args[3].(string), args[4].(string), args[5].(string), args[6].(hash.Encoded), variadicArgs...)
	})
	return _c
}

func (_c *MockRepository_CreateUser_Call) Return(_a0 *ent.User, _a1 error) *MockRepository_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_CreateUser_Call) RunAndReturn(run func(context.Context, int, institution.Role, string, string, string, hash.Encoded, ...user.Option) (*ent.User, error)) *MockRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// FindInstitutionInviteLinkWithInstitution provides a mock function with given fields: ctx, code
func (_m *MockRepository) FindInstitutionInviteLinkWithInstitution(ctx context.Context, code string) (*ent.InstitutionInviteLink, error) {
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

// MockRepository_FindInstitutionInviteLinkWithInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitutionInviteLinkWithInstitution'
type MockRepository_FindInstitutionInviteLinkWithInstitution_Call struct {
	*mock.Call
}

// FindInstitutionInviteLinkWithInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - code string
func (_e *MockRepository_Expecter) FindInstitutionInviteLinkWithInstitution(ctx interface{}, code interface{}) *MockRepository_FindInstitutionInviteLinkWithInstitution_Call {
	return &MockRepository_FindInstitutionInviteLinkWithInstitution_Call{Call: _e.mock.On("FindInstitutionInviteLinkWithInstitution", ctx, code)}
}

func (_c *MockRepository_FindInstitutionInviteLinkWithInstitution_Call) Run(run func(ctx context.Context, code string)) *MockRepository_FindInstitutionInviteLinkWithInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindInstitutionInviteLinkWithInstitution_Call) Return(_a0 *ent.InstitutionInviteLink, _a1 error) *MockRepository_FindInstitutionInviteLinkWithInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindInstitutionInviteLinkWithInstitution_Call) RunAndReturn(run func(context.Context, string) (*ent.InstitutionInviteLink, error)) *MockRepository_FindInstitutionInviteLinkWithInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// FindJWTRevocation provides a mock function with given fields: ctx, jti
func (_m *MockRepository) FindJWTRevocation(ctx context.Context, jti string) (*ent.JWTRevocation, error) {
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

// MockRepository_FindJWTRevocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindJWTRevocation'
type MockRepository_FindJWTRevocation_Call struct {
	*mock.Call
}

// FindJWTRevocation is a helper method to define mock.On call
//   - ctx context.Context
//   - jti string
func (_e *MockRepository_Expecter) FindJWTRevocation(ctx interface{}, jti interface{}) *MockRepository_FindJWTRevocation_Call {
	return &MockRepository_FindJWTRevocation_Call{Call: _e.mock.On("FindJWTRevocation", ctx, jti)}
}

func (_c *MockRepository_FindJWTRevocation_Call) Run(run func(ctx context.Context, jti string)) *MockRepository_FindJWTRevocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindJWTRevocation_Call) Return(_a0 *ent.JWTRevocation, _a1 error) *MockRepository_FindJWTRevocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindJWTRevocation_Call) RunAndReturn(run func(context.Context, string) (*ent.JWTRevocation, error)) *MockRepository_FindJWTRevocation_Call {
	_c.Call.Return(run)
	return _c
}

// FindUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
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

// MockRepository_FindUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUserByEmail'
type MockRepository_FindUserByEmail_Call struct {
	*mock.Call
}

// FindUserByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockRepository_Expecter) FindUserByEmail(ctx interface{}, email interface{}) *MockRepository_FindUserByEmail_Call {
	return &MockRepository_FindUserByEmail_Call{Call: _e.mock.On("FindUserByEmail", ctx, email)}
}

func (_c *MockRepository_FindUserByEmail_Call) Run(run func(ctx context.Context, email string)) *MockRepository_FindUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindUserByEmail_Call) Return(_a0 *ent.User, _a1 error) *MockRepository_FindUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindUserByEmail_Call) RunAndReturn(run func(context.Context, string) (*ent.User, error)) *MockRepository_FindUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
