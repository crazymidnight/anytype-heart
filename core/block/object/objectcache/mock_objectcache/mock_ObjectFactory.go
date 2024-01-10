// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_objectcache

import (
	mock "github.com/stretchr/testify/mock"

	smartblock "github.com/anyproto/anytype-heart/core/block/editor/smartblock"
)

// MockObjectFactory is an autogenerated mock type for the ObjectFactory type
type MockObjectFactory struct {
	mock.Mock
}

type MockObjectFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *MockObjectFactory) EXPECT() *MockObjectFactory_Expecter {
	return &MockObjectFactory_Expecter{mock: &_m.Mock}
}

// InitObject provides a mock function with given fields: space, id, initCtx
func (_m *MockObjectFactory) InitObject(space smartblock.Space, id string, initCtx *smartblock.InitContext) (smartblock.SmartBlock, error) {
	ret := _m.Called(space, id, initCtx)

	if len(ret) == 0 {
		panic("no return value specified for InitObject")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(smartblock.Space, string, *smartblock.InitContext) (smartblock.SmartBlock, error)); ok {
		return rf(space, id, initCtx)
	}
	if rf, ok := ret.Get(0).(func(smartblock.Space, string, *smartblock.InitContext) smartblock.SmartBlock); ok {
		r0 = rf(space, id, initCtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(smartblock.Space, string, *smartblock.InitContext) error); ok {
		r1 = rf(space, id, initCtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockObjectFactory_InitObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitObject'
type MockObjectFactory_InitObject_Call struct {
	*mock.Call
}

// InitObject is a helper method to define mock.On call
//   - space smartblock.Space
//   - id string
//   - initCtx *smartblock.InitContext
func (_e *MockObjectFactory_Expecter) InitObject(space interface{}, id interface{}, initCtx interface{}) *MockObjectFactory_InitObject_Call {
	return &MockObjectFactory_InitObject_Call{Call: _e.mock.On("InitObject", space, id, initCtx)}
}

func (_c *MockObjectFactory_InitObject_Call) Run(run func(space smartblock.Space, id string, initCtx *smartblock.InitContext)) *MockObjectFactory_InitObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(smartblock.Space), args[1].(string), args[2].(*smartblock.InitContext))
	})
	return _c
}

func (_c *MockObjectFactory_InitObject_Call) Return(sb smartblock.SmartBlock, err error) *MockObjectFactory_InitObject_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockObjectFactory_InitObject_Call) RunAndReturn(run func(smartblock.Space, string, *smartblock.InitContext) (smartblock.SmartBlock, error)) *MockObjectFactory_InitObject_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockObjectFactory creates a new instance of MockObjectFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockObjectFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockObjectFactory {
	mock := &MockObjectFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
