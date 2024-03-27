// Code generated by mockery. DO NOT EDIT.

package mock_fileobject

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	domain "github.com/anyproto/anytype-heart/core/domain"

	fileobject "github.com/anyproto/anytype-heart/core/files/fileobject"

	mock "github.com/stretchr/testify/mock"

	objectorigin "github.com/anyproto/anytype-heart/core/domain/objectorigin"

	pb "github.com/anyproto/anytype-heart/pb"

	source "github.com/anyproto/anytype-heart/core/block/source"

	state "github.com/anyproto/anytype-heart/core/block/editor/state"

	types "github.com/gogo/protobuf/types"
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

// Create provides a mock function with given fields: ctx, spaceId, req
func (_m *MockService) Create(ctx context.Context, spaceId string, req fileobject.CreateRequest) (string, *types.Struct, error) {
	ret := _m.Called(ctx, spaceId, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 *types.Struct
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, fileobject.CreateRequest) (string, *types.Struct, error)); ok {
		return rf(ctx, spaceId, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, fileobject.CreateRequest) string); ok {
		r0 = rf(ctx, spaceId, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, fileobject.CreateRequest) *types.Struct); ok {
		r1 = rf(ctx, spaceId, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Struct)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, fileobject.CreateRequest) error); ok {
		r2 = rf(ctx, spaceId, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - req fileobject.CreateRequest
func (_e *MockService_Expecter) Create(ctx interface{}, spaceId interface{}, req interface{}) *MockService_Create_Call {
	return &MockService_Create_Call{Call: _e.mock.On("Create", ctx, spaceId, req)}
}

func (_c *MockService_Create_Call) Run(run func(ctx context.Context, spaceId string, req fileobject.CreateRequest)) *MockService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(fileobject.CreateRequest))
	})
	return _c
}

func (_c *MockService_Create_Call) Return(id string, object *types.Struct, err error) *MockService_Create_Call {
	_c.Call.Return(id, object, err)
	return _c
}

func (_c *MockService_Create_Call) RunAndReturn(run func(context.Context, string, fileobject.CreateRequest) (string, *types.Struct, error)) *MockService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateFromImport provides a mock function with given fields: fileId, origin
func (_m *MockService) CreateFromImport(fileId domain.FullFileId, origin objectorigin.ObjectOrigin) (string, error) {
	ret := _m.Called(fileId, origin)

	if len(ret) == 0 {
		panic("no return value specified for CreateFromImport")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.FullFileId, objectorigin.ObjectOrigin) (string, error)); ok {
		return rf(fileId, origin)
	}
	if rf, ok := ret.Get(0).(func(domain.FullFileId, objectorigin.ObjectOrigin) string); ok {
		r0 = rf(fileId, origin)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.FullFileId, objectorigin.ObjectOrigin) error); ok {
		r1 = rf(fileId, origin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_CreateFromImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateFromImport'
type MockService_CreateFromImport_Call struct {
	*mock.Call
}

// CreateFromImport is a helper method to define mock.On call
//   - fileId domain.FullFileId
//   - origin objectorigin.ObjectOrigin
func (_e *MockService_Expecter) CreateFromImport(fileId interface{}, origin interface{}) *MockService_CreateFromImport_Call {
	return &MockService_CreateFromImport_Call{Call: _e.mock.On("CreateFromImport", fileId, origin)}
}

func (_c *MockService_CreateFromImport_Call) Run(run func(fileId domain.FullFileId, origin objectorigin.ObjectOrigin)) *MockService_CreateFromImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FullFileId), args[1].(objectorigin.ObjectOrigin))
	})
	return _c
}

