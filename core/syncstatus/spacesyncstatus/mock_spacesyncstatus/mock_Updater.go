// Code generated by mockery. DO NOT EDIT.

package mock_spacesyncstatus

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	domain "github.com/anyproto/anytype-heart/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockUpdater is an autogenerated mock type for the Updater type
type MockUpdater struct {
	mock.Mock
}

type MockUpdater_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdater) EXPECT() *MockUpdater_Expecter {
	return &MockUpdater_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockUpdater) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUpdater_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockUpdater_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUpdater_Expecter) Close(ctx interface{}) *MockUpdater_Close_Call {
	return &MockUpdater_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockUpdater_Close_Call) Run(run func(ctx context.Context)) *MockUpdater_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUpdater_Close_Call) Return(err error) *MockUpdater_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUpdater_Close_Call) RunAndReturn(run func(context.Context) error) *MockUpdater_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockUpdater) Init(a *app.App) error {
	ret := _m.Called(a)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUpdater_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockUpdater_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockUpdater_Expecter) Init(a interface{}) *MockUpdater_Init_Call {
	return &MockUpdater_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockUpdater_Init_Call) Run(run func(a *app.App)) *MockUpdater_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockUpdater_Init_Call) Return(err error) *MockUpdater_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUpdater_Init_Call) RunAndReturn(run func(*app.App) error) *MockUpdater_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockUpdater) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockUpdater_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockUpdater_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockUpdater_Expecter) Name() *MockUpdater_Name_Call {
	return &MockUpdater_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockUpdater_Name_Call) Run(run func()) *MockUpdater_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUpdater_Name_Call) Return(name string) *MockUpdater_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockUpdater_Name_Call) RunAndReturn(run func() string) *MockUpdater_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Reload provides a mock function with given fields: ctx
func (_m *MockUpdater) Reload(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Reload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUpdater_Reload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reload'
type MockUpdater_Reload_Call struct {
	*mock.Call
}

// Reload is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUpdater_Expecter) Reload(ctx interface{}) *MockUpdater_Reload_Call {
	return &MockUpdater_Reload_Call{Call: _e.mock.On("Reload", ctx)}
}

func (_c *MockUpdater_Reload_Call) Run(run func(ctx context.Context)) *MockUpdater_Reload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUpdater_Reload_Call) Return(err error) *MockUpdater_Reload_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUpdater_Reload_Call) RunAndReturn(run func(context.Context) error) *MockUpdater_Reload_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockUpdater) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUpdater_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockUpdater_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUpdater_Expecter) Run(ctx interface{}) *MockUpdater_Run_Call {
	return &MockUpdater_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockUpdater_Run_Call) Run(run func(ctx context.Context)) *MockUpdater_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUpdater_Run_Call) Return(err error) *MockUpdater_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUpdater_Run_Call) RunAndReturn(run func(context.Context) error) *MockUpdater_Run_Call {
	_c.Call.Return(run)
	return _c
}

// SendUpdate provides a mock function with given fields: spaceSync
func (_m *MockUpdater) SendUpdate(spaceSync *domain.SpaceSync) {
	_m.Called(spaceSync)
}

// MockUpdater_SendUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendUpdate'
type MockUpdater_SendUpdate_Call struct {
	*mock.Call
}

// SendUpdate is a helper method to define mock.On call
//   - spaceSync *domain.SpaceSync
func (_e *MockUpdater_Expecter) SendUpdate(spaceSync interface{}) *MockUpdater_SendUpdate_Call {
	return &MockUpdater_SendUpdate_Call{Call: _e.mock.On("SendUpdate", spaceSync)}
}

func (_c *MockUpdater_SendUpdate_Call) Run(run func(spaceSync *domain.SpaceSync)) *MockUpdater_SendUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.SpaceSync))
	})
	return _c
}

func (_c *MockUpdater_SendUpdate_Call) Return() *MockUpdater_SendUpdate_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUpdater_SendUpdate_Call) RunAndReturn(run func(*domain.SpaceSync)) *MockUpdater_SendUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdater creates a new instance of MockUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdater(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdater {
	mock := &MockUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
