// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore (interfaces: FileStore)
//
// Generated by this command:
//
//	mockgen -package filesync -destination filestore_mock.go github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore FileStore
//

// Package filesync is a generated GoMock package.
package filesync

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	domain "github.com/anyproto/anytype-heart/core/domain"
	objectorigin "github.com/anyproto/anytype-heart/core/domain/objectorigin"
	localstore "github.com/anyproto/anytype-heart/pkg/lib/localstore"
	storage "github.com/anyproto/anytype-heart/pkg/lib/pb/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockFileStore is a mock of FileStore interface.
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore.
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance.
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return m.recorder
}

// AddFileKeys mocks base method.
func (m *MockFileStore) AddFileKeys(arg0 ...domain.FileEncryptionKeys) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddFileKeys", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFileKeys indicates an expected call of AddFileKeys.
func (mr *MockFileStoreMockRecorder) AddFileKeys(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFileKeys", reflect.TypeOf((*MockFileStore)(nil).AddFileKeys), arg0...)
}

// AddFileVariant mocks base method.
func (m *MockFileStore) AddFileVariant(arg0 *storage.FileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFileVariant", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFileVariant indicates an expected call of AddFileVariant.
func (mr *MockFileStoreMockRecorder) AddFileVariant(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFileVariant", reflect.TypeOf((*MockFileStore)(nil).AddFileVariant), arg0)
}

// AddFileVariants mocks base method.
func (m *MockFileStore) AddFileVariants(arg0 bool, arg1 ...*storage.FileInfo) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddFileVariants", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFileVariants indicates an expected call of AddFileVariants.
func (mr *MockFileStoreMockRecorder) AddFileVariants(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFileVariants", reflect.TypeOf((*MockFileStore)(nil).AddFileVariants), varargs...)
}

// Close mocks base method.
func (m *MockFileStore) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileStoreMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileStore)(nil).Close), arg0)
}

// DeleteFile mocks base method.
func (m *MockFileStore) DeleteFile(arg0 domain.FileId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockFileStoreMockRecorder) DeleteFile(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileStore)(nil).DeleteFile), arg0)
}

// DeleteFileVariants mocks base method.
func (m *MockFileStore) DeleteFileVariants(arg0 []domain.FileContentId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFileVariants", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFileVariants indicates an expected call of DeleteFileVariants.
func (mr *MockFileStoreMockRecorder) DeleteFileVariants(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileVariants", reflect.TypeOf((*MockFileStore)(nil).DeleteFileVariants), arg0)
}

// GetChunksCount mocks base method.
func (m *MockFileStore) GetChunksCount(arg0 domain.FileId) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChunksCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChunksCount indicates an expected call of GetChunksCount.
func (mr *MockFileStoreMockRecorder) GetChunksCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChunksCount", reflect.TypeOf((*MockFileStore)(nil).GetChunksCount), arg0)
}

// GetFileKeys mocks base method.
func (m *MockFileStore) GetFileKeys(arg0 domain.FileId) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileKeys", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileKeys indicates an expected call of GetFileKeys.
func (mr *MockFileStoreMockRecorder) GetFileKeys(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileKeys", reflect.TypeOf((*MockFileStore)(nil).GetFileKeys), arg0)
}

