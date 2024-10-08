// Code generated by mockery. DO NOT EDIT.

package mock_spacesyncstatus

import (
	app "github.com/anyproto/any-sync/app"
	mock "github.com/stretchr/testify/mock"
)

// MockSpaceIdGetter is an autogenerated mock type for the SpaceIdGetter type
type MockSpaceIdGetter struct {
	mock.Mock
}

type MockSpaceIdGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSpaceIdGetter) EXPECT() *MockSpaceIdGetter_Expecter {
	return &MockSpaceIdGetter_Expecter{mock: &_m.Mock}
}

// AllSpaceIds provides a mock function with given fields:
func (_m *MockSpaceIdGetter) AllSpaceIds() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AllSpaceIds")
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

// MockSpaceIdGetter_AllSpaceIds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllSpaceIds'
type MockSpaceIdGetter_AllSpaceIds_Call struct {
	*mock.Call
}

// AllSpaceIds is a helper method to define mock.On call
func (_e *MockSpaceIdGetter_Expecter) AllSpaceIds() *MockSpaceIdGetter_AllSpaceIds_Call {
	return &MockSpaceIdGetter_AllSpaceIds_Call{Call: _e.mock.On("AllSpaceIds")}
}

func (_c *MockSpaceIdGetter_AllSpaceIds_Call) Run(run func()) *MockSpaceIdGetter_AllSpaceIds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceIdGetter_AllSpaceIds_Call) Return(_a0 []string) *MockSpaceIdGetter_AllSpaceIds_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceIdGetter_AllSpaceIds_Call) RunAndReturn(run func() []string) *MockSpaceIdGetter_AllSpaceIds_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockSpaceIdGetter) Init(a *app.App) error {
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

// MockSpaceIdGetter_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockSpaceIdGetter_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockSpaceIdGetter_Expecter) Init(a interface{}) *MockSpaceIdGetter_Init_Call {
	return &MockSpaceIdGetter_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockSpaceIdGetter_Init_Call) Run(run func(a *app.App)) *MockSpaceIdGetter_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockSpaceIdGetter_Init_Call) Return(err error) *MockSpaceIdGetter_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceIdGetter_Init_Call) RunAndReturn(run func(*app.App) error) *MockSpaceIdGetter_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockSpaceIdGetter) Name() string {
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

// MockSpaceIdGetter_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockSpaceIdGetter_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockSpaceIdGetter_Expecter) Name() *MockSpaceIdGetter_Name_Call {
	return &MockSpaceIdGetter_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockSpaceIdGetter_Name_Call) Run(run func()) *MockSpaceIdGetter_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceIdGetter_Name_Call) Return(name string) *MockSpaceIdGetter_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockSpaceIdGetter_Name_Call) RunAndReturn(run func() string) *MockSpaceIdGetter_Name_Call {
	_c.Call.Return(run)
	return _c
}

// TechSpaceId provides a mock function with given fields:
func (_m *MockSpaceIdGetter) TechSpaceId() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TechSpaceId")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSpaceIdGetter_TechSpaceId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TechSpaceId'
type MockSpaceIdGetter_TechSpaceId_Call struct {
	*mock.Call
}

// TechSpaceId is a helper method to define mock.On call
func (_e *MockSpaceIdGetter_Expecter) TechSpaceId() *MockSpaceIdGetter_TechSpaceId_Call {
	return &MockSpaceIdGetter_TechSpaceId_Call{Call: _e.mock.On("TechSpaceId")}
}

func (_c *MockSpaceIdGetter_TechSpaceId_Call) Run(run func()) *MockSpaceIdGetter_TechSpaceId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceIdGetter_TechSpaceId_Call) Return(_a0 string) *MockSpaceIdGetter_TechSpaceId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceIdGetter_TechSpaceId_Call) RunAndReturn(run func() string) *MockSpaceIdGetter_TechSpaceId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSpaceIdGetter creates a new instance of MockSpaceIdGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpaceIdGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpaceIdGetter {
	mock := &MockSpaceIdGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