func (_c *MockService_CreateFromImport_Call) Return(_a0 string, _a1 error) *MockService_CreateFromImport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_CreateFromImport_Call) RunAndReturn(run func(domain.FullFileId, objectorigin.ObjectOrigin) (string, error)) *MockService_CreateFromImport_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteFileData provides a mock function with given fields: objectId
func (_m *MockService) DeleteFileData(objectId string) error {
	ret := _m.Called(objectId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFileData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(objectId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_DeleteFileData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFileData'
type MockService_DeleteFileData_Call struct {
	*mock.Call
}

// DeleteFileData is a helper method to define mock.On call
//   - objectId string
func (_e *MockService_Expecter) DeleteFileData(objectId interface{}) *MockService_DeleteFileData_Call {
	return &MockService_DeleteFileData_Call{Call: _e.mock.On("DeleteFileData", objectId)}
}

func (_c *MockService_DeleteFileData_Call) Run(run func(objectId string)) *MockService_DeleteFileData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_DeleteFileData_Call) Return(_a0 error) *MockService_DeleteFileData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_DeleteFileData_Call) RunAndReturn(run func(string) error) *MockService_DeleteFileData_Call {
	_c.Call.Return(run)
	return _c
}

// FileOffload provides a mock function with given fields: ctx, objectId, includeNotPinned
func (_m *MockService) FileOffload(ctx context.Context, objectId string, includeNotPinned bool) (uint64, error) {
	ret := _m.Called(ctx, objectId, includeNotPinned)

	if len(ret) == 0 {
		panic("no return value specified for FileOffload")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) (uint64, error)); ok {
		return rf(ctx, objectId, includeNotPinned)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) uint64); ok {
		r0 = rf(ctx, objectId, includeNotPinned)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, objectId, includeNotPinned)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_FileOffload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FileOffload'
type MockService_FileOffload_Call struct {
	*mock.Call
}

// FileOffload is a helper method to define mock.On call
//   - ctx context.Context
//   - objectId string
//   - includeNotPinned bool
func (_e *MockService_Expecter) FileOffload(ctx interface{}, objectId interface{}, includeNotPinned interface{}) *MockService_FileOffload_Call {
	return &MockService_FileOffload_Call{Call: _e.mock.On("FileOffload", ctx, objectId, includeNotPinned)}
}

func (_c *MockService_FileOffload_Call) Run(run func(ctx context.Context, objectId string, includeNotPinned bool)) *MockService_FileOffload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockService_FileOffload_Call) Return(totalSize uint64, err error) *MockService_FileOffload_Call {
	_c.Call.Return(totalSize, err)
	return _c
}

func (_c *MockService_FileOffload_Call) RunAndReturn(run func(context.Context, string, bool) (uint64, error)) *MockService_FileOffload_Call {
	_c.Call.Return(run)
	return _c
}

// FileSpaceOffload provides a mock function with given fields: ctx, spaceId, includeNotPinned
func (_m *MockService) FileSpaceOffload(ctx context.Context, spaceId string, includeNotPinned bool) (int, uint64, error) {
	ret := _m.Called(ctx, spaceId, includeNotPinned)

	if len(ret) == 0 {
		panic("no return value specified for FileSpaceOffload")
	}

	var r0 int
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) (int, uint64, error)); ok {
		return rf(ctx, spaceId, includeNotPinned)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) int); ok {
		r0 = rf(ctx, spaceId, includeNotPinned)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, bool) uint64); ok {
		r1 = rf(ctx, spaceId, includeNotPinned)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, bool) error); ok {
		r2 = rf(ctx, spaceId, includeNotPinned)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockService_FileSpaceOffload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FileSpaceOffload'
type MockService_FileSpaceOffload_Call struct {
	*mock.Call
}

// FileSpaceOffload is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - includeNotPinned bool
func (_e *MockService_Expecter) FileSpaceOffload(ctx interface{}, spaceId interface{}, includeNotPinned interface{}) *MockService_FileSpaceOffload_Call {
	return &MockService_FileSpaceOffload_Call{Call: _e.mock.On("FileSpaceOffload", ctx, spaceId, includeNotPinned)}
}

func (_c *MockService_FileSpaceOffload_Call) Run(run func(ctx context.Context, spaceId string, includeNotPinned bool)) *MockService_FileSpaceOffload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockService_FileSpaceOffload_Call) Return(filesOffloaded int, totalSize uint64, err error) *MockService_FileSpaceOffload_Call {
	_c.Call.Return(filesOffloaded, totalSize, err)
	return _c
}

