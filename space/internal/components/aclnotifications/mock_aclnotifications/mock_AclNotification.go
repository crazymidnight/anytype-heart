// Code generated by mockery. DO NOT EDIT.

package mock_aclnotifications

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	list "github.com/anyproto/any-sync/commonspace/object/acl/list"

	mock "github.com/stretchr/testify/mock"

	spaceinfo "github.com/anyproto/anytype-heart/space/spaceinfo"
)

// MockAclNotification is an autogenerated mock type for the AclNotification type
type MockAclNotification struct {
	mock.Mock
}

type MockAclNotification_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAclNotification) EXPECT() *MockAclNotification_Expecter {
	return &MockAclNotification_Expecter{mock: &_m.Mock}
}

// AddRecords provides a mock function with given fields: acl, permissions, spaceId, accountStatus
func (_m *MockAclNotification) AddRecords(acl list.AclList, permissions list.AclPermissions, spaceId string, accountStatus spaceinfo.AccountStatus) {
	_m.Called(acl, permissions, spaceId, accountStatus)
}

// MockAclNotification_AddRecords_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRecords'
type MockAclNotification_AddRecords_Call struct {
	*mock.Call
}

// AddRecords is a helper method to define mock.On call
//   - acl list.AclList
//   - permissions list.AclPermissions
//   - spaceId string
//   - accountStatus spaceinfo.AccountStatus
func (_e *MockAclNotification_Expecter) AddRecords(acl interface{}, permissions interface{}, spaceId interface{}, accountStatus interface{}) *MockAclNotification_AddRecords_Call {
	return &MockAclNotification_AddRecords_Call{Call: _e.mock.On("AddRecords", acl, permissions, spaceId, accountStatus)}
}

func (_c *MockAclNotification_AddRecords_Call) Run(run func(acl list.AclList, permissions list.AclPermissions, spaceId string, accountStatus spaceinfo.AccountStatus)) *MockAclNotification_AddRecords_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(list.AclList), args[1].(list.AclPermissions), args[2].(string), args[3].(spaceinfo.AccountStatus))
	})
	return _c
}

func (_c *MockAclNotification_AddRecords_Call) Return() *MockAclNotification_AddRecords_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAclNotification_AddRecords_Call) RunAndReturn(run func(list.AclList, list.AclPermissions, string, spaceinfo.AccountStatus)) *MockAclNotification_AddRecords_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx
func (_m *MockAclNotification) Close(ctx context.Context) error {
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

// MockAclNotification_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockAclNotification_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAclNotification_Expecter) Close(ctx interface{}) *MockAclNotification_Close_Call {
	return &MockAclNotification_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockAclNotification_Close_Call) Run(run func(ctx context.Context)) *MockAclNotification_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAclNotification_Close_Call) Return(err error) *MockAclNotification_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockAclNotification_Close_Call) RunAndReturn(run func(context.Context) error) *MockAclNotification_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockAclNotification) Init(a *app.App) error {
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

// MockAclNotification_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockAclNotification_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockAclNotification_Expecter) Init(a interface{}) *MockAclNotification_Init_Call {
	return &MockAclNotification_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockAclNotification_Init_Call) Run(run func(a *app.App)) *MockAclNotification_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockAclNotification_Init_Call) Return(err error) *MockAclNotification_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockAclNotification_Init_Call) RunAndReturn(run func(*app.App) error) *MockAclNotification_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockAclNotification) Name() string {
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

// MockAclNotification_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockAclNotification_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockAclNotification_Expecter) Name() *MockAclNotification_Name_Call {
	return &MockAclNotification_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockAclNotification_Name_Call) Run(run func()) *MockAclNotification_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAclNotification_Name_Call) Return(name string) *MockAclNotification_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockAclNotification_Name_Call) RunAndReturn(run func() string) *MockAclNotification_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockAclNotification) Run(ctx context.Context) error {
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

// MockAclNotification_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockAclNotification_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAclNotification_Expecter) Run(ctx interface{}) *MockAclNotification_Run_Call {
	return &MockAclNotification_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockAclNotification_Run_Call) Run(run func(ctx context.Context)) *MockAclNotification_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAclNotification_Run_Call) Return(err error) *MockAclNotification_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockAclNotification_Run_Call) RunAndReturn(run func(context.Context) error) *MockAclNotification_Run_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAclNotification creates a new instance of MockAclNotification. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAclNotification(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAclNotification {
	mock := &MockAclNotification{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
