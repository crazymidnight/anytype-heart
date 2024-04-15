// Code generated by mockery. DO NOT EDIT.

package mock_cache

import (
	context "context"

	domain "github.com/anyproto/anytype-heart/core/domain"
	mock "github.com/stretchr/testify/mock"

	smartblock "github.com/anyproto/anytype-heart/core/block/editor/smartblock"
)

// MockObjectGetter is an autogenerated mock type for the ObjectGetter type
type MockObjectGetter struct {
	mock.Mock
}

type MockObjectGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockObjectGetter) EXPECT() *MockObjectGetter_Expecter {
	return &MockObjectGetter_Expecter{mock: &_m.Mock}
}

// GetObject provides a mock function with given fields: ctx, objectID
func (_m *MockObjectGetter) GetObject(ctx context.Context, objectID string) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, objectID)

	if len(ret) == 0 {
		panic("no return value specified for GetObject")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, objectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) smartblock.SmartBlock); ok {
		r0 = rf(ctx, objectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, objectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockObjectGetter_GetObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObject'
type MockObjectGetter_GetObject_Call struct {
	*mock.Call
}

// GetObject is a helper method to define mock.On call
//   - ctx context.Context
//   - objectID string
func (_e *MockObjectGetter_Expecter) GetObject(ctx interface{}, objectID interface{}) *MockObjectGetter_GetObject_Call {
	return &MockObjectGetter_GetObject_Call{Call: _e.mock.On("GetObject", ctx, objectID)}
}

func (_c *MockObjectGetter_GetObject_Call) Run(run func(ctx context.Context, objectID string)) *MockObjectGetter_GetObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockObjectGetter_GetObject_Call) Return(sb smartblock.SmartBlock, err error) *MockObjectGetter_GetObject_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockObjectGetter_GetObject_Call) RunAndReturn(run func(context.Context, string) (smartblock.SmartBlock, error)) *MockObjectGetter_GetObject_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectByFullID provides a mock function with given fields: ctx, id
func (_m *MockObjectGetter) GetObjectByFullID(ctx context.Context, id domain.FullID) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetObjectByFullID")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.FullID) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.FullID) smartblock.SmartBlock); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.FullID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockObjectGetter_GetObjectByFullID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectByFullID'
type MockObjectGetter_GetObjectByFullID_Call struct {
	*mock.Call
}

// GetObjectByFullID is a helper method to define mock.On call
//   - ctx context.Context
//   - id domain.FullID
func (_e *MockObjectGetter_Expecter) GetObjectByFullID(ctx interface{}, id interface{}) *MockObjectGetter_GetObjectByFullID_Call {
	return &MockObjectGetter_GetObjectByFullID_Call{Call: _e.mock.On("GetObjectByFullID", ctx, id)}
}

func (_c *MockObjectGetter_GetObjectByFullID_Call) Run(run func(ctx context.Context, id domain.FullID)) *MockObjectGetter_GetObjectByFullID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.FullID))
	})
	return _c
}

func (_c *MockObjectGetter_GetObjectByFullID_Call) Return(sb smartblock.SmartBlock, err error) *MockObjectGetter_GetObjectByFullID_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockObjectGetter_GetObjectByFullID_Call) RunAndReturn(run func(context.Context, domain.FullID) (smartblock.SmartBlock, error)) *MockObjectGetter_GetObjectByFullID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockObjectGetter creates a new instance of MockObjectGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockObjectGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockObjectGetter {
	mock := &MockObjectGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
