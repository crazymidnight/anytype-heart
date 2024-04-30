// Code generated by mockery. DO NOT EDIT.

package mock_invitemigrator

import (
	app "github.com/anyproto/any-sync/app"
	clientspace "github.com/anyproto/anytype-heart/space/clientspace"

	mock "github.com/stretchr/testify/mock"
)

// MockInviteMigrator is an autogenerated mock type for the InviteMigrator type
type MockInviteMigrator struct {
	mock.Mock
}

type MockInviteMigrator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInviteMigrator) EXPECT() *MockInviteMigrator_Expecter {
	return &MockInviteMigrator_Expecter{mock: &_m.Mock}
}

// Init provides a mock function with given fields: a
func (_m *MockInviteMigrator) Init(a *app.App) error {
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

// MockInviteMigrator_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockInviteMigrator_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockInviteMigrator_Expecter) Init(a interface{}) *MockInviteMigrator_Init_Call {
	return &MockInviteMigrator_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockInviteMigrator_Init_Call) Run(run func(a *app.App)) *MockInviteMigrator_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockInviteMigrator_Init_Call) Return(err error) *MockInviteMigrator_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockInviteMigrator_Init_Call) RunAndReturn(run func(*app.App) error) *MockInviteMigrator_Init_Call {
	_c.Call.Return(run)
	return _c
}

// MigrateExistingInvites provides a mock function with given fields: space
func (_m *MockInviteMigrator) MigrateExistingInvites(space clientspace.Space) error {
	ret := _m.Called(space)

	if len(ret) == 0 {
		panic("no return value specified for MigrateExistingInvites")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(clientspace.Space) error); ok {
		r0 = rf(space)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInviteMigrator_MigrateExistingInvites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MigrateExistingInvites'
type MockInviteMigrator_MigrateExistingInvites_Call struct {
	*mock.Call
}

// MigrateExistingInvites is a helper method to define mock.On call
//   - space clientspace.Space
func (_e *MockInviteMigrator_Expecter) MigrateExistingInvites(space interface{}) *MockInviteMigrator_MigrateExistingInvites_Call {
	return &MockInviteMigrator_MigrateExistingInvites_Call{Call: _e.mock.On("MigrateExistingInvites", space)}
}

func (_c *MockInviteMigrator_MigrateExistingInvites_Call) Run(run func(space clientspace.Space)) *MockInviteMigrator_MigrateExistingInvites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(clientspace.Space))
	})
	return _c
}

func (_c *MockInviteMigrator_MigrateExistingInvites_Call) Return(_a0 error) *MockInviteMigrator_MigrateExistingInvites_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInviteMigrator_MigrateExistingInvites_Call) RunAndReturn(run func(clientspace.Space) error) *MockInviteMigrator_MigrateExistingInvites_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockInviteMigrator) Name() string {
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

// MockInviteMigrator_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockInviteMigrator_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockInviteMigrator_Expecter) Name() *MockInviteMigrator_Name_Call {
	return &MockInviteMigrator_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockInviteMigrator_Name_Call) Run(run func()) *MockInviteMigrator_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInviteMigrator_Name_Call) Return(name string) *MockInviteMigrator_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockInviteMigrator_Name_Call) RunAndReturn(run func() string) *MockInviteMigrator_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInviteMigrator creates a new instance of MockInviteMigrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInviteMigrator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInviteMigrator {
	mock := &MockInviteMigrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