func (_c *MockService_FileSpaceOffload_Call) RunAndReturn(run func(context.Context, string, bool) (int, uint64, error)) *MockService_FileSpaceOffload_Call {
	_c.Call.Return(run)
	return _c
}

// FilesOffload provides a mock function with given fields: ctx, objectIds, includeNotPinned
func (_m *MockService) FilesOffload(ctx context.Context, objectIds []string, includeNotPinned bool) (int, uint64, error) {
	ret := _m.Called(ctx, objectIds, includeNotPinned)

	if len(ret) == 0 {
		panic("no return value specified for FilesOffload")
	}

	var r0 int
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, bool) (int, uint64, error)); ok {
		return rf(ctx, objectIds, includeNotPinned)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, bool) int); ok {
		r0 = rf(ctx, objectIds, includeNotPinned)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, bool) uint64); ok {
		r1 = rf(ctx, objectIds, includeNotPinned)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []string, bool) error); ok {
		r2 = rf(ctx, objectIds, includeNotPinned)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockService_FilesOffload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilesOffload'
type MockService_FilesOffload_Call struct {
	*mock.Call
}

// FilesOffload is a helper method to define mock.On call
//   - ctx context.Context
//   - objectIds []string
//   - includeNotPinned bool
func (_e *MockService_Expecter) FilesOffload(ctx interface{}, objectIds interface{}, includeNotPinned interface{}) *MockService_FilesOffload_Call {
	return &MockService_FilesOffload_Call{Call: _e.mock.On("FilesOffload", ctx, objectIds, includeNotPinned)}
}

func (_c *MockService_FilesOffload_Call) Run(run func(ctx context.Context, objectIds []string, includeNotPinned bool)) *MockService_FilesOffload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(bool))
	})
	return _c
}

func (_c *MockService_FilesOffload_Call) Return(filesOffloaded int, totalSize uint64, err error) *MockService_FilesOffload_Call {
	_c.Call.Return(filesOffloaded, totalSize, err)
	return _c
}

func (_c *MockService_FilesOffload_Call) RunAndReturn(run func(context.Context, []string, bool) (int, uint64, error)) *MockService_FilesOffload_Call {
	_c.Call.Return(run)
	return _c
}

// GetFileIdFromObject provides a mock function with given fields: objectId
func (_m *MockService) GetFileIdFromObject(objectId string) (domain.FullFileId, error) {
	ret := _m.Called(objectId)

	if len(ret) == 0 {
		panic("no return value specified for GetFileIdFromObject")
	}

	var r0 domain.FullFileId
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.FullFileId, error)); ok {
		return rf(objectId)
	}
	if rf, ok := ret.Get(0).(func(string) domain.FullFileId); ok {
		r0 = rf(objectId)
	} else {
		r0 = ret.Get(0).(domain.FullFileId)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(objectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetFileIdFromObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFileIdFromObject'
type MockService_GetFileIdFromObject_Call struct {
	*mock.Call
}

// GetFileIdFromObject is a helper method to define mock.On call
//   - objectId string
func (_e *MockService_Expecter) GetFileIdFromObject(objectId interface{}) *MockService_GetFileIdFromObject_Call {
	return &MockService_GetFileIdFromObject_Call{Call: _e.mock.On("GetFileIdFromObject", objectId)}
}

func (_c *MockService_GetFileIdFromObject_Call) Run(run func(objectId string)) *MockService_GetFileIdFromObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_GetFileIdFromObject_Call) Return(_a0 domain.FullFileId, _a1 error) *MockService_GetFileIdFromObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetFileIdFromObject_Call) RunAndReturn(run func(string) (domain.FullFileId, error)) *MockService_GetFileIdFromObject_Call {
	_c.Call.Return(run)
	return _c
}

