// Code generated by mockery v2.35.2. DO NOT EDIT.

package mock_spacecore

import (
	context "context"

	app "github.com/anyproto/any-sync/app"
	streampool "github.com/anyproto/any-sync/net/streampool"
	mock "github.com/stretchr/testify/mock"

	spacecore "github.com/anyproto/anytype-heart/space/spacecore"
)

// MockSpaceCoreService is an autogenerated mock type for the SpaceCoreService type
type MockSpaceCoreService struct {
	mock.Mock
}

type MockSpaceCoreService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSpaceCoreService) EXPECT() *MockSpaceCoreService_Expecter {
	return &MockSpaceCoreService_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockSpaceCoreService) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSpaceCoreService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceCoreService_Expecter) Close(ctx interface{}) *MockSpaceCoreService_Close_Call {
	return &MockSpaceCoreService_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockSpaceCoreService_Close_Call) Run(run func(ctx context.Context)) *MockSpaceCoreService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceCoreService_Close_Call) Return(err error) *MockSpaceCoreService_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_Close_Call) RunAndReturn(run func(context.Context) error) *MockSpaceCoreService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, replicationKey
func (_m *MockSpaceCoreService) Create(ctx context.Context, replicationKey uint64) (*spacecore.AnySpace, error) {
	ret := _m.Called(ctx, replicationKey)

	var r0 *spacecore.AnySpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*spacecore.AnySpace, error)); ok {
		return rf(ctx, replicationKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *spacecore.AnySpace); ok {
		r0 = rf(ctx, replicationKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spacecore.AnySpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, replicationKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockSpaceCoreService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - replicationKey uint64
func (_e *MockSpaceCoreService_Expecter) Create(ctx interface{}, replicationKey interface{}) *MockSpaceCoreService_Create_Call {
	return &MockSpaceCoreService_Create_Call{Call: _e.mock.On("Create", ctx, replicationKey)}
}

func (_c *MockSpaceCoreService_Create_Call) Run(run func(ctx context.Context, replicationKey uint64)) *MockSpaceCoreService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockSpaceCoreService_Create_Call) Return(_a0 *spacecore.AnySpace, _a1 error) *MockSpaceCoreService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceCoreService_Create_Call) RunAndReturn(run func(context.Context, uint64) (*spacecore.AnySpace, error)) *MockSpaceCoreService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, spaceID
func (_m *MockSpaceCoreService) Delete(ctx context.Context, spaceID string) (spacecore.NetworkStatus, error) {
	ret := _m.Called(ctx, spaceID)

	var r0 spacecore.NetworkStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (spacecore.NetworkStatus, error)); ok {
		return rf(ctx, spaceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) spacecore.NetworkStatus); ok {
		r0 = rf(ctx, spaceID)
	} else {
		r0 = ret.Get(0).(spacecore.NetworkStatus)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockSpaceCoreService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
func (_e *MockSpaceCoreService_Expecter) Delete(ctx interface{}, spaceID interface{}) *MockSpaceCoreService_Delete_Call {
	return &MockSpaceCoreService_Delete_Call{Call: _e.mock.On("Delete", ctx, spaceID)}
}

func (_c *MockSpaceCoreService_Delete_Call) Run(run func(ctx context.Context, spaceID string)) *MockSpaceCoreService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_Delete_Call) Return(payload spacecore.NetworkStatus, err error) *MockSpaceCoreService_Delete_Call {
	_c.Call.Return(payload, err)
	return _c
}

func (_c *MockSpaceCoreService_Delete_Call) RunAndReturn(run func(context.Context, string) (spacecore.NetworkStatus, error)) *MockSpaceCoreService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Derive provides a mock function with given fields: ctx, spaceType
func (_m *MockSpaceCoreService) Derive(ctx context.Context, spaceType string) (*spacecore.AnySpace, error) {
	ret := _m.Called(ctx, spaceType)

	var r0 *spacecore.AnySpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*spacecore.AnySpace, error)); ok {
		return rf(ctx, spaceType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *spacecore.AnySpace); ok {
		r0 = rf(ctx, spaceType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spacecore.AnySpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Derive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Derive'
type MockSpaceCoreService_Derive_Call struct {
	*mock.Call
}

// Derive is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceType string
func (_e *MockSpaceCoreService_Expecter) Derive(ctx interface{}, spaceType interface{}) *MockSpaceCoreService_Derive_Call {
	return &MockSpaceCoreService_Derive_Call{Call: _e.mock.On("Derive", ctx, spaceType)}
}

func (_c *MockSpaceCoreService_Derive_Call) Run(run func(ctx context.Context, spaceType string)) *MockSpaceCoreService_Derive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_Derive_Call) Return(space *spacecore.AnySpace, err error) *MockSpaceCoreService_Derive_Call {
	_c.Call.Return(space, err)
	return _c
}

func (_c *MockSpaceCoreService_Derive_Call) RunAndReturn(run func(context.Context, string) (*spacecore.AnySpace, error)) *MockSpaceCoreService_Derive_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveID provides a mock function with given fields: ctx, spaceType
func (_m *MockSpaceCoreService) DeriveID(ctx context.Context, spaceType string) (string, error) {
	ret := _m.Called(ctx, spaceType)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, spaceType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, spaceType)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_DeriveID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveID'
type MockSpaceCoreService_DeriveID_Call struct {
	*mock.Call
}

// DeriveID is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceType string
func (_e *MockSpaceCoreService_Expecter) DeriveID(ctx interface{}, spaceType interface{}) *MockSpaceCoreService_DeriveID_Call {
	return &MockSpaceCoreService_DeriveID_Call{Call: _e.mock.On("DeriveID", ctx, spaceType)}
}

func (_c *MockSpaceCoreService_DeriveID_Call) Run(run func(ctx context.Context, spaceType string)) *MockSpaceCoreService_DeriveID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_DeriveID_Call) Return(id string, err error) *MockSpaceCoreService_DeriveID_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockSpaceCoreService_DeriveID_Call) RunAndReturn(run func(context.Context, string) (string, error)) *MockSpaceCoreService_DeriveID_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockSpaceCoreService) Get(ctx context.Context, id string) (*spacecore.AnySpace, error) {
	ret := _m.Called(ctx, id)

	var r0 *spacecore.AnySpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*spacecore.AnySpace, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *spacecore.AnySpace); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spacecore.AnySpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockSpaceCoreService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockSpaceCoreService_Expecter) Get(ctx interface{}, id interface{}) *MockSpaceCoreService_Get_Call {
	return &MockSpaceCoreService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockSpaceCoreService_Get_Call) Run(run func(ctx context.Context, id string)) *MockSpaceCoreService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_Get_Call) Return(_a0 *spacecore.AnySpace, _a1 error) *MockSpaceCoreService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceCoreService_Get_Call) RunAndReturn(run func(context.Context, string) (*spacecore.AnySpace, error)) *MockSpaceCoreService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockSpaceCoreService) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockSpaceCoreService_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockSpaceCoreService_Expecter) Init(a interface{}) *MockSpaceCoreService_Init_Call {
	return &MockSpaceCoreService_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockSpaceCoreService_Init_Call) Run(run func(a *app.App)) *MockSpaceCoreService_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockSpaceCoreService_Init_Call) Return(err error) *MockSpaceCoreService_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_Init_Call) RunAndReturn(run func(*app.App) error) *MockSpaceCoreService_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockSpaceCoreService) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSpaceCoreService_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockSpaceCoreService_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockSpaceCoreService_Expecter) Name() *MockSpaceCoreService_Name_Call {
	return &MockSpaceCoreService_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockSpaceCoreService_Name_Call) Run(run func()) *MockSpaceCoreService_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceCoreService_Name_Call) Return(name string) *MockSpaceCoreService_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockSpaceCoreService_Name_Call) RunAndReturn(run func() string) *MockSpaceCoreService_Name_Call {
	_c.Call.Return(run)
	return _c
}