// GetFileOrigin mocks base method.
func (m *MockFileStore) GetFileOrigin(arg0 domain.FileId) (objectorigin.ObjectOrigin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileOrigin", arg0)
	ret0, _ := ret[0].(objectorigin.ObjectOrigin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileOrigin indicates an expected call of GetFileOrigin.
func (mr *MockFileStoreMockRecorder) GetFileOrigin(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileOrigin", reflect.TypeOf((*MockFileStore)(nil).GetFileOrigin), arg0)
}

// GetFileSize mocks base method.
func (m *MockFileStore) GetFileSize(arg0 domain.FileId) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileSize", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileSize indicates an expected call of GetFileSize.
func (mr *MockFileStoreMockRecorder) GetFileSize(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileSize", reflect.TypeOf((*MockFileStore)(nil).GetFileSize), arg0)
}

// GetFileVariant mocks base method.
func (m *MockFileStore) GetFileVariant(arg0 domain.FileContentId) (*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileVariant", arg0)
	ret0, _ := ret[0].(*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileVariant indicates an expected call of GetFileVariant.
func (mr *MockFileStoreMockRecorder) GetFileVariant(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileVariant", reflect.TypeOf((*MockFileStore)(nil).GetFileVariant), arg0)
}

// GetFileVariantByChecksum mocks base method.
func (m *MockFileStore) GetFileVariantByChecksum(arg0, arg1 string) (*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileVariantByChecksum", arg0, arg1)
	ret0, _ := ret[0].(*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileVariantByChecksum indicates an expected call of GetFileVariantByChecksum.
func (mr *MockFileStoreMockRecorder) GetFileVariantByChecksum(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileVariantByChecksum", reflect.TypeOf((*MockFileStore)(nil).GetFileVariantByChecksum), arg0, arg1)
}

// GetFileVariantBySource mocks base method.
func (m *MockFileStore) GetFileVariantBySource(arg0, arg1, arg2 string) (*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileVariantBySource", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileVariantBySource indicates an expected call of GetFileVariantBySource.
func (mr *MockFileStoreMockRecorder) GetFileVariantBySource(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileVariantBySource", reflect.TypeOf((*MockFileStore)(nil).GetFileVariantBySource), arg0, arg1, arg2)
}

// Indexes mocks base method.
func (m *MockFileStore) Indexes() []localstore.Index {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].([]localstore.Index)
	return ret0
}

// Indexes indicates an expected call of Indexes.
func (mr *MockFileStoreMockRecorder) Indexes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexes", reflect.TypeOf((*MockFileStore)(nil).Indexes))
}

// Init mocks base method.
func (m *MockFileStore) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockFileStoreMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockFileStore)(nil).Init), arg0)
}

// IsFileImported mocks base method.
func (m *MockFileStore) IsFileImported(arg0 domain.FileId) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFileImported", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFileImported indicates an expected call of IsFileImported.
func (mr *MockFileStoreMockRecorder) IsFileImported(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFileImported", reflect.TypeOf((*MockFileStore)(nil).IsFileImported), arg0)
}

// ListAllFileVariants mocks base method.
func (m *MockFileStore) ListAllFileVariants() ([]*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllFileVariants")
	ret0, _ := ret[0].([]*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllFileVariants indicates an expected call of ListAllFileVariants.
func (mr *MockFileStoreMockRecorder) ListAllFileVariants() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllFileVariants", reflect.TypeOf((*MockFileStore)(nil).ListAllFileVariants))
}

// ListFileIds mocks base method.
func (m *MockFileStore) ListFileIds() ([]domain.FileId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFileIds")
	ret0, _ := ret[0].([]domain.FileId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFileIds indicates an expected call of ListFileIds.
func (mr *MockFileStoreMockRecorder) ListFileIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFileIds", reflect.TypeOf((*MockFileStore)(nil).ListFileIds))
}

// ListFileVariants mocks base method.
func (m *MockFileStore) ListFileVariants(arg0 domain.FileId) ([]*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFileVariants", arg0)
	ret0, _ := ret[0].([]*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFileVariants indicates an expected call of ListFileVariants.
func (mr *MockFileStoreMockRecorder) ListFileVariants(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFileVariants", reflect.TypeOf((*MockFileStore)(nil).ListFileVariants), arg0)
}

// Name mocks base method.
func (m *MockFileStore) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockFileStoreMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFileStore)(nil).Name))
}

// RemoveEmptyFileKeys mocks base method.
func (m *MockFileStore) RemoveEmptyFileKeys() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEmptyFileKeys")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEmptyFileKeys indicates an expected call of RemoveEmptyFileKeys.
func (mr *MockFileStoreMockRecorder) RemoveEmptyFileKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEmptyFileKeys", reflect.TypeOf((*MockFileStore)(nil).RemoveEmptyFileKeys))
}

// Run mocks base method.
func (m *MockFileStore) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockFileStoreMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockFileStore)(nil).Run), arg0)
}

// SetChunksCount mocks base method.
func (m *MockFileStore) SetChunksCount(arg0 domain.FileId, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetChunksCount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetChunksCount indicates an expected call of SetChunksCount.
func (mr *MockFileStoreMockRecorder) SetChunksCount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChunksCount", reflect.TypeOf((*MockFileStore)(nil).SetChunksCount), arg0, arg1)
}

// SetFileOrigin mocks base method.
func (m *MockFileStore) SetFileOrigin(arg0 domain.FileId, arg1 objectorigin.ObjectOrigin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFileOrigin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFileOrigin indicates an expected call of SetFileOrigin.
func (mr *MockFileStoreMockRecorder) SetFileOrigin(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFileOrigin", reflect.TypeOf((*MockFileStore)(nil).SetFileOrigin), arg0, arg1)
}

// SetFileSize mocks base method.
func (m *MockFileStore) SetFileSize(arg0 domain.FileId, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFileSize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFileSize indicates an expected call of SetFileSize.
func (mr *MockFileStoreMockRecorder) SetFileSize(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFileSize", reflect.TypeOf((*MockFileStore)(nil).SetFileSize), arg0, arg1)
}

// SetIsFileImported mocks base method.
func (m *MockFileStore) SetIsFileImported(arg0 domain.FileId, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIsFileImported", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIsFileImported indicates an expected call of SetIsFileImported.
func (mr *MockFileStoreMockRecorder) SetIsFileImported(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIsFileImported", reflect.TypeOf((*MockFileStore)(nil).SetIsFileImported), arg0, arg1)
}
