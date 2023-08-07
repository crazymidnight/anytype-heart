// Code generated by mockery v2.26.1. DO NOT EDIT.

package mock_core

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	core "github.com/anyproto/anytype-heart/pkg/lib/core"

	mock "github.com/stretchr/testify/mock"

	threads "github.com/anyproto/anytype-heart/pkg/lib/threads"

	uniquekey "github.com/anyproto/anytype-heart/core/block/uniquekey"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// AccountObjects provides a mock function with given fields:
func (_m *MockService) AccountObjects() threads.DerivedSmartblockIds {
	ret := _m.Called()

	var r0 threads.DerivedSmartblockIds
	if rf, ok := ret.Get(0).(func() threads.DerivedSmartblockIds); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(threads.DerivedSmartblockIds)
	}

	return r0
}

// MockService_AccountObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AccountObjects'
type MockService_AccountObjects_Call struct {
	*mock.Call
}

// AccountObjects is a helper method to define mock.On call
func (_e *MockService_Expecter) AccountObjects() *MockService_AccountObjects_Call {
	return &MockService_AccountObjects_Call{Call: _e.mock.On("AccountObjects")}
}

func (_c *MockService_AccountObjects_Call) Run(run func()) *MockService_AccountObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_AccountObjects_Call) Return(_a0 threads.DerivedSmartblockIds) *MockService_AccountObjects_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_AccountObjects_Call) RunAndReturn(run func() threads.DerivedSmartblockIds) *MockService_AccountObjects_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx
func (_m *MockService) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) Close(ctx interface{}) *MockService_Close_Call {
	return &MockService_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockService_Close_Call) Run(run func(ctx context.Context)) *MockService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_Close_Call) Return(err error) *MockService_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Close_Call) RunAndReturn(run func(context.Context) error) *MockService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveObjectId provides a mock function with given fields: ctx, spaceID, key
func (_m *MockService) DeriveObjectId(ctx context.Context, spaceID string, key uniquekey.UniqueKey) (string, error) {
	ret := _m.Called(ctx, spaceID, key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uniquekey.UniqueKey) (string, error)); ok {
		return rf(ctx, spaceID, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uniquekey.UniqueKey) string); ok {
		r0 = rf(ctx, spaceID, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uniquekey.UniqueKey) error); ok {
		r1 = rf(ctx, spaceID, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_DeriveObjectId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveObjectId'
type MockService_DeriveObjectId_Call struct {
	*mock.Call
}

// DeriveObjectId is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
//   - key uniquekey.UniqueKey
func (_e *MockService_Expecter) DeriveObjectId(ctx interface{}, spaceID interface{}, key interface{}) *MockService_DeriveObjectId_Call {
	return &MockService_DeriveObjectId_Call{Call: _e.mock.On("DeriveObjectId", ctx, spaceID, key)}
}

func (_c *MockService_DeriveObjectId_Call) Run(run func(ctx context.Context, spaceID string, key uniquekey.UniqueKey)) *MockService_DeriveObjectId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uniquekey.UniqueKey))
	})
	return _c
}

func (_c *MockService_DeriveObjectId_Call) Return(_a0 string, _a1 error) *MockService_DeriveObjectId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_DeriveObjectId_Call) RunAndReturn(run func(context.Context, string, uniquekey.UniqueKey) (string, error)) *MockService_DeriveObjectId_Call {
	_c.Call.Return(run)
	return _c
}

// DerivePredefinedObjects provides a mock function with given fields: ctx, spaceID, createTrees
func (_m *MockService) DerivePredefinedObjects(ctx context.Context, spaceID string, createTrees bool) (threads.DerivedSmartblockIds, error) {
	ret := _m.Called(ctx, spaceID, createTrees)

	var r0 threads.DerivedSmartblockIds
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) (threads.DerivedSmartblockIds, error)); ok {
		return rf(ctx, spaceID, createTrees)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) threads.DerivedSmartblockIds); ok {
		r0 = rf(ctx, spaceID, createTrees)
	} else {
		r0 = ret.Get(0).(threads.DerivedSmartblockIds)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, spaceID, createTrees)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_DerivePredefinedObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DerivePredefinedObjects'
type MockService_DerivePredefinedObjects_Call struct {
	*mock.Call
}

// DerivePredefinedObjects is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
//   - createTrees bool
func (_e *MockService_Expecter) DerivePredefinedObjects(ctx interface{}, spaceID interface{}, createTrees interface{}) *MockService_DerivePredefinedObjects_Call {
	return &MockService_DerivePredefinedObjects_Call{Call: _e.mock.On("DerivePredefinedObjects", ctx, spaceID, createTrees)}
}

func (_c *MockService_DerivePredefinedObjects_Call) Run(run func(ctx context.Context, spaceID string, createTrees bool)) *MockService_DerivePredefinedObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockService_DerivePredefinedObjects_Call) Return(predefinedObjectIDs threads.DerivedSmartblockIds, err error) *MockService_DerivePredefinedObjects_Call {
	_c.Call.Return(predefinedObjectIDs, err)
	return _c
}

