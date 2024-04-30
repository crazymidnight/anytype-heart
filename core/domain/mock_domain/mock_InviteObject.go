// Code generated by mockery. DO NOT EDIT.

package mock_domain

import mock "github.com/stretchr/testify/mock"

// MockInviteObject is an autogenerated mock type for the InviteObject type
type MockInviteObject struct {
	mock.Mock
}

type MockInviteObject_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInviteObject) EXPECT() *MockInviteObject_Expecter {
	return &MockInviteObject_Expecter{mock: &_m.Mock}
}

// GetExistingInviteInfo provides a mock function with given fields:
func (_m *MockInviteObject) GetExistingInviteInfo() (string, string) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExistingInviteInfo")
	}

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func() (string, string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MockInviteObject_GetExistingInviteInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExistingInviteInfo'
type MockInviteObject_GetExistingInviteInfo_Call struct {
	*mock.Call
}

// GetExistingInviteInfo is a helper method to define mock.On call
func (_e *MockInviteObject_Expecter) GetExistingInviteInfo() *MockInviteObject_GetExistingInviteInfo_Call {
	return &MockInviteObject_GetExistingInviteInfo_Call{Call: _e.mock.On("GetExistingInviteInfo")}
}

func (_c *MockInviteObject_GetExistingInviteInfo_Call) Run(run func()) *MockInviteObject_GetExistingInviteInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInviteObject_GetExistingInviteInfo_Call) Return(fileCid string, fileKey string) *MockInviteObject_GetExistingInviteInfo_Call {
	_c.Call.Return(fileCid, fileKey)
	return _c
}

func (_c *MockInviteObject_GetExistingInviteInfo_Call) RunAndReturn(run func() (string, string)) *MockInviteObject_GetExistingInviteInfo_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveExistingInviteInfo provides a mock function with given fields:
func (_m *MockInviteObject) RemoveExistingInviteInfo() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoveExistingInviteInfo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInviteObject_RemoveExistingInviteInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveExistingInviteInfo'
type MockInviteObject_RemoveExistingInviteInfo_Call struct {
	*mock.Call
}

// RemoveExistingInviteInfo is a helper method to define mock.On call
func (_e *MockInviteObject_Expecter) RemoveExistingInviteInfo() *MockInviteObject_RemoveExistingInviteInfo_Call {
	return &MockInviteObject_RemoveExistingInviteInfo_Call{Call: _e.mock.On("RemoveExistingInviteInfo")}
}

func (_c *MockInviteObject_RemoveExistingInviteInfo_Call) Run(run func()) *MockInviteObject_RemoveExistingInviteInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInviteObject_RemoveExistingInviteInfo_Call) Return(fileCid string, err error) *MockInviteObject_RemoveExistingInviteInfo_Call {
	_c.Call.Return(fileCid, err)
	return _c
}

func (_c *MockInviteObject_RemoveExistingInviteInfo_Call) RunAndReturn(run func() (string, error)) *MockInviteObject_RemoveExistingInviteInfo_Call {
	_c.Call.Return(run)
	return _c
}

// SetInviteFileInfo provides a mock function with given fields: fileCid, fileKey
func (_m *MockInviteObject) SetInviteFileInfo(fileCid string, fileKey string) error {
	ret := _m.Called(fileCid, fileKey)

	if len(ret) == 0 {
		panic("no return value specified for SetInviteFileInfo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(fileCid, fileKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInviteObject_SetInviteFileInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetInviteFileInfo'
type MockInviteObject_SetInviteFileInfo_Call struct {
	*mock.Call
}

// SetInviteFileInfo is a helper method to define mock.On call
//   - fileCid string
//   - fileKey string
func (_e *MockInviteObject_Expecter) SetInviteFileInfo(fileCid interface{}, fileKey interface{}) *MockInviteObject_SetInviteFileInfo_Call {
	return &MockInviteObject_SetInviteFileInfo_Call{Call: _e.mock.On("SetInviteFileInfo", fileCid, fileKey)}
}

func (_c *MockInviteObject_SetInviteFileInfo_Call) Run(run func(fileCid string, fileKey string)) *MockInviteObject_SetInviteFileInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockInviteObject_SetInviteFileInfo_Call) Return(err error) *MockInviteObject_SetInviteFileInfo_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockInviteObject_SetInviteFileInfo_Call) RunAndReturn(run func(string, string) error) *MockInviteObject_SetInviteFileInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInviteObject creates a new instance of MockInviteObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInviteObject(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInviteObject {
	mock := &MockInviteObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
