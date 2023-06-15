// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	ent "github.com/np-inprove/server/internal/ent"
	hash "github.com/np-inprove/server/internal/hash"

	mock "github.com/stretchr/testify/mock"

	time "time"

	user "github.com/np-inprove/server/internal/entity/user"
)

// MockWriter is an autogenerated mock type for the Writer type
type MockWriter struct {
	mock.Mock
}

type MockWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWriter) EXPECT() *MockWriter_Expecter {
	return &MockWriter_Expecter{mock: &_m.Mock}
}

// CreateJWTRevocation provides a mock function with given fields: ctx, jti, expiry
func (_m *MockWriter) CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (*ent.JWTRevocation, error) {
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

// MockWriter_CreateJWTRevocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJWTRevocation'
type MockWriter_CreateJWTRevocation_Call struct {
	*mock.Call
}

// CreateJWTRevocation is a helper method to define mock.On call
//   - ctx context.Context
//   - jti string
//   - expiry time.Time
func (_e *MockWriter_Expecter) CreateJWTRevocation(ctx interface{}, jti interface{}, expiry interface{}) *MockWriter_CreateJWTRevocation_Call {
	return &MockWriter_CreateJWTRevocation_Call{Call: _e.mock.On("CreateJWTRevocation", ctx, jti, expiry)}
}

func (_c *MockWriter_CreateJWTRevocation_Call) Run(run func(ctx context.Context, jti string, expiry time.Time)) *MockWriter_CreateJWTRevocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Time))
	})
	return _c
}

func (_c *MockWriter_CreateJWTRevocation_Call) Return(_a0 *ent.JWTRevocation, _a1 error) *MockWriter_CreateJWTRevocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWriter_CreateJWTRevocation_Call) RunAndReturn(run func(context.Context, string, time.Time) (*ent.JWTRevocation, error)) *MockWriter_CreateJWTRevocation_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, instID, firstName, lastName, email, password, opts
func (_m *MockWriter) CreateUser(ctx context.Context, instID int, firstName string, lastName string, email string, password hash.Encoded, opts ...user.Option) (*ent.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, instID, firstName, lastName, email, password)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string, hash.Encoded, ...user.Option) (*ent.User, error)); ok {
		return rf(ctx, instID, firstName, lastName, email, password, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string, hash.Encoded, ...user.Option) *ent.User); ok {
		r0 = rf(ctx, instID, firstName, lastName, email, password, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, string, string, hash.Encoded, ...user.Option) error); ok {
		r1 = rf(ctx, instID, firstName, lastName, email, password, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWriter_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockWriter_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - instID int
//   - firstName string
//   - lastName string
//   - email string
//   - password hash.Encoded
//   - opts ...user.Option
func (_e *MockWriter_Expecter) CreateUser(ctx interface{}, instID interface{}, firstName interface{}, lastName interface{}, email interface{}, password interface{}, opts ...interface{}) *MockWriter_CreateUser_Call {
	return &MockWriter_CreateUser_Call{Call: _e.mock.On("CreateUser",
		append([]interface{}{ctx, instID, firstName, lastName, email, password}, opts...)...)}
}

func (_c *MockWriter_CreateUser_Call) Run(run func(ctx context.Context, instID int, firstName string, lastName string, email string, password hash.Encoded, opts ...user.Option)) *MockWriter_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]user.Option, len(args)-6)
		for i, a := range args[6:] {
			if a != nil {
				variadicArgs[i] = a.(user.Option)
			}
		}
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(string), args[4].(string), args[5].(hash.Encoded), variadicArgs...)
	})
	return _c
}

func (_c *MockWriter_CreateUser_Call) Return(_a0 *ent.User, _a1 error) *MockWriter_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWriter_CreateUser_Call) RunAndReturn(run func(context.Context, int, string, string, string, hash.Encoded, ...user.Option) (*ent.User, error)) *MockWriter_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWriter creates a new instance of MockWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockWriter(t mockConstructorTestingTNewMockWriter) *MockWriter {
	mock := &MockWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
