// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	ent "github.com/np-inprove/server/internal/ent"
	entityinstitution "github.com/np-inprove/server/internal/entity/institution"

	mock "github.com/stretchr/testify/mock"
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

// CreateInstitution provides a mock function with given fields: ctx, name, shortName, description
func (_m *MockRepository) CreateInstitution(ctx context.Context, name string, shortName string, description string) (*ent.Institution, error) {
	ret := _m.Called(ctx, name, shortName, description)

	var r0 *ent.Institution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*ent.Institution, error)); ok {
		return rf(ctx, name, shortName, description)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *ent.Institution); ok {
		r0 = rf(ctx, name, shortName, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Institution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, shortName, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_CreateInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateInstitution'
type MockRepository_CreateInstitution_Call struct {
	*mock.Call
}

// CreateInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - shortName string
//   - description string
func (_e *MockRepository_Expecter) CreateInstitution(ctx interface{}, name interface{}, shortName interface{}, description interface{}) *MockRepository_CreateInstitution_Call {
	return &MockRepository_CreateInstitution_Call{Call: _e.mock.On("CreateInstitution", ctx, name, shortName, description)}
}

func (_c *MockRepository_CreateInstitution_Call) Run(run func(ctx context.Context, name string, shortName string, description string)) *MockRepository_CreateInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockRepository_CreateInstitution_Call) Return(_a0 *ent.Institution, _a1 error) *MockRepository_CreateInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_CreateInstitution_Call) RunAndReturn(run func(context.Context, string, string, string) (*ent.Institution, error)) *MockRepository_CreateInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// CreateInviteLink provides a mock function with given fields: ctx, id, code, role
func (_m *MockRepository) CreateInviteLink(ctx context.Context, id int, code string, role entityinstitution.Role) (*ent.InstitutionInviteLink, error) {
	ret := _m.Called(ctx, id, code, role)

	var r0 *ent.InstitutionInviteLink
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, entityinstitution.Role) (*ent.InstitutionInviteLink, error)); ok {
		return rf(ctx, id, code, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, entityinstitution.Role) *ent.InstitutionInviteLink); ok {
		r0 = rf(ctx, id, code, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.InstitutionInviteLink)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, entityinstitution.Role) error); ok {
		r1 = rf(ctx, id, code, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_CreateInviteLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateInviteLink'
type MockRepository_CreateInviteLink_Call struct {
	*mock.Call
}

// CreateInviteLink is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
//   - code string
//   - role entityinstitution.Role
func (_e *MockRepository_Expecter) CreateInviteLink(ctx interface{}, id interface{}, code interface{}, role interface{}) *MockRepository_CreateInviteLink_Call {
	return &MockRepository_CreateInviteLink_Call{Call: _e.mock.On("CreateInviteLink", ctx, id, code, role)}
}

func (_c *MockRepository_CreateInviteLink_Call) Run(run func(ctx context.Context, id int, code string, role entityinstitution.Role)) *MockRepository_CreateInviteLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(entityinstitution.Role))
	})
	return _c
}

func (_c *MockRepository_CreateInviteLink_Call) Return(_a0 *ent.InstitutionInviteLink, _a1 error) *MockRepository_CreateInviteLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_CreateInviteLink_Call) RunAndReturn(run func(context.Context, int, string, entityinstitution.Role) (*ent.InstitutionInviteLink, error)) *MockRepository_CreateInviteLink_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteInstitution provides a mock function with given fields: ctx, id
func (_m *MockRepository) DeleteInstitution(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_DeleteInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteInstitution'
type MockRepository_DeleteInstitution_Call struct {
	*mock.Call
}

// DeleteInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockRepository_Expecter) DeleteInstitution(ctx interface{}, id interface{}) *MockRepository_DeleteInstitution_Call {
	return &MockRepository_DeleteInstitution_Call{Call: _e.mock.On("DeleteInstitution", ctx, id)}
}

func (_c *MockRepository_DeleteInstitution_Call) Run(run func(ctx context.Context, id int)) *MockRepository_DeleteInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepository_DeleteInstitution_Call) Return(_a0 error) *MockRepository_DeleteInstitution_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_DeleteInstitution_Call) RunAndReturn(run func(context.Context, int) error) *MockRepository_DeleteInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteInviteLink provides a mock function with given fields: ctx, id
func (_m *MockRepository) DeleteInviteLink(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_DeleteInviteLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteInviteLink'
type MockRepository_DeleteInviteLink_Call struct {
	*mock.Call
}

// DeleteInviteLink is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockRepository_Expecter) DeleteInviteLink(ctx interface{}, id interface{}) *MockRepository_DeleteInviteLink_Call {
	return &MockRepository_DeleteInviteLink_Call{Call: _e.mock.On("DeleteInviteLink", ctx, id)}
}

