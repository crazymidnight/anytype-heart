// Code generated by mockery v2.35.0. DO NOT EDIT.

package mock_space

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	mock "github.com/stretchr/testify/mock"

	types "github.com/gogo/protobuf/types"
)

// MockbundledObjectsInstaller is an autogenerated mock type for the bundledObjectsInstaller type
type MockbundledObjectsInstaller struct {
	mock.Mock
}

type MockbundledObjectsInstaller_Expecter struct {
	mock *mock.Mock
}

func (_m *MockbundledObjectsInstaller) EXPECT() *MockbundledObjectsInstaller_Expecter {
	return &MockbundledObjectsInstaller_Expecter{mock: &_m.Mock}
}

// Init provides a mock function with given fields: a
func (_m *MockbundledObjectsInstaller) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockbundledObjectsInstaller_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockbundledObjectsInstaller_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockbundledObjectsInstaller_Expecter) Init(a interface{}) *MockbundledObjectsInstaller_Init_Call {
	return &MockbundledObjectsInstaller_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockbundledObjectsInstaller_Init_Call) Run(run func(a *app.App)) *MockbundledObjectsInstaller_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockbundledObjectsInstaller_Init_Call) Return(err error) *MockbundledObjectsInstaller_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockbundledObjectsInstaller_Init_Call) RunAndReturn(run func(*app.App) error) *MockbundledObjectsInstaller_Init_Call {
	_c.Call.Return(run)
	return _c
}

// InstallBundledObjects provides a mock function with given fields: ctx, spaceID, ids
func (_m *MockbundledObjectsInstaller) InstallBundledObjects(ctx context.Context, spaceID string, ids []string) ([]string, []*types.Struct, error) {
	ret := _m.Called(ctx, spaceID, ids)

	var r0 []string
	var r1 []*types.Struct
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]string, []*types.Struct, error)); ok {
		return rf(ctx, spaceID, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []string); ok {
		r0 = rf(ctx, spaceID, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) []*types.Struct); ok {
		r1 = rf(ctx, spaceID, ids)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.Struct)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, []string) error); ok {
		r2 = rf(ctx, spaceID, ids)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockbundledObjectsInstaller_InstallBundledObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InstallBundledObjects'
type MockbundledObjectsInstaller_InstallBundledObjects_Call struct {
	*mock.Call
}

// InstallBundledObjects is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
//   - ids []string
func (_e *MockbundledObjectsInstaller_Expecter) InstallBundledObjects(ctx interface{}, spaceID interface{}, ids interface{}) *MockbundledObjectsInstaller_InstallBundledObjects_Call {
	return &MockbundledObjectsInstaller_InstallBundledObjects_Call{Call: _e.mock.On("InstallBundledObjects", ctx, spaceID, ids)}
}

func (_c *MockbundledObjectsInstaller_InstallBundledObjects_Call) Run(run func(ctx context.Context, spaceID string, ids []string)) *MockbundledObjectsInstaller_InstallBundledObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockbundledObjectsInstaller_InstallBundledObjects_Call) Return(_a0 []string, _a1 []*types.Struct, _a2 error) *MockbundledObjectsInstaller_InstallBundledObjects_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockbundledObjectsInstaller_InstallBundledObjects_Call) RunAndReturn(run func(context.Context, string, []string) ([]string, []*types.Struct, error)) *MockbundledObjectsInstaller_InstallBundledObjects_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockbundledObjectsInstaller) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockbundledObjectsInstaller_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockbundledObjectsInstaller_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockbundledObjectsInstaller_Expecter) Name() *MockbundledObjectsInstaller_Name_Call {
	return &MockbundledObjectsInstaller_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockbundledObjectsInstaller_Name_Call) Run(run func()) *MockbundledObjectsInstaller_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockbundledObjectsInstaller_Name_Call) Return(name string) *MockbundledObjectsInstaller_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockbundledObjectsInstaller_Name_Call) RunAndReturn(run func() string) *MockbundledObjectsInstaller_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockbundledObjectsInstaller creates a new instance of MockbundledObjectsInstaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockbundledObjectsInstaller(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockbundledObjectsInstaller {
	mock := &MockbundledObjectsInstaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
