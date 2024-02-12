// Code generated by mockery. DO NOT EDIT.

package mock_space

import (
	context "context"

	app "github.com/anyproto/any-sync/app"
	clientspace "github.com/anyproto/anytype-heart/space/clientspace"

	crypto "github.com/anyproto/any-sync/util/crypto"

	mock "github.com/stretchr/testify/mock"
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

// AccountMetadataPayload provides a mock function with given fields:
func (_m *MockService) AccountMetadataPayload() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AccountMetadataPayload")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockService_AccountMetadataPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AccountMetadataPayload'
type MockService_AccountMetadataPayload_Call struct {
	*mock.Call
}

// AccountMetadataPayload is a helper method to define mock.On call
func (_e *MockService_Expecter) AccountMetadataPayload() *MockService_AccountMetadataPayload_Call {
	return &MockService_AccountMetadataPayload_Call{Call: _e.mock.On("AccountMetadataPayload")}
}

func (_c *MockService_AccountMetadataPayload_Call) Run(run func()) *MockService_AccountMetadataPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_AccountMetadataPayload_Call) Return(_a0 []byte) *MockService_AccountMetadataPayload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_AccountMetadataPayload_Call) RunAndReturn(run func() []byte) *MockService_AccountMetadataPayload_Call {
	_c.Call.Return(run)
	return _c
}

// AccountMetadataSymKey provides a mock function with given fields:
func (_m *MockService) AccountMetadataSymKey() crypto.SymKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AccountMetadataSymKey")
	}

	var r0 crypto.SymKey
	if rf, ok := ret.Get(0).(func() crypto.SymKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.SymKey)
		}
	}

	return r0
}

// MockService_AccountMetadataSymKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AccountMetadataSymKey'
type MockService_AccountMetadataSymKey_Call struct {
	*mock.Call
}

// AccountMetadataSymKey is a helper method to define mock.On call
func (_e *MockService_Expecter) AccountMetadataSymKey() *MockService_AccountMetadataSymKey_Call {
	return &MockService_AccountMetadataSymKey_Call{Call: _e.mock.On("AccountMetadataSymKey")}
}

func (_c *MockService_AccountMetadataSymKey_Call) Run(run func()) *MockService_AccountMetadataSymKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_AccountMetadataSymKey_Call) Return(_a0 crypto.SymKey) *MockService_AccountMetadataSymKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_AccountMetadataSymKey_Call) RunAndReturn(run func() crypto.SymKey) *MockService_AccountMetadataSymKey_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx
func (_m *MockService) Close(ctx context.Context) error {
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

// Create provides a mock function with given fields: ctx
func (_m *MockService) Create(ctx context.Context) (clientspace.Space, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 clientspace.Space
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (clientspace.Space, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) clientspace.Space); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientspace.Space)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) Create(ctx interface{}) *MockService_Create_Call {
	return &MockService_Create_Call{Call: _e.mock.On("Create", ctx)}
}

func (_c *MockService_Create_Call) Run(run func(ctx context.Context)) *MockService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_Create_Call) Return(_a0 clientspace.Space, err error) *MockService_Create_Call {
	_c.Call.Return(_a0, err)
	return _c
}

func (_c *MockService_Create_Call) RunAndReturn(run func(context.Context) (clientspace.Space, error)) *MockService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockService_Expecter) Delete(ctx interface{}, id interface{}) *MockService_Delete_Call {
	return &MockService_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockService_Delete_Call) Run(run func(ctx context.Context, id string)) *MockService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_Delete_Call) Return(err error) *MockService_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Delete_Call) RunAndReturn(run func(context.Context, string) error) *MockService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockService) Get(ctx context.Context, id string) (clientspace.Space, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 clientspace.Space
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (clientspace.Space, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) clientspace.Space); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientspace.Space)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockService_Expecter) Get(ctx interface{}, id interface{}) *MockService_Get_Call {
	return &MockService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockService_Get_Call) Run(run func(ctx context.Context, id string)) *MockService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_Get_Call) Return(_a0 clientspace.Space, err error) *MockService_Get_Call {
	_c.Call.Return(_a0, err)
	return _c
}

func (_c *MockService_Get_Call) RunAndReturn(run func(context.Context, string) (clientspace.Space, error)) *MockService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonalSpace provides a mock function with given fields: ctx
func (_m *MockService) GetPersonalSpace(ctx context.Context) (clientspace.Space, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonalSpace")
	}

	var r0 clientspace.Space
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (clientspace.Space, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) clientspace.Space); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientspace.Space)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetPersonalSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonalSpace'
type MockService_GetPersonalSpace_Call struct {
	*mock.Call
}

// GetPersonalSpace is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) GetPersonalSpace(ctx interface{}) *MockService_GetPersonalSpace_Call {
	return &MockService_GetPersonalSpace_Call{Call: _e.mock.On("GetPersonalSpace", ctx)}
}

func (_c *MockService_GetPersonalSpace_Call) Run(run func(ctx context.Context)) *MockService_GetPersonalSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_GetPersonalSpace_Call) Return(_a0 clientspace.Space, err error) *MockService_GetPersonalSpace_Call {
	_c.Call.Return(_a0, err)
	return _c
}

func (_c *MockService_GetPersonalSpace_Call) RunAndReturn(run func(context.Context) (clientspace.Space, error)) *MockService_GetPersonalSpace_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockService) Init(a *app.App) error {
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

// Join provides a mock function with given fields: ctx, id
func (_m *MockService) Join(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Join")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Join_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Join'
type MockService_Join_Call struct {
	*mock.Call
}

// Join is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockService_Expecter) Join(ctx interface{}, id interface{}) *MockService_Join_Call {
	return &MockService_Join_Call{Call: _e.mock.On("Join", ctx, id)}
}

func (_c *MockService_Join_Call) Run(run func(ctx context.Context, id string)) *MockService_Join_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_Join_Call) Return(err error) *MockService_Join_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Join_Call) RunAndReturn(run func(context.Context, string) error) *MockService_Join_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockService) Name() string {
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

// Run provides a mock function with given fields: ctx
func (_m *MockService) Run(ctx context.Context) error {
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

// SpaceViewId provides a mock function with given fields: spaceId
func (_m *MockService) SpaceViewId(spaceId string) (string, error) {
	ret := _m.Called(spaceId)

	if len(ret) == 0 {
		panic("no return value specified for SpaceViewId")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(spaceId)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(spaceId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(spaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_SpaceViewId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpaceViewId'
type MockService_SpaceViewId_Call struct {
	*mock.Call
}

// SpaceViewId is a helper method to define mock.On call
//   - spaceId string
func (_e *MockService_Expecter) SpaceViewId(spaceId interface{}) *MockService_SpaceViewId_Call {
	return &MockService_SpaceViewId_Call{Call: _e.mock.On("SpaceViewId", spaceId)}
}

func (_c *MockService_SpaceViewId_Call) Run(run func(spaceId string)) *MockService_SpaceViewId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_SpaceViewId_Call) Return(spaceViewId string, err error) *MockService_SpaceViewId_Call {
	_c.Call.Return(spaceViewId, err)
	return _c
}

func (_c *MockService_SpaceViewId_Call) RunAndReturn(run func(string) (string, error)) *MockService_SpaceViewId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
