// Code generated by mockery v2.35.2. DO NOT EDIT.

package mock_system_object

import (
	context "context"

	app "github.com/anyproto/any-sync/app"
	mock "github.com/stretchr/testify/mock"

	domain "github.com/anyproto/anytype-heart/core/domain"
)

// Mockderiver is an autogenerated mock type for the deriver type
type Mockderiver struct {
	mock.Mock
}

type Mockderiver_Expecter struct {
	mock *mock.Mock
}

func (_m *Mockderiver) EXPECT() *Mockderiver_Expecter {
	return &Mockderiver_Expecter{mock: &_m.Mock}
}

// DeriveObjectID provides a mock function with given fields: ctx, spaceID, uniqueKey
func (_m *Mockderiver) DeriveObjectID(ctx context.Context, spaceID string, uniqueKey domain.UniqueKey) (string, error) {
	ret := _m.Called(ctx, spaceID, uniqueKey)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UniqueKey) (string, error)); ok {
		return rf(ctx, spaceID, uniqueKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UniqueKey) string); ok {
		r0 = rf(ctx, spaceID, uniqueKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.UniqueKey) error); ok {
		r1 = rf(ctx, spaceID, uniqueKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mockderiver_DeriveObjectID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveObjectID'
type Mockderiver_DeriveObjectID_Call struct {
	*mock.Call
}

// DeriveObjectID is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
//   - uniqueKey domain.UniqueKey
func (_e *Mockderiver_Expecter) DeriveObjectID(ctx interface{}, spaceID interface{}, uniqueKey interface{}) *Mockderiver_DeriveObjectID_Call {
	return &Mockderiver_DeriveObjectID_Call{Call: _e.mock.On("DeriveObjectID", ctx, spaceID, uniqueKey)}
}

func (_c *Mockderiver_DeriveObjectID_Call) Run(run func(ctx context.Context, spaceID string, uniqueKey domain.UniqueKey)) *Mockderiver_DeriveObjectID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.UniqueKey))
	})
	return _c
}

func (_c *Mockderiver_DeriveObjectID_Call) Return(id string, err error) *Mockderiver_DeriveObjectID_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *Mockderiver_DeriveObjectID_Call) RunAndReturn(run func(context.Context, string, domain.UniqueKey) (string, error)) *Mockderiver_DeriveObjectID_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *Mockderiver) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mockderiver_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type Mockderiver_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *Mockderiver_Expecter) Init(a interface{}) *Mockderiver_Init_Call {
	return &Mockderiver_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *Mockderiver_Init_Call) Run(run func(a *app.App)) *Mockderiver_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *Mockderiver_Init_Call) Return(err error) *Mockderiver_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Mockderiver_Init_Call) RunAndReturn(run func(*app.App) error) *Mockderiver_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Mockderiver) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Mockderiver_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Mockderiver_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Mockderiver_Expecter) Name() *Mockderiver_Name_Call {
	return &Mockderiver_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Mockderiver_Name_Call) Run(run func()) *Mockderiver_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Mockderiver_Name_Call) Return(name string) *Mockderiver_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *Mockderiver_Name_Call) RunAndReturn(run func() string) *Mockderiver_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockderiver creates a new instance of Mockderiver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockderiver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mockderiver {
	mock := &Mockderiver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
