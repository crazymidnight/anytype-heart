// Code generated by mockery v2.35.0. DO NOT EDIT.

package mock_restriction

import (
	app "github.com/anyproto/any-sync/app"
	mock "github.com/stretchr/testify/mock"

	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"

	restriction "github.com/anyproto/anytype-heart/core/block/restriction"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// CheckRestrictions provides a mock function with given fields: spaceID, id, cr
func (_m *MockService) CheckRestrictions(spaceID string, id string, cr ...model.RestrictionsObjectRestriction) error {
	_va := make([]interface{}, len(cr))
	for _i := range cr {
		_va[_i] = cr[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, spaceID, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...model.RestrictionsObjectRestriction) error); ok {
		r0 = rf(spaceID, id, cr...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_CheckRestrictions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckRestrictions'
type MockService_CheckRestrictions_Call struct {
	*mock.Call
}

// CheckRestrictions is a helper method to define mock.On call
//   - spaceID string
//   - id string
//   - cr ...model.RestrictionsObjectRestriction
func (_e *MockService_Expecter) CheckRestrictions(spaceID interface{}, id interface{}, cr ...interface{}) *MockService_CheckRestrictions_Call {
	return &MockService_CheckRestrictions_Call{Call: _e.mock.On("CheckRestrictions",
		append([]interface{}{spaceID, id}, cr...)...)}
}

func (_c *MockService_CheckRestrictions_Call) Run(run func(spaceID string, id string, cr ...model.RestrictionsObjectRestriction)) *MockService_CheckRestrictions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]model.RestrictionsObjectRestriction, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(model.RestrictionsObjectRestriction)
			}
		}
		run(args[0].(string), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockService_CheckRestrictions_Call) Return(_a0 error) *MockService_CheckRestrictions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_CheckRestrictions_Call) RunAndReturn(run func(string, string, ...model.RestrictionsObjectRestriction) error) *MockService_CheckRestrictions_Call {
	_c.Call.Return(run)
	return _c
}

// GetRestrictions provides a mock function with given fields: _a0
func (_m *MockService) GetRestrictions(_a0 restriction.RestrictionHolder) restriction.Restrictions {
	ret := _m.Called(_a0)

	var r0 restriction.Restrictions
	if rf, ok := ret.Get(0).(func(restriction.RestrictionHolder) restriction.Restrictions); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(restriction.Restrictions)
	}

	return r0
}

// MockService_GetRestrictions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRestrictions'
type MockService_GetRestrictions_Call struct {
	*mock.Call
}

// GetRestrictions is a helper method to define mock.On call
//   - _a0 restriction.RestrictionHolder
func (_e *MockService_Expecter) GetRestrictions(_a0 interface{}) *MockService_GetRestrictions_Call {
	return &MockService_GetRestrictions_Call{Call: _e.mock.On("GetRestrictions", _a0)}
}

func (_c *MockService_GetRestrictions_Call) Run(run func(_a0 restriction.RestrictionHolder)) *MockService_GetRestrictions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(restriction.RestrictionHolder))
	})
	return _c
}

func (_c *MockService_GetRestrictions_Call) Return(_a0 restriction.Restrictions) *MockService_GetRestrictions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_GetRestrictions_Call) RunAndReturn(run func(restriction.RestrictionHolder) restriction.Restrictions) *MockService_GetRestrictions_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockService) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockService_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockService_Expecter) Init(a interface{}) *MockService_Init_Call {
	return &MockService_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockService_Init_Call) Run(run func(a *app.App)) *MockService_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockService_Init_Call) Return(err error) *MockService_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Init_Call) RunAndReturn(run func(*app.App) error) *MockService_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockService) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockService_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockService_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockService_Expecter) Name() *MockService_Name_Call {
	return &MockService_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockService_Name_Call) Run(run func()) *MockService_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Name_Call) Return(name string) *MockService_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockService_Name_Call) RunAndReturn(run func() string) *MockService_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