func (_c *MockRepository_DeleteInviteLink_Call) Run(run func(ctx context.Context, id int)) *MockRepository_DeleteInviteLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepository_DeleteInviteLink_Call) Return(_a0 error) *MockRepository_DeleteInviteLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_DeleteInviteLink_Call) RunAndReturn(run func(context.Context, int) error) *MockRepository_DeleteInviteLink_Call {
	_c.Call.Return(run)
	return _c
}

// FindInstitution provides a mock function with given fields: ctx, shortName
func (_m *MockRepository) FindInstitution(ctx context.Context, shortName string) (*ent.Institution, error) {
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

// MockRepository_FindInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitution'
type MockRepository_FindInstitution_Call struct {
	*mock.Call
}

// FindInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - shortName string
func (_e *MockRepository_Expecter) FindInstitution(ctx interface{}, shortName interface{}) *MockRepository_FindInstitution_Call {
	return &MockRepository_FindInstitution_Call{Call: _e.mock.On("FindInstitution", ctx, shortName)}
}

func (_c *MockRepository_FindInstitution_Call) Run(run func(ctx context.Context, shortName string)) *MockRepository_FindInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindInstitution_Call) Return(_a0 *ent.Institution, _a1 error) *MockRepository_FindInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindInstitution_Call) RunAndReturn(run func(context.Context, string) (*ent.Institution, error)) *MockRepository_FindInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// FindInstitutionWithInvites provides a mock function with given fields: ctx, shortName
func (_m *MockRepository) FindInstitutionWithInvites(ctx context.Context, shortName string) (*ent.Institution, error) {
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

// MockRepository_FindInstitutionWithInvites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitutionWithInvites'
type MockRepository_FindInstitutionWithInvites_Call struct {
	*mock.Call
}

// FindInstitutionWithInvites is a helper method to define mock.On call
//   - ctx context.Context
//   - shortName string
func (_e *MockRepository_Expecter) FindInstitutionWithInvites(ctx interface{}, shortName interface{}) *MockRepository_FindInstitutionWithInvites_Call {
	return &MockRepository_FindInstitutionWithInvites_Call{Call: _e.mock.On("FindInstitutionWithInvites", ctx, shortName)}
}

func (_c *MockRepository_FindInstitutionWithInvites_Call) Run(run func(ctx context.Context, shortName string)) *MockRepository_FindInstitutionWithInvites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindInstitutionWithInvites_Call) Return(_a0 *ent.Institution, _a1 error) *MockRepository_FindInstitutionWithInvites_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindInstitutionWithInvites_Call) RunAndReturn(run func(context.Context, string) (*ent.Institution, error)) *MockRepository_FindInstitutionWithInvites_Call {
	_c.Call.Return(run)
	return _c
}

// FindInstitutions provides a mock function with given fields: ctx
func (_m *MockRepository) FindInstitutions(ctx context.Context) ([]*ent.Institution, error) {
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

// MockRepository_FindInstitutions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInstitutions'
type MockRepository_FindInstitutions_Call struct {
	*mock.Call
}

// FindInstitutions is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepository_Expecter) FindInstitutions(ctx interface{}) *MockRepository_FindInstitutions_Call {
	return &MockRepository_FindInstitutions_Call{Call: _e.mock.On("FindInstitutions", ctx)}
}

func (_c *MockRepository_FindInstitutions_Call) Run(run func(ctx context.Context)) *MockRepository_FindInstitutions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepository_FindInstitutions_Call) Return(_a0 []*ent.Institution, _a1 error) *MockRepository_FindInstitutions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindInstitutions_Call) RunAndReturn(run func(context.Context) ([]*ent.Institution, error)) *MockRepository_FindInstitutions_Call {
	_c.Call.Return(run)
	return _c
}

// FindInviteLinkWithInstitution provides a mock function with given fields: ctx, code
func (_m *MockRepository) FindInviteLinkWithInstitution(ctx context.Context, code string) (*ent.InstitutionInviteLink, error) {
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

// MockRepository_FindInviteLinkWithInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInviteLinkWithInstitution'
type MockRepository_FindInviteLinkWithInstitution_Call struct {
	*mock.Call
}

// FindInviteLinkWithInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - code string
func (_e *MockRepository_Expecter) FindInviteLinkWithInstitution(ctx interface{}, code interface{}) *MockRepository_FindInviteLinkWithInstitution_Call {
	return &MockRepository_FindInviteLinkWithInstitution_Call{Call: _e.mock.On("FindInviteLinkWithInstitution", ctx, code)}
}

func (_c *MockRepository_FindInviteLinkWithInstitution_Call) Run(run func(ctx context.Context, code string)) *MockRepository_FindInviteLinkWithInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindInviteLinkWithInstitution_Call) Return(_a0 *ent.InstitutionInviteLink, _a1 error) *MockRepository_FindInviteLinkWithInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindInviteLinkWithInstitution_Call) RunAndReturn(run func(context.Context, string) (*ent.InstitutionInviteLink, error)) *MockRepository_FindInviteLinkWithInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// FindInviteLinks provides a mock function with given fields: ctx, id
func (_m *MockRepository) FindInviteLinks(ctx context.Context, id int) ([]*ent.InstitutionInviteLink, error) {
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

// MockRepository_FindInviteLinks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindInviteLinks'
type MockRepository_FindInviteLinks_Call struct {
	*mock.Call
}

// FindInviteLinks is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockRepository_Expecter) FindInviteLinks(ctx interface{}, id interface{}) *MockRepository_FindInviteLinks_Call {
	return &MockRepository_FindInviteLinks_Call{Call: _e.mock.On("FindInviteLinks", ctx, id)}
}