func (_c *MockService_DerivePredefinedObjects_Call) RunAndReturn(run func(context.Context, string, bool) (threads.DerivedSmartblockIds, error)) *MockService_DerivePredefinedObjects_Call {
	_c.Call.Return(run)
	return _c
}

// EnsurePredefinedBlocks provides a mock function with given fields: ctx, spaceID
func (_m *MockService) EnsurePredefinedBlocks(ctx context.Context, spaceID string) (threads.DerivedSmartblockIds, error) {
	ret := _m.Called(ctx, spaceID)

	var r0 threads.DerivedSmartblockIds
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (threads.DerivedSmartblockIds, error)); ok {
		return rf(ctx, spaceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) threads.DerivedSmartblockIds); ok {
		r0 = rf(ctx, spaceID)
	} else {
		r0 = ret.Get(0).(threads.DerivedSmartblockIds)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_EnsurePredefinedBlocks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnsurePredefinedBlocks'
type MockService_EnsurePredefinedBlocks_Call struct {
	*mock.Call
}

// EnsurePredefinedBlocks is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
func (_e *MockService_Expecter) EnsurePredefinedBlocks(ctx interface{}, spaceID interface{}) *MockService_EnsurePredefinedBlocks_Call {
	return &MockService_EnsurePredefinedBlocks_Call{Call: _e.mock.On("EnsurePredefinedBlocks", ctx, spaceID)}
}

func (_c *MockService_EnsurePredefinedBlocks_Call) Run(run func(ctx context.Context, spaceID string)) *MockService_EnsurePredefinedBlocks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_EnsurePredefinedBlocks_Call) Return(predefinedObjectIDs threads.DerivedSmartblockIds, err error) *MockService_EnsurePredefinedBlocks_Call {
	_c.Call.Return(predefinedObjectIDs, err)
	return _c
}

func (_c *MockService_EnsurePredefinedBlocks_Call) RunAndReturn(run func(context.Context, string) (threads.DerivedSmartblockIds, error)) *MockService_EnsurePredefinedBlocks_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllWorkspaces provides a mock function with given fields:
func (_m *MockService) GetAllWorkspaces() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetAllWorkspaces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllWorkspaces'
type MockService_GetAllWorkspaces_Call struct {
	*mock.Call
}

// GetAllWorkspaces is a helper method to define mock.On call
func (_e *MockService_Expecter) GetAllWorkspaces() *MockService_GetAllWorkspaces_Call {
	return &MockService_GetAllWorkspaces_Call{Call: _e.mock.On("GetAllWorkspaces")}
}

func (_c *MockService_GetAllWorkspaces_Call) Run(run func()) *MockService_GetAllWorkspaces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_GetAllWorkspaces_Call) Return(_a0 []string, _a1 error) *MockService_GetAllWorkspaces_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetAllWorkspaces_Call) RunAndReturn(run func() ([]string, error)) *MockService_GetAllWorkspaces_Call {
	_c.Call.Return(run)
	return _c
}

// GetWorkspaceIdForObject provides a mock function with given fields: spaceID, objectID
func (_m *MockService) GetWorkspaceIdForObject(spaceID string, objectID string) (string, error) {
	ret := _m.Called(spaceID, objectID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(spaceID, objectID)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(spaceID, objectID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(spaceID, objectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetWorkspaceIdForObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWorkspaceIdForObject'
type MockService_GetWorkspaceIdForObject_Call struct {
	*mock.Call
}

// GetWorkspaceIdForObject is a helper method to define mock.On call
//   - spaceID string
//   - objectID string
func (_e *MockService_Expecter) GetWorkspaceIdForObject(spaceID interface{}, objectID interface{}) *MockService_GetWorkspaceIdForObject_Call {
	return &MockService_GetWorkspaceIdForObject_Call{Call: _e.mock.On("GetWorkspaceIdForObject", spaceID, objectID)}
}

func (_c *MockService_GetWorkspaceIdForObject_Call) Run(run func(spaceID string, objectID string)) *MockService_GetWorkspaceIdForObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockService_GetWorkspaceIdForObject_Call) Return(_a0 string, _a1 error) *MockService_GetWorkspaceIdForObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetWorkspaceIdForObject_Call) RunAndReturn(run func(string, string) (string, error)) *MockService_GetWorkspaceIdForObject_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockService) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockService_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockService_Expecter) Init(a interface{}) *MockService_Init_Call {
	return &MockService_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockService_Init_Call) Run(run func(a *app.App)) *MockService_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockService_Init_Call) Return(err error) *MockService_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Init_Call) RunAndReturn(run func(*app.App) error) *MockService_Init_Call {
	_c.Call.Return(run)
	return _c
}

// IsStarted provides a mock function with given fields:
func (_m *MockService) IsStarted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockService_IsStarted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsStarted'
type MockService_IsStarted_Call struct {
	*mock.Call
}