// RevertDeletion provides a mock function with given fields: ctx, spaceID
func (_m *MockSpaceCoreService) RevertDeletion(ctx context.Context, spaceID string) error {
	ret := _m.Called(ctx, spaceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, spaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_RevertDeletion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevertDeletion'
type MockSpaceCoreService_RevertDeletion_Call struct {
	*mock.Call
}

// RevertDeletion is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
func (_e *MockSpaceCoreService_Expecter) RevertDeletion(ctx interface{}, spaceID interface{}) *MockSpaceCoreService_RevertDeletion_Call {
	return &MockSpaceCoreService_RevertDeletion_Call{Call: _e.mock.On("RevertDeletion", ctx, spaceID)}
}

func (_c *MockSpaceCoreService_RevertDeletion_Call) Run(run func(ctx context.Context, spaceID string)) *MockSpaceCoreService_RevertDeletion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_RevertDeletion_Call) Return(err error) *MockSpaceCoreService_RevertDeletion_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_RevertDeletion_Call) RunAndReturn(run func(context.Context, string) error) *MockSpaceCoreService_RevertDeletion_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockSpaceCoreService) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockSpaceCoreService_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceCoreService_Expecter) Run(ctx interface{}) *MockSpaceCoreService_Run_Call {
	return &MockSpaceCoreService_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockSpaceCoreService_Run_Call) Run(run func(ctx context.Context)) *MockSpaceCoreService_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceCoreService_Run_Call) Return(err error) *MockSpaceCoreService_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_Run_Call) RunAndReturn(run func(context.Context) error) *MockSpaceCoreService_Run_Call {
	_c.Call.Return(run)
	return _c
}

// StreamPool provides a mock function with given fields:
func (_m *MockSpaceCoreService) StreamPool() streampool.StreamPool {
	ret := _m.Called()

	var r0 streampool.StreamPool
	if rf, ok := ret.Get(0).(func() streampool.StreamPool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(streampool.StreamPool)
		}
	}

	return r0
}

// MockSpaceCoreService_StreamPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamPool'
type MockSpaceCoreService_StreamPool_Call struct {
	*mock.Call
}

// StreamPool is a helper method to define mock.On call
func (_e *MockSpaceCoreService_Expecter) StreamPool() *MockSpaceCoreService_StreamPool_Call {
	return &MockSpaceCoreService_StreamPool_Call{Call: _e.mock.On("StreamPool")}
}

func (_c *MockSpaceCoreService_StreamPool_Call) Run(run func()) *MockSpaceCoreService_StreamPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceCoreService_StreamPool_Call) Return(_a0 streampool.StreamPool) *MockSpaceCoreService_StreamPool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceCoreService_StreamPool_Call) RunAndReturn(run func() streampool.StreamPool) *MockSpaceCoreService_StreamPool_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSpaceCoreService creates a new instance of MockSpaceCoreService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpaceCoreService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpaceCoreService {
	mock := &MockSpaceCoreService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