func (_c *MockRepository_FindInviteLinks_Call) Run(run func(ctx context.Context, id int)) *MockRepository_FindInviteLinks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepository_FindInviteLinks_Call) Return(_a0 []*ent.InstitutionInviteLink, _a1 error) *MockRepository_FindInviteLinks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindInviteLinks_Call) RunAndReturn(run func(context.Context, int) ([]*ent.InstitutionInviteLink, error)) *MockRepository_FindInviteLinks_Call {
	_c.Call.Return(run)
	return _c
}

// FindUser provides a mock function with given fields: ctx, principal
func (_m *MockRepository) FindUser(ctx context.Context, principal string) (*ent.User, error) {
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

// MockRepository_FindUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUser'
type MockRepository_FindUser_Call struct {
	*mock.Call
}

// FindUser is a helper method to define mock.On call
//   - ctx context.Context
//   - principal string
func (_e *MockRepository_Expecter) FindUser(ctx interface{}, principal interface{}) *MockRepository_FindUser_Call {
	return &MockRepository_FindUser_Call{Call: _e.mock.On("FindUser", ctx, principal)}
}

func (_c *MockRepository_FindUser_Call) Run(run func(ctx context.Context, principal string)) *MockRepository_FindUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindUser_Call) Return(_a0 *ent.User, _a1 error) *MockRepository_FindUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindUser_Call) RunAndReturn(run func(context.Context, string) (*ent.User, error)) *MockRepository_FindUser_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateInstitution provides a mock function with given fields: ctx, id, name, shortName, description
func (_m *MockRepository) UpdateInstitution(ctx context.Context, id int, name string, shortName string, description string) (*ent.Institution, error) {
	ret := _m.Called(ctx, id, name, shortName, description)

	var r0 *ent.Institution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) (*ent.Institution, error)); ok {
		return rf(ctx, id, name, shortName, description)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) *ent.Institution); ok {
		r0 = rf(ctx, id, name, shortName, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Institution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, string, string) error); ok {
		r1 = rf(ctx, id, name, shortName, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_UpdateInstitution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateInstitution'
type MockRepository_UpdateInstitution_Call struct {
	*mock.Call
}

// UpdateInstitution is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
//   - name string
//   - shortName string
//   - description string
func (_e *MockRepository_Expecter) UpdateInstitution(ctx interface{}, id interface{}, name interface{}, shortName interface{}, description interface{}) *MockRepository_UpdateInstitution_Call {
	return &MockRepository_UpdateInstitution_Call{Call: _e.mock.On("UpdateInstitution", ctx, id, name, shortName, description)}
}

func (_c *MockRepository_UpdateInstitution_Call) Run(run func(ctx context.Context, id int, name string, shortName string, description string)) *MockRepository_UpdateInstitution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockRepository_UpdateInstitution_Call) Return(_a0 *ent.Institution, _a1 error) *MockRepository_UpdateInstitution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_UpdateInstitution_Call) RunAndReturn(run func(context.Context, int, string, string, string) (*ent.Institution, error)) *MockRepository_UpdateInstitution_Call {
	_c.Call.Return(run)
	return _c
}

// WithTx provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) WithTx(_a0 context.Context, _a1 func(context.Context) (interface{}, error)) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) (interface{}, error)) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, func(context.Context) (interface{}, error)) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_WithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTx'
type MockRepository_WithTx_Call struct {
	*mock.Call
}

// WithTx is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 func(context.Context)(interface{} , error)
func (_e *MockRepository_Expecter) WithTx(_a0 interface{}, _a1 interface{}) *MockRepository_WithTx_Call {
	return &MockRepository_WithTx_Call{Call: _e.mock.On("WithTx", _a0, _a1)}
}

func (_c *MockRepository_WithTx_Call) Run(run func(_a0 context.Context, _a1 func(context.Context) (interface{}, error))) *MockRepository_WithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) (interface{}, error)))
	})
	return _c
}

func (_c *MockRepository_WithTx_Call) Return(_a0 interface{}, _a1 error) *MockRepository_WithTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_WithTx_Call) RunAndReturn(run func(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)) *MockRepository_WithTx_Call {
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
