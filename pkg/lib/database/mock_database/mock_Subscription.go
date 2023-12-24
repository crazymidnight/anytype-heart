// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_database

import (
	types "github.com/gogo/protobuf/types"
	mock "github.com/stretchr/testify/mock"
)

// MockSubscription is an autogenerated mock type for the Subscription type
type MockSubscription struct {
	mock.Mock
}

type MockSubscription_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSubscription) EXPECT() *MockSubscription_Expecter {
	return &MockSubscription_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockSubscription) Close() {
	_m.Called()
}

// MockSubscription_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSubscription_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockSubscription_Expecter) Close() *MockSubscription_Close_Call {
	return &MockSubscription_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockSubscription_Close_Call) Run(run func()) *MockSubscription_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSubscription_Close_Call) Return() *MockSubscription_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSubscription_Close_Call) RunAndReturn(run func()) *MockSubscription_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Publish provides a mock function with given fields: id, msg
func (_m *MockSubscription) Publish(id string, msg *types.Struct) bool {
	ret := _m.Called(id, msg)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, *types.Struct) bool); ok {
		r0 = rf(id, msg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSubscription_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type MockSubscription_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - id string
//   - msg *types.Struct
func (_e *MockSubscription_Expecter) Publish(id interface{}, msg interface{}) *MockSubscription_Publish_Call {
	return &MockSubscription_Publish_Call{Call: _e.mock.On("Publish", id, msg)}
}

func (_c *MockSubscription_Publish_Call) Run(run func(id string, msg *types.Struct)) *MockSubscription_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*types.Struct))
	})
	return _c
}

func (_c *MockSubscription_Publish_Call) Return(_a0 bool) *MockSubscription_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSubscription_Publish_Call) RunAndReturn(run func(string, *types.Struct) bool) *MockSubscription_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// RecordChan provides a mock function with given fields:
func (_m *MockSubscription) RecordChan() chan *types.Struct {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RecordChan")
	}

	var r0 chan *types.Struct
	if rf, ok := ret.Get(0).(func() chan *types.Struct); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.Struct)
		}
	}

	return r0
}

// MockSubscription_RecordChan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecordChan'
type MockSubscription_RecordChan_Call struct {
	*mock.Call
}

// RecordChan is a helper method to define mock.On call
func (_e *MockSubscription_Expecter) RecordChan() *MockSubscription_RecordChan_Call {
	return &MockSubscription_RecordChan_Call{Call: _e.mock.On("RecordChan")}
}

func (_c *MockSubscription_RecordChan_Call) Run(run func()) *MockSubscription_RecordChan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSubscription_RecordChan_Call) Return(_a0 chan *types.Struct) *MockSubscription_RecordChan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSubscription_RecordChan_Call) RunAndReturn(run func() chan *types.Struct) *MockSubscription_RecordChan_Call {
	_c.Call.Return(run)
	return _c
}

// Subscribe provides a mock function with given fields: ids
func (_m *MockSubscription) Subscribe(ids []string) []string {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for Subscribe")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockSubscription_Subscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscribe'
type MockSubscription_Subscribe_Call struct {
	*mock.Call
}

// Subscribe is a helper method to define mock.On call
//   - ids []string
func (_e *MockSubscription_Expecter) Subscribe(ids interface{}) *MockSubscription_Subscribe_Call {
	return &MockSubscription_Subscribe_Call{Call: _e.mock.On("Subscribe", ids)}
}

func (_c *MockSubscription_Subscribe_Call) Run(run func(ids []string)) *MockSubscription_Subscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockSubscription_Subscribe_Call) Return(added []string) *MockSubscription_Subscribe_Call {
	_c.Call.Return(added)
	return _c
}

func (_c *MockSubscription_Subscribe_Call) RunAndReturn(run func([]string) []string) *MockSubscription_Subscribe_Call {
	_c.Call.Return(run)
	return _c
}

// Subscriptions provides a mock function with given fields:
func (_m *MockSubscription) Subscriptions() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Subscriptions")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockSubscription_Subscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscriptions'
type MockSubscription_Subscriptions_Call struct {
	*mock.Call
}

// Subscriptions is a helper method to define mock.On call
func (_e *MockSubscription_Expecter) Subscriptions() *MockSubscription_Subscriptions_Call {
	return &MockSubscription_Subscriptions_Call{Call: _e.mock.On("Subscriptions")}
}

func (_c *MockSubscription_Subscriptions_Call) Run(run func()) *MockSubscription_Subscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSubscription_Subscriptions_Call) Return(_a0 []string) *MockSubscription_Subscriptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSubscription_Subscriptions_Call) RunAndReturn(run func() []string) *MockSubscription_Subscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSubscription creates a new instance of MockSubscription. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSubscription(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSubscription {
	mock := &MockSubscription{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