// GetFileIdFromObjectWaitLoad provides a mock function with given fields: ctx, objectId
func (_m *MockService) GetFileIdFromObjectWaitLoad(ctx context.Context, objectId string) (domain.FullFileId, error) {
	ret := _m.Called(ctx, objectId)

	if len(ret) == 0 {
		panic("no return value specified for GetFileIdFromObjectWaitLoad")
	}

	var r0 domain.FullFileId
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.FullFileId, error)); ok {
		return rf(ctx, objectId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.FullFileId); ok {
		r0 = rf(ctx, objectId)
	} else {
		r0 = ret.Get(0).(domain.FullFileId)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, objectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetFileIdFromObjectWaitLoad_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFileIdFromObjectWaitLoad'
type MockService_GetFileIdFromObjectWaitLoad_Call struct {
	*mock.Call
}

// GetFileIdFromObjectWaitLoad is a helper method to define mock.On call
//   - ctx context.Context
//   - objectId string
func (_e *MockService_Expecter) GetFileIdFromObjectWaitLoad(ctx interface{}, objectId interface{}) *MockService_GetFileIdFromObjectWaitLoad_Call {
	return &MockService_GetFileIdFromObjectWaitLoad_Call{Call: _e.mock.On("GetFileIdFromObjectWaitLoad", ctx, objectId)}
}

func (_c *MockService_GetFileIdFromObjectWaitLoad_Call) Run(run func(ctx context.Context, objectId string)) *MockService_GetFileIdFromObjectWaitLoad_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_GetFileIdFromObjectWaitLoad_Call) Return(_a0 domain.FullFileId, _a1 error) *MockService_GetFileIdFromObjectWaitLoad_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetFileIdFromObjectWaitLoad_Call) RunAndReturn(run func(context.Context, string) (domain.FullFileId, error)) *MockService_GetFileIdFromObjectWaitLoad_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectDetailsByFileId provides a mock function with given fields: fileId
func (_m *MockService) GetObjectDetailsByFileId(fileId domain.FullFileId) (string, *types.Struct, error) {
	ret := _m.Called(fileId)

	if len(ret) == 0 {
		panic("no return value specified for GetObjectDetailsByFileId")
	}

	var r0 string
	var r1 *types.Struct
	var r2 error
	if rf, ok := ret.Get(0).(func(domain.FullFileId) (string, *types.Struct, error)); ok {
		return rf(fileId)
	}
	if rf, ok := ret.Get(0).(func(domain.FullFileId) string); ok {
		r0 = rf(fileId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.FullFileId) *types.Struct); ok {
		r1 = rf(fileId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Struct)
		}
	}

	if rf, ok := ret.Get(2).(func(domain.FullFileId) error); ok {
		r2 = rf(fileId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockService_GetObjectDetailsByFileId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectDetailsByFileId'
type MockService_GetObjectDetailsByFileId_Call struct {
	*mock.Call
}

// GetObjectDetailsByFileId is a helper method to define mock.On call
//   - fileId domain.FullFileId
func (_e *MockService_Expecter) GetObjectDetailsByFileId(fileId interface{}) *MockService_GetObjectDetailsByFileId_Call {
	return &MockService_GetObjectDetailsByFileId_Call{Call: _e.mock.On("GetObjectDetailsByFileId", fileId)}
}

func (_c *MockService_GetObjectDetailsByFileId_Call) Run(run func(fileId domain.FullFileId)) *MockService_GetObjectDetailsByFileId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FullFileId))
	})
	return _c
}

