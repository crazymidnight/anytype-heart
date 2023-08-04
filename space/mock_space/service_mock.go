// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/space (interfaces: Service)

// Package mock_space is a generated GoMock package.
package mock_space

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	commonspace "github.com/anyproto/any-sync/commonspace"
	streampool "github.com/anyproto/any-sync/net/streampool"
	space "github.com/anyproto/anytype-heart/space"
	gomock "go.uber.org/mock/gomock"
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

// AccountId mocks base method.
func (m *MockService) AccountId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountId")
	ret0, _ := ret[0].(string)
	return ret0
}

// AccountId indicates an expected call of AccountId.
func (mr *MockServiceMockRecorder) AccountId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountId", reflect.TypeOf((*MockService)(nil).AccountId))
}

// AccountSpace mocks base method.
func (m *MockService) AccountSpace(arg0 context.Context) (commonspace.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountSpace", arg0)
	ret0, _ := ret[0].(commonspace.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountSpace indicates an expected call of AccountSpace.
func (mr *MockServiceMockRecorder) AccountSpace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountSpace", reflect.TypeOf((*MockService)(nil).AccountSpace), arg0)
}

// Close mocks base method.
func (m *MockService) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockServiceMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockService)(nil).Close), arg0)
}

// CreateSpace mocks base method.
func (m *MockService) CreateSpace(arg0 context.Context) (commonspace.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSpace", arg0)
	ret0, _ := ret[0].(commonspace.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSpace indicates an expected call of CreateSpace.
func (mr *MockServiceMockRecorder) CreateSpace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpace", reflect.TypeOf((*MockService)(nil).CreateSpace), arg0)
}

// DeleteAccount mocks base method.
func (m *MockService) DeleteAccount(arg0 context.Context, arg1 bool) (space.StatusPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(space.StatusPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockServiceMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockService)(nil).DeleteAccount), arg0, arg1)
}

// DeleteSpace mocks base method.
func (m *MockService) DeleteSpace(arg0 context.Context, arg1 string, arg2 bool) (space.StatusPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSpace", arg0, arg1, arg2)
	ret0, _ := ret[0].(space.StatusPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSpace indicates an expected call of DeleteSpace.
func (mr *MockServiceMockRecorder) DeleteSpace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSpace", reflect.TypeOf((*MockService)(nil).DeleteSpace), arg0, arg1, arg2)
}

// DeriveSpace mocks base method.
func (m *MockService) DeriveSpace(arg0 context.Context, arg1 commonspace.SpaceDerivePayload) (commonspace.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeriveSpace", arg0, arg1)
	ret0, _ := ret[0].(commonspace.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeriveSpace indicates an expected call of DeriveSpace.
func (mr *MockServiceMockRecorder) DeriveSpace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeriveSpace", reflect.TypeOf((*MockService)(nil).DeriveSpace), arg0, arg1)
}

// GetSpace mocks base method.
func (m *MockService) GetSpace(arg0 context.Context, arg1 string) (commonspace.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpace", arg0, arg1)
	ret0, _ := ret[0].(commonspace.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpace indicates an expected call of GetSpace.
func (mr *MockServiceMockRecorder) GetSpace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpace", reflect.TypeOf((*MockService)(nil).GetSpace), arg0, arg1)
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

// Run mocks base method.
func (m *MockService) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockServiceMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockService)(nil).Run), arg0)
}

// StreamPool mocks base method.
func (m *MockService) StreamPool() streampool.StreamPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamPool")
	ret0, _ := ret[0].(streampool.StreamPool)
	return ret0
}

// StreamPool indicates an expected call of StreamPool.
func (mr *MockServiceMockRecorder) StreamPool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamPool", reflect.TypeOf((*MockService)(nil).StreamPool))
}
