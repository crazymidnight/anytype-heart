// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/core/block/import (interfaces: Creator,IDGetter)
//
// Generated by this command:
//
//	mockgen -package importer -destination mock.go github.com/anyproto/anytype-heart/core/block/import Creator,IDGetter
//
// Package importer is a generated GoMock package.
package importer

import (
	reflect "reflect"
	time "time"

	treestorage "github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	converter "github.com/anyproto/anytype-heart/core/block/import/converter"
	types "github.com/gogo/protobuf/types"
	gomock "go.uber.org/mock/gomock"
)

// MockCreator is a mock of Creator interface.
type MockCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCreatorMockRecorder
}

// MockCreatorMockRecorder is the mock recorder for MockCreator.
type MockCreatorMockRecorder struct {
	mock *MockCreator
}

// NewMockCreator creates a new mock instance.
func NewMockCreator(ctrl *gomock.Controller) *MockCreator {
	mock := &MockCreator{ctrl: ctrl}
	mock.recorder = &MockCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreator) EXPECT() *MockCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCreator) Create(arg0 *DataObject, arg1 *converter.Snapshot) (*types.Struct, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*types.Struct)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockCreatorMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreator)(nil).Create), arg0, arg1)
}

// MockIDGetter is a mock of IDGetter interface.
type MockIDGetter struct {
	ctrl     *gomock.Controller
	recorder *MockIDGetterMockRecorder
}

// MockIDGetterMockRecorder is the mock recorder for MockIDGetter.
type MockIDGetterMockRecorder struct {
	mock *MockIDGetter
}

// NewMockIDGetter creates a new mock instance.
func NewMockIDGetter(ctrl *gomock.Controller) *MockIDGetter {
	mock := &MockIDGetter{ctrl: ctrl}
	mock.recorder = &MockIDGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDGetter) EXPECT() *MockIDGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIDGetter) Get(arg0 string, arg1 *converter.Snapshot, arg2 time.Time, arg3 bool) (string, treestorage.TreeStorageCreatePayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(treestorage.TreeStorageCreatePayload)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockIDGetterMockRecorder) Get(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIDGetter)(nil).Get), arg0, arg1, arg2, arg3)
}