func (_c *MockService_GetObjectDetailsByFileId_Call) Return(_a0 string, _a1 *types.Struct, _a2 error) *MockService_GetObjectDetailsByFileId_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockService_GetObjectDetailsByFileId_Call) RunAndReturn(run func(domain.FullFileId) (string, *types.Struct, error)) *MockService_GetObjectDetailsByFileId_Call {
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

// InitEmptyFileState provides a mock function with given fields: st
func (_m *MockService) InitEmptyFileState(st *state.State) {
	_m.Called(st)
}

// MockService_InitEmptyFileState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitEmptyFileState'
type MockService_InitEmptyFileState_Call struct {
	*mock.Call
}

// InitEmptyFileState is a helper method to define mock.On call
//   - st *state.State
func (_e *MockService_Expecter) InitEmptyFileState(st interface{}) *MockService_InitEmptyFileState_Call {
	return &MockService_InitEmptyFileState_Call{Call: _e.mock.On("InitEmptyFileState", st)}
}

func (_c *MockService_InitEmptyFileState_Call) Run(run func(st *state.State)) *MockService_InitEmptyFileState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*state.State))
	})
	return _c
}

func (_c *MockService_InitEmptyFileState_Call) Return() *MockService_InitEmptyFileState_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockService_InitEmptyFileState_Call) RunAndReturn(run func(*state.State)) *MockService_InitEmptyFileState_Call {
	_c.Call.Return(run)
	return _c
}

// MigrateBlocks provides a mock function with given fields: st, spc, keys
func (_m *MockService) MigrateBlocks(st *state.State, spc source.Space, keys []*pb.ChangeFileKeys) {
	_m.Called(st, spc, keys)
}

// MockService_MigrateBlocks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MigrateBlocks'
type MockService_MigrateBlocks_Call struct {
	*mock.Call
}

// MigrateBlocks is a helper method to define mock.On call
//   - st *state.State
//   - spc source.Space
//   - keys []*pb.ChangeFileKeys
func (_e *MockService_Expecter) MigrateBlocks(st interface{}, spc interface{}, keys interface{}) *MockService_MigrateBlocks_Call {
	return &MockService_MigrateBlocks_Call{Call: _e.mock.On("MigrateBlocks", st, spc, keys)}
}

func (_c *MockService_MigrateBlocks_Call) Run(run func(st *state.State, spc source.Space, keys []*pb.ChangeFileKeys)) *MockService_MigrateBlocks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*state.State), args[1].(source.Space), args[2].([]*pb.ChangeFileKeys))
	})
	return _c
}

func (_c *MockService_MigrateBlocks_Call) Return() *MockService_MigrateBlocks_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockService_MigrateBlocks_Call) RunAndReturn(run func(*state.State, source.Space, []*pb.ChangeFileKeys)) *MockService_MigrateBlocks_Call {
	_c.Call.Return(run)
	return _c
}

// MigrateDetails provides a mock function with given fields: st, spc, keys
func (_m *MockService) MigrateDetails(st *state.State, spc source.Space, keys []*pb.ChangeFileKeys) {
	_m.Called(st, spc, keys)
}

// MockService_MigrateDetails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MigrateDetails'
type MockService_MigrateDetails_Call struct {
	*mock.Call
}

// MigrateDetails is a helper method to define mock.On call
//   - st *state.State
//   - spc source.Space
//   - keys []*pb.ChangeFileKeys
func (_e *MockService_Expecter) MigrateDetails(st interface{}, spc interface{}, keys interface{}) *MockService_MigrateDetails_Call {
	return &MockService_MigrateDetails_Call{Call: _e.mock.On("MigrateDetails", st, spc, keys)}
}

func (_c *MockService_MigrateDetails_Call) Run(run func(st *state.State, spc source.Space, keys []*pb.ChangeFileKeys)) *MockService_MigrateDetails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*state.State), args[1].(source.Space), args[2].([]*pb.ChangeFileKeys))
	})
	return _c
}

func (_c *MockService_MigrateDetails_Call) Return() *MockService_MigrateDetails_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockService_MigrateDetails_Call) RunAndReturn(run func(*state.State, source.Space, []*pb.ChangeFileKeys)) *MockService_MigrateDetails_Call {
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
