// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/core/block/source (interfaces: Service,Source)

// Package mockSource is a generated GoMock package.
package mockSource

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	state "github.com/anyproto/anytype-heart/core/block/editor/state"
	source "github.com/anyproto/anytype-heart/core/block/source"
	pb "github.com/anyproto/anytype-heart/pb"
	smartblock "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DetailsFromIdBasedSource mocks base method.
func (m *MockService) DetailsFromIdBasedSource(arg0 string) (*types.Struct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetailsFromIdBasedSource", arg0)
	ret0, _ := ret[0].(*types.Struct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetailsFromIdBasedSource indicates an expected call of DetailsFromIdBasedSource.
func (mr *MockServiceMockRecorder) DetailsFromIdBasedSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetailsFromIdBasedSource", reflect.TypeOf((*MockService)(nil).DetailsFromIdBasedSource), arg0)
}

// IDsListerBySmartblockType mocks base method.
func (m *MockService) IDsListerBySmartblockType(arg0 string, arg1 smartblock.SmartBlockType) (source.IDsLister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDsListerBySmartblockType", arg0, arg1)
	ret0, _ := ret[0].(source.IDsLister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IDsListerBySmartblockType indicates an expected call of IDsListerBySmartblockType.
func (mr *MockServiceMockRecorder) IDsListerBySmartblockType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDsListerBySmartblockType", reflect.TypeOf((*MockService)(nil).IDsListerBySmartblockType), arg0, arg1)
}

// Init mocks base method.
func (m *MockService) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockServiceMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockService)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockService)(nil).Name))
}

// NewSource mocks base method.
func (m *MockService) NewSource(arg0 context.Context, arg1, arg2 string, arg3 source.BuildOptions) (source.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSource", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(source.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSource indicates an expected call of NewSource.
func (mr *MockServiceMockRecorder) NewSource(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSource", reflect.TypeOf((*MockService)(nil).NewSource), arg0, arg1, arg2, arg3)
}

// NewStaticSource mocks base method.
func (m *MockService) NewStaticSource(arg0 string, arg1 model.SmartBlockType, arg2 *state.State, arg3 func(source.PushChangeParams) (string, error)) source.SourceWithType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStaticSource", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(source.SourceWithType)
	return ret0
}

// NewStaticSource indicates an expected call of NewStaticSource.
func (mr *MockServiceMockRecorder) NewStaticSource(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStaticSource", reflect.TypeOf((*MockService)(nil).NewStaticSource), arg0, arg1, arg2, arg3)
}

// RegisterStaticSource mocks base method.
func (m *MockService) RegisterStaticSource(arg0 string, arg1 source.Source) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterStaticSource", arg0, arg1)
}

// RegisterStaticSource indicates an expected call of RegisterStaticSource.
func (mr *MockServiceMockRecorder) RegisterStaticSource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStaticSource", reflect.TypeOf((*MockService)(nil).RegisterStaticSource), arg0, arg1)
}

// RemoveStaticSource mocks base method.
func (m *MockService) RemoveStaticSource(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveStaticSource", arg0)
}

// RemoveStaticSource indicates an expected call of RemoveStaticSource.
func (mr *MockServiceMockRecorder) RemoveStaticSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStaticSource", reflect.TypeOf((*MockService)(nil).RemoveStaticSource), arg0)
}

// MockSource is a mock of Source interface.
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource.
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance.
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSourceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSource)(nil).Close))
}

// GetFileKeysSnapshot mocks base method.
func (m *MockSource) GetFileKeysSnapshot() []*pb.ChangeFileKeys {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileKeysSnapshot")
	ret0, _ := ret[0].([]*pb.ChangeFileKeys)
	return ret0
}

// GetFileKeysSnapshot indicates an expected call of GetFileKeysSnapshot.
func (mr *MockSourceMockRecorder) GetFileKeysSnapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileKeysSnapshot", reflect.TypeOf((*MockSource)(nil).GetFileKeysSnapshot))
}

// Heads mocks base method.
func (m *MockSource) Heads() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Heads")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Heads indicates an expected call of Heads.
func (mr *MockSourceMockRecorder) Heads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heads", reflect.TypeOf((*MockSource)(nil).Heads))
}

// Id mocks base method.
func (m *MockSource) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockSourceMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockSource)(nil).Id))
}

// PushChange mocks base method.
func (m *MockSource) PushChange(arg0 source.PushChangeParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushChange", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushChange indicates an expected call of PushChange.
func (mr *MockSourceMockRecorder) PushChange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushChange", reflect.TypeOf((*MockSource)(nil).PushChange), arg0)
}

// ReadDoc mocks base method.
func (m *MockSource) ReadDoc(arg0 context.Context, arg1 source.ChangeReceiver, arg2 bool) (state.Doc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDoc", arg0, arg1, arg2)
	ret0, _ := ret[0].(state.Doc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDoc indicates an expected call of ReadDoc.
func (mr *MockSourceMockRecorder) ReadDoc(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDoc", reflect.TypeOf((*MockSource)(nil).ReadDoc), arg0, arg1, arg2)
}

// ReadOnly mocks base method.
func (m *MockSource) ReadOnly() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOnly")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ReadOnly indicates an expected call of ReadOnly.
func (mr *MockSourceMockRecorder) ReadOnly() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOnly", reflect.TypeOf((*MockSource)(nil).ReadOnly))
}

// Type mocks base method.
func (m *MockSource) Type() model.SmartBlockType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(model.SmartBlockType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockSourceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockSource)(nil).Type))
}
