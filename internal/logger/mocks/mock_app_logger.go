// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	zapcore "go.uber.org/zap/zapcore"
)

// MockAppLogger is an autogenerated mock type for the AppLogger type
type MockAppLogger struct {
	mock.Mock
}

type MockAppLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAppLogger) EXPECT() *MockAppLogger_Expecter {
	return &MockAppLogger_Expecter{mock: &_m.Mock}
}

// DPanic provides a mock function with given fields: msg, fields
func (_m *MockAppLogger) DPanic(msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAppLogger_DPanic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DPanic'
type MockAppLogger_DPanic_Call struct {
	*mock.Call
}

// DPanic is a helper method to define mock.On call
//   - msg string
//   - fields ...zapcore.Field
func (_e *MockAppLogger_Expecter) DPanic(msg interface{}, fields ...interface{}) *MockAppLogger_DPanic_Call {
	return &MockAppLogger_DPanic_Call{Call: _e.mock.On("DPanic",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockAppLogger_DPanic_Call) Run(run func(msg string, fields ...zapcore.Field)) *MockAppLogger_DPanic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAppLogger_DPanic_Call) Return() *MockAppLogger_DPanic_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAppLogger_DPanic_Call) RunAndReturn(run func(string, ...zapcore.Field)) *MockAppLogger_DPanic_Call {
	_c.Call.Return(run)
	return _c
}

// Debug provides a mock function with given fields: msg, fields
func (_m *MockAppLogger) Debug(msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAppLogger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type MockAppLogger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//   - msg string
//   - fields ...zapcore.Field
func (_e *MockAppLogger_Expecter) Debug(msg interface{}, fields ...interface{}) *MockAppLogger_Debug_Call {
	return &MockAppLogger_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockAppLogger_Debug_Call) Run(run func(msg string, fields ...zapcore.Field)) *MockAppLogger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAppLogger_Debug_Call) Return() *MockAppLogger_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAppLogger_Debug_Call) RunAndReturn(run func(string, ...zapcore.Field)) *MockAppLogger_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: msg, fields
func (_m *MockAppLogger) Error(msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAppLogger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type MockAppLogger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - msg string
//   - fields ...zapcore.Field
func (_e *MockAppLogger_Expecter) Error(msg interface{}, fields ...interface{}) *MockAppLogger_Error_Call {
	return &MockAppLogger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockAppLogger_Error_Call) Run(run func(msg string, fields ...zapcore.Field)) *MockAppLogger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAppLogger_Error_Call) Return() *MockAppLogger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAppLogger_Error_Call) RunAndReturn(run func(string, ...zapcore.Field)) *MockAppLogger_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Fatal provides a mock function with given fields: msg, fields
func (_m *MockAppLogger) Fatal(msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAppLogger_Fatal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatal'
type MockAppLogger_Fatal_Call struct {
	*mock.Call
}

// Fatal is a helper method to define mock.On call
//   - msg string
//   - fields ...zapcore.Field
func (_e *MockAppLogger_Expecter) Fatal(msg interface{}, fields ...interface{}) *MockAppLogger_Fatal_Call {
	return &MockAppLogger_Fatal_Call{Call: _e.mock.On("Fatal",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockAppLogger_Fatal_Call) Run(run func(msg string, fields ...zapcore.Field)) *MockAppLogger_Fatal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAppLogger_Fatal_Call) Return() *MockAppLogger_Fatal_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAppLogger_Fatal_Call) RunAndReturn(run func(string, ...zapcore.Field)) *MockAppLogger_Fatal_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: msg, fields
func (_m *MockAppLogger) Info(msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAppLogger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type MockAppLogger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - msg string
//   - fields ...zapcore.Field
func (_e *MockAppLogger_Expecter) Info(msg interface{}, fields ...interface{}) *MockAppLogger_Info_Call {
	return &MockAppLogger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockAppLogger_Info_Call) Run(run func(msg string, fields ...zapcore.Field)) *MockAppLogger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAppLogger_Info_Call) Return() *MockAppLogger_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAppLogger_Info_Call) RunAndReturn(run func(string, ...zapcore.Field)) *MockAppLogger_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Panic provides a mock function with given fields: msg, fields
func (_m *MockAppLogger) Panic(msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAppLogger_Panic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Panic'
type MockAppLogger_Panic_Call struct {
	*mock.Call
}

// Panic is a helper method to define mock.On call
//   - msg string
//   - fields ...zapcore.Field
func (_e *MockAppLogger_Expecter) Panic(msg interface{}, fields ...interface{}) *MockAppLogger_Panic_Call {
	return &MockAppLogger_Panic_Call{Call: _e.mock.On("Panic",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockAppLogger_Panic_Call) Run(run func(msg string, fields ...zapcore.Field)) *MockAppLogger_Panic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAppLogger_Panic_Call) Return() *MockAppLogger_Panic_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAppLogger_Panic_Call) RunAndReturn(run func(string, ...zapcore.Field)) *MockAppLogger_Panic_Call {
	_c.Call.Return(run)
	return _c
}

// Warn provides a mock function with given fields: msg, fields
func (_m *MockAppLogger) Warn(msg string, fields ...zapcore.Field) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAppLogger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type MockAppLogger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//   - msg string
//   - fields ...zapcore.Field
func (_e *MockAppLogger_Expecter) Warn(msg interface{}, fields ...interface{}) *MockAppLogger_Warn_Call {
	return &MockAppLogger_Warn_Call{Call: _e.mock.On("Warn",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockAppLogger_Warn_Call) Run(run func(msg string, fields ...zapcore.Field)) *MockAppLogger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAppLogger_Warn_Call) Return() *MockAppLogger_Warn_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAppLogger_Warn_Call) RunAndReturn(run func(string, ...zapcore.Field)) *MockAppLogger_Warn_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAppLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAppLogger creates a new instance of MockAppLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAppLogger(t mockConstructorTestingTNewMockAppLogger) *MockAppLogger {
	mock := &MockAppLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
