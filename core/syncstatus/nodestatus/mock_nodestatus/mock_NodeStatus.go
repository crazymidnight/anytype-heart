// Code generated by mockery. DO NOT EDIT.

package mock_nodestatus

import (
	app "github.com/anyproto/any-sync/app"
	mock "github.com/stretchr/testify/mock"

	nodestatus "github.com/anyproto/anytype-heart/core/syncstatus/nodestatus"
)

// MockNodeStatus is an autogenerated mock type for the NodeStatus type
type MockNodeStatus struct {
	mock.Mock
}

type MockNodeStatus_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNodeStatus) EXPECT() *MockNodeStatus_Expecter {
	return &MockNodeStatus_Expecter{mock: &_m.Mock}
}

// GetNodeStatus provides a mock function with given fields: spaceId
func (_m *MockNodeStatus) GetNodeStatus(spaceId string) nodestatus.ConnectionStatus {
	ret := _m.Called(spaceId)

	if len(ret) == 0 {
		panic("no return value specified for GetNodeStatus")
	}

	var r0 nodestatus.ConnectionStatus
	if rf, ok := ret.Get(0).(func(string) nodestatus.ConnectionStatus); ok {
		r0 = rf(spaceId)
	} else {
		r0 = ret.Get(0).(nodestatus.ConnectionStatus)
	}

	return r0
}

// MockNodeStatus_GetNodeStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNodeStatus'
type MockNodeStatus_GetNodeStatus_Call struct {
	*mock.Call
}

// GetNodeStatus is a helper method to define mock.On call
//   - spaceId string
func (_e *MockNodeStatus_Expecter) GetNodeStatus(spaceId interface{}) *MockNodeStatus_GetNodeStatus_Call {
	return &MockNodeStatus_GetNodeStatus_Call{Call: _e.mock.On("GetNodeStatus", spaceId)}
}

func (_c *MockNodeStatus_GetNodeStatus_Call) Run(run func(spaceId string)) *MockNodeStatus_GetNodeStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockNodeStatus_GetNodeStatus_Call) Return(_a0 nodestatus.ConnectionStatus) *MockNodeStatus_GetNodeStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNodeStatus_GetNodeStatus_Call) RunAndReturn(run func(string) nodestatus.ConnectionStatus) *MockNodeStatus_GetNodeStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockNodeStatus) Init(a *app.App) error {
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

// MockNodeStatus_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockNodeStatus_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockNodeStatus_Expecter) Init(a interface{}) *MockNodeStatus_Init_Call {
	return &MockNodeStatus_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockNodeStatus_Init_Call) Run(run func(a *app.App)) *MockNodeStatus_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockNodeStatus_Init_Call) Return(err error) *MockNodeStatus_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockNodeStatus_Init_Call) RunAndReturn(run func(*app.App) error) *MockNodeStatus_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockNodeStatus) Name() string {
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

// MockNodeStatus_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockNodeStatus_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockNodeStatus_Expecter) Name() *MockNodeStatus_Name_Call {
	return &MockNodeStatus_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockNodeStatus_Name_Call) Run(run func()) *MockNodeStatus_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNodeStatus_Name_Call) Return(name string) *MockNodeStatus_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockNodeStatus_Name_Call) RunAndReturn(run func() string) *MockNodeStatus_Name_Call {
	_c.Call.Return(run)
	return _c
}

// SetNodesStatus provides a mock function with given fields: spaceId, status
func (_m *MockNodeStatus) SetNodesStatus(spaceId string, status nodestatus.ConnectionStatus) {
	_m.Called(spaceId, status)
}

// MockNodeStatus_SetNodesStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNodesStatus'
type MockNodeStatus_SetNodesStatus_Call struct {
	*mock.Call
}

// SetNodesStatus is a helper method to define mock.On call
//   - spaceId string
//   - status nodestatus.ConnectionStatus
func (_e *MockNodeStatus_Expecter) SetNodesStatus(spaceId interface{}, status interface{}) *MockNodeStatus_SetNodesStatus_Call {
	return &MockNodeStatus_SetNodesStatus_Call{Call: _e.mock.On("SetNodesStatus", spaceId, status)}
}

func (_c *MockNodeStatus_SetNodesStatus_Call) Run(run func(spaceId string, status nodestatus.ConnectionStatus)) *MockNodeStatus_SetNodesStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(nodestatus.ConnectionStatus))
	})
	return _c
}

func (_c *MockNodeStatus_SetNodesStatus_Call) Return() *MockNodeStatus_SetNodesStatus_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNodeStatus_SetNodesStatus_Call) RunAndReturn(run func(string, nodestatus.ConnectionStatus)) *MockNodeStatus_SetNodesStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNodeStatus creates a new instance of MockNodeStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNodeStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNodeStatus {
	mock := &MockNodeStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
