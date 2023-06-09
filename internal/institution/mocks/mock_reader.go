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

// FindInstitution provides a mock function with given fields: ctx, shortName
func (_m *MockReader) FindInstitution(ctx context.Context, shortName string) (*ent.Institution, error) {
	ret := _m.Called(ctx, shortName)

	var r0 *ent.Institution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.Institution, error)); ok {
		return rf(ctx, shortName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.Institution); ok {
		r0 = rf(ctx, shortName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Institution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitution'
type MockReader_FindInstitution_Call struct {
	*mock.Call
}

// FindInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - shortName string
func (_e *MockReader_Expecter) FindInstitution(ctx interface{}, shortName interface{}) *MockReader_FindInstitution_Call {
	return &MockReader_FindInstitution_Call{Call: _e.mock.On("FindInstitution", ctx, shortName)}
}

func (_c *MockReader_FindInstitution_Call) Run(run func(ctx context.Context, shortName string)) *MockReader_FindInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReader_FindInstitution_Call) Return(_a0 *ent.Institution, _a1 error) *MockReader_FindInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindInstitution_Call) RunAndReturn(run func(context.Context, string) (*ent.Institution, error)) *MockReader_FindInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// FindInstitutionWithInvites provides a mock function with given fields: ctx, shortName
func (_m *MockReader) FindInstitutionWithInvites(ctx context.Context, shortName string) (*ent.Institution, error) {
	ret := _m.Called(ctx, shortName)

	var r0 *ent.Institution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.Institution, error)); ok {
		return rf(ctx, shortName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.Institution); ok {
		r0 = rf(ctx, shortName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Institution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindInstitutionWithInvites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitutionWithInvites'
type MockReader_FindInstitutionWithInvites_Call struct {
	*mock.Call
}

// FindInstitutionWithInvites is a helper method to define mock.On call
//   - ctx context.Context
//   - shortName string
func (_e *MockReader_Expecter) FindInstitutionWithInvites(ctx interface{}, shortName interface{}) *MockReader_FindInstitutionWithInvites_Call {
	return &MockReader_FindInstitutionWithInvites_Call{Call: _e.mock.On("FindInstitutionWithInvites", ctx, shortName)}
}

func (_c *MockReader_FindInstitutionWithInvites_Call) Run(run func(ctx context.Context, shortName string)) *MockReader_FindInstitutionWithInvites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReader_FindInstitutionWithInvites_Call) Return(_a0 *ent.Institution, _a1 error) *MockReader_FindInstitutionWithInvites_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindInstitutionWithInvites_Call) RunAndReturn(run func(context.Context, string) (*ent.Institution, error)) *MockReader_FindInstitutionWithInvites_Call {
	_c.Call.Return(run)
	return _c
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

// FindInviteLinkWithInstitution provides a mock function with given fields: ctx, code
func (_m *MockReader) FindInviteLinkWithInstitution(ctx context.Context, code string) (*ent.InstitutionInviteLink, error) {
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

// MockReader_FindInviteLinkWithInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInviteLinkWithInstitution'
type MockReader_FindInviteLinkWithInstitution_Call struct {
	*mock.Call
}

// FindInviteLinkWithInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - code string
func (_e *MockReader_Expecter) FindInviteLinkWithInstitution(ctx interface{}, code interface{}) *MockReader_FindInviteLinkWithInstitution_Call {
	return &MockReader_FindInviteLinkWithInstitution_Call{Call: _e.mock.On("FindInviteLinkWithInstitution", ctx, code)}
}

func (_c *MockReader_FindInviteLinkWithInstitution_Call) Run(run func(ctx context.Context, code string)) *MockReader_FindInviteLinkWithInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReader_FindInviteLinkWithInstitution_Call) Return(_a0 *ent.InstitutionInviteLink, _a1 error) *MockReader_FindInviteLinkWithInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindInviteLinkWithInstitution_Call) RunAndReturn(run func(context.Context, string) (*ent.InstitutionInviteLink, error)) *MockReader_FindInviteLinkWithInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// FindInviteLinks provides a mock function with given fields: ctx, id
func (_m *MockReader) FindInviteLinks(ctx context.Context, id int) ([]*ent.InstitutionInviteLink, error) {
	ret := _m.Called(ctx, id)

	var r0 []*ent.InstitutionInviteLink
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]*ent.InstitutionInviteLink, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []*ent.InstitutionInviteLink); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.InstitutionInviteLink)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindInviteLinks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInviteLinks'
type MockReader_FindInviteLinks_Call struct {
	*mock.Call
}

// FindInviteLinks is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockReader_Expecter) FindInviteLinks(ctx interface{}, id interface{}) *MockReader_FindInviteLinks_Call {
	return &MockReader_FindInviteLinks_Call{Call: _e.mock.On("FindInviteLinks", ctx, id)}
}

func (_c *MockReader_FindInviteLinks_Call) Run(run func(ctx context.Context, id int)) *MockReader_FindInviteLinks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockReader_FindInviteLinks_Call) Return(_a0 []*ent.InstitutionInviteLink, _a1 error) *MockReader_FindInviteLinks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindInviteLinks_Call) RunAndReturn(run func(context.Context, int) ([]*ent.InstitutionInviteLink, error)) *MockReader_FindInviteLinks_Call {
	_c.Call.Return(run)
	return _c
}

// FindUser provides a mock function with given fields: ctx, principal
func (_m *MockReader) FindUser(ctx context.Context, principal string) (*ent.User, error) {
	ret := _m.Called(ctx, principal)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.User, error)); ok {
		return rf(ctx, principal)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.User); ok {
		r0 = rf(ctx, principal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, principal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_FindUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUser'
type MockReader_FindUser_Call struct {
	*mock.Call
}

// FindUser is a helper method to define mock.On call
//   - ctx context.Context
//   - principal string
func (_e *MockReader_Expecter) FindUser(ctx interface{}, principal interface{}) *MockReader_FindUser_Call {
	return &MockReader_FindUser_Call{Call: _e.mock.On("FindUser", ctx, principal)}
}

func (_c *MockReader_FindUser_Call) Run(run func(ctx context.Context, principal string)) *MockReader_FindUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReader_FindUser_Call) Return(_a0 *ent.User, _a1 error) *MockReader_FindUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_FindUser_Call) RunAndReturn(run func(context.Context, string) (*ent.User, error)) *MockReader_FindUser_Call {
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