// IsStarted is a helper method to define mock.On call
func (_e *MockService_Expecter) IsStarted() *MockService_IsStarted_Call {
	return &MockService_IsStarted_Call{Call: _e.mock.On("IsStarted")}
}

func (_c *MockService_IsStarted_Call) Run(run func()) *MockService_IsStarted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_IsStarted_Call) Return(_a0 bool) *MockService_IsStarted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_IsStarted_Call) RunAndReturn(run func() bool) *MockService_IsStarted_Call {
	_c.Call.Return(run)
	return _c
}

// LocalProfile provides a mock function with given fields: spaceID
func (_m *MockService) LocalProfile(spaceID string) (core.Profile, error) {
	ret := _m.Called(spaceID)

	var r0 core.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (core.Profile, error)); ok {
		return rf(spaceID)
	}
	if rf, ok := ret.Get(0).(func(string) core.Profile); ok {
		r0 = rf(spaceID)
	} else {
		r0 = ret.Get(0).(core.Profile)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(spaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_LocalProfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LocalProfile'
type MockService_LocalProfile_Call struct {
	*mock.Call
}

// LocalProfile is a helper method to define mock.On call
//   - spaceID string
func (_e *MockService_Expecter) LocalProfile(spaceID interface{}) *MockService_LocalProfile_Call {
	return &MockService_LocalProfile_Call{Call: _e.mock.On("LocalProfile", spaceID)}
}

func (_c *MockService_LocalProfile_Call) Run(run func(spaceID string)) *MockService_LocalProfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_LocalProfile_Call) Return(_a0 core.Profile, _a1 error) *MockService_LocalProfile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_LocalProfile_Call) RunAndReturn(run func(string) (core.Profile, error)) *MockService_LocalProfile_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockService) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockService_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockService_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockService_Expecter) Name() *MockService_Name_Call {
	return &MockService_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockService_Name_Call) Run(run func()) *MockService_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Name_Call) Return(name string) *MockService_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockService_Name_Call) RunAndReturn(run func() string) *MockService_Name_Call {
	_c.Call.Return(run)
	return _c
}

// PredefinedObjects provides a mock function with given fields: spaceID
func (_m *MockService) PredefinedObjects(spaceID string) threads.DerivedSmartblockIds {
	ret := _m.Called(spaceID)

	var r0 threads.DerivedSmartblockIds
	if rf, ok := ret.Get(0).(func(string) threads.DerivedSmartblockIds); ok {
		r0 = rf(spaceID)
	} else {
		r0 = ret.Get(0).(threads.DerivedSmartblockIds)
	}

	return r0
}

// MockService_PredefinedObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PredefinedObjects'
type MockService_PredefinedObjects_Call struct {
	*mock.Call
}

// PredefinedObjects is a helper method to define mock.On call
//   - spaceID string
func (_e *MockService_Expecter) PredefinedObjects(spaceID interface{}) *MockService_PredefinedObjects_Call {
	return &MockService_PredefinedObjects_Call{Call: _e.mock.On("PredefinedObjects", spaceID)}
}

func (_c *MockService_PredefinedObjects_Call) Run(run func(spaceID string)) *MockService_PredefinedObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_PredefinedObjects_Call) Return(_a0 threads.DerivedSmartblockIds) *MockService_PredefinedObjects_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_PredefinedObjects_Call) RunAndReturn(run func(string) threads.DerivedSmartblockIds) *MockService_PredefinedObjects_Call {
	_c.Call.Return(run)
	return _c
}

// ProfileID provides a mock function with given fields: spaceID
func (_m *MockService) ProfileID(spaceID string) string {
	ret := _m.Called(spaceID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(spaceID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockService_ProfileID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProfileID'
type MockService_ProfileID_Call struct {
	*mock.Call
}

// ProfileID is a helper method to define mock.On call
//   - spaceID string
func (_e *MockService_Expecter) ProfileID(spaceID interface{}) *MockService_ProfileID_Call {
	return &MockService_ProfileID_Call{Call: _e.mock.On("ProfileID", spaceID)}
}

func (_c *MockService_ProfileID_Call) Run(run func(spaceID string)) *MockService_ProfileID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_ProfileID_Call) Return(_a0 string) *MockService_ProfileID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_ProfileID_Call) RunAndReturn(run func(string) string) *MockService_ProfileID_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockService) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockService_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) Run(ctx interface{}) *MockService_Run_Call {
	return &MockService_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockService_Run_Call) Run(run func(ctx context.Context)) *MockService_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_Run_Call) Return(err error) *MockService_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Run_Call) RunAndReturn(run func(context.Context) error) *MockService_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockService) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockService_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockService_Expecter) Stop() *MockService_Stop_Call {
	return &MockService_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockService_Stop_Call) Run(run func()) *MockService_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Stop_Call) Return(_a0 error) *MockService_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_Stop_Call) RunAndReturn(run func() error) *MockService_Stop_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockService(t mockConstructorTestingTNewMockService) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
