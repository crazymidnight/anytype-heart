// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore (interfaces: ObjectStore)

// Package testMock is a generated GoMock package.
package testMock

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	coordinatorproto "github.com/anyproto/any-sync/coordinator/coordinatorproto"
	types "github.com/gogo/protobuf/types"
	gomock "go.uber.org/mock/gomock"

	smartblock "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	database "github.com/anyproto/anytype-heart/pkg/lib/database"
	ftsearch "github.com/anyproto/anytype-heart/pkg/lib/localstore/ftsearch"
	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"
)

// MockObjectStore is a mock of ObjectStore interface.
type MockObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreMockRecorder
}

// MockObjectStoreMockRecorder is the mock recorder for MockObjectStore.
type MockObjectStoreMockRecorder struct {
	mock *MockObjectStore
}

// NewMockObjectStore creates a new mock instance.
func NewMockObjectStore(ctrl *gomock.Controller) *MockObjectStore {
	mock := &MockObjectStore{ctrl: ctrl}
	mock.recorder = &MockObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStore) EXPECT() *MockObjectStoreMockRecorder {
	return m.recorder
}

// AddToIndexQueue mocks base method.
func (m *MockObjectStore) AddToIndexQueue(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToIndexQueue", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToIndexQueue indicates an expected call of AddToIndexQueue.
func (mr *MockObjectStoreMockRecorder) AddToIndexQueue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToIndexQueue", reflect.TypeOf((*MockObjectStore)(nil).AddToIndexQueue), arg0)
}

// Close mocks base method.
func (m *MockObjectStore) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockObjectStoreMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObjectStore)(nil).Close), arg0)
}

// DeleteDetails mocks base method.
func (m *MockObjectStore) DeleteDetails(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDetails", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDetails indicates an expected call of DeleteDetails.
func (mr *MockObjectStoreMockRecorder) DeleteDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDetails", reflect.TypeOf((*MockObjectStore)(nil).DeleteDetails), arg0)
}

// DeleteObject mocks base method.
func (m *MockObjectStore) DeleteObject(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockObjectStoreMockRecorder) DeleteObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockObjectStore)(nil).DeleteObject), arg0)
}

// EraseIndexes mocks base method.
func (m *MockObjectStore) EraseIndexes() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EraseIndexes")
	ret0, _ := ret[0].(error)
	return ret0
}

// EraseIndexes indicates an expected call of EraseIndexes.
func (mr *MockObjectStoreMockRecorder) EraseIndexes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EraseIndexes", reflect.TypeOf((*MockObjectStore)(nil).EraseIndexes))
}

// FTSearch mocks base method.
func (m *MockObjectStore) FTSearch() ftsearch.FTSearch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FTSearch")
	ret0, _ := ret[0].(ftsearch.FTSearch)
	return ret0
}

// FTSearch indicates an expected call of FTSearch.
func (mr *MockObjectStoreMockRecorder) FTSearch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FTSearch", reflect.TypeOf((*MockObjectStore)(nil).FTSearch))
}

// GetAccountStatus mocks base method.
func (m *MockObjectStore) GetAccountStatus() (*coordinatorproto.SpaceStatusPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountStatus")
	ret0, _ := ret[0].(*coordinatorproto.SpaceStatusPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountStatus indicates an expected call of GetAccountStatus.
func (mr *MockObjectStoreMockRecorder) GetAccountStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountStatus", reflect.TypeOf((*MockObjectStore)(nil).GetAccountStatus))
}

// GetByIDs mocks base method.
func (m *MockObjectStore) GetByIDs(arg0 string, arg1 []string) ([]*model.ObjectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", arg0, arg1)
	ret0, _ := ret[0].([]*model.ObjectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockObjectStoreMockRecorder) GetByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockObjectStore)(nil).GetByIDs), arg0, arg1)
}

// GetChecksums mocks base method.
func (m *MockObjectStore) GetChecksums() (*model.ObjectStoreChecksums, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChecksums")
	ret0, _ := ret[0].(*model.ObjectStoreChecksums)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecksums indicates an expected call of GetChecksums.
func (mr *MockObjectStoreMockRecorder) GetChecksums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChecksums", reflect.TypeOf((*MockObjectStore)(nil).GetChecksums))
}

// GetCurrentWorkspaceID mocks base method.
func (m *MockObjectStore) GetCurrentWorkspaceID() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentWorkspaceID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentWorkspaceID indicates an expected call of GetCurrentWorkspaceID.
func (mr *MockObjectStoreMockRecorder) GetCurrentWorkspaceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentWorkspaceID", reflect.TypeOf((*MockObjectStore)(nil).GetCurrentWorkspaceID))
}

// GetDetails mocks base method.
func (m *MockObjectStore) GetDetails(arg0 string) (*model.ObjectDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetails", arg0)
	ret0, _ := ret[0].(*model.ObjectDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetails indicates an expected call of GetDetails.
func (mr *MockObjectStoreMockRecorder) GetDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetails", reflect.TypeOf((*MockObjectStore)(nil).GetDetails), arg0)
}

// GetInboundLinksByID mocks base method.
func (m *MockObjectStore) GetInboundLinksByID(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInboundLinksByID", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInboundLinksByID indicates an expected call of GetInboundLinksByID.
func (mr *MockObjectStoreMockRecorder) GetInboundLinksByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInboundLinksByID", reflect.TypeOf((*MockObjectStore)(nil).GetInboundLinksByID), arg0)
}

// GetLastIndexedHeadsHash mocks base method.
func (m *MockObjectStore) GetLastIndexedHeadsHash(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastIndexedHeadsHash", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastIndexedHeadsHash indicates an expected call of GetLastIndexedHeadsHash.
func (mr *MockObjectStoreMockRecorder) GetLastIndexedHeadsHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastIndexedHeadsHash", reflect.TypeOf((*MockObjectStore)(nil).GetLastIndexedHeadsHash), arg0)
}

// GetOutboundLinksByID mocks base method.
func (m *MockObjectStore) GetOutboundLinksByID(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundLinksByID", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutboundLinksByID indicates an expected call of GetOutboundLinksByID.
func (mr *MockObjectStoreMockRecorder) GetOutboundLinksByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundLinksByID", reflect.TypeOf((*MockObjectStore)(nil).GetOutboundLinksByID), arg0)
}

// GetWithLinksInfoByID mocks base method.
func (m *MockObjectStore) GetWithLinksInfoByID(arg0, arg1 string) (*model.ObjectInfoWithLinks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithLinksInfoByID", arg0, arg1)
	ret0, _ := ret[0].(*model.ObjectInfoWithLinks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithLinksInfoByID indicates an expected call of GetWithLinksInfoByID.
func (mr *MockObjectStoreMockRecorder) GetWithLinksInfoByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithLinksInfoByID", reflect.TypeOf((*MockObjectStore)(nil).GetWithLinksInfoByID), arg0, arg1)
}

// HasIDs mocks base method.
func (m *MockObjectStore) HasIDs(arg0 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasIDs", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasIDs indicates an expected call of HasIDs.
func (mr *MockObjectStoreMockRecorder) HasIDs(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasIDs", reflect.TypeOf((*MockObjectStore)(nil).HasIDs), arg0...)
}

// Init mocks base method.
func (m *MockObjectStore) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockObjectStoreMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockObjectStore)(nil).Init), arg0)
}

// List mocks base method.
func (m *MockObjectStore) List(arg0 string) ([]*model.ObjectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*model.ObjectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockObjectStoreMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockObjectStore)(nil).List), arg0)
}

// ListIDsFromFullTextQueue mocks base method.
func (m *MockObjectStore) ListIDsFromFullTextQueue() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIDsFromFullTextQueue")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIDsFromFullTextQueue indicates an expected call of ListIDsFromFullTextQueue.
func (mr *MockObjectStoreMockRecorder) ListIDsFromFullTextQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIDsFromFullTextQueue", reflect.TypeOf((*MockObjectStore)(nil).ListIDsFromFullTextQueue))
}

// ListIds mocks base method.
func (m *MockObjectStore) ListIds() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIds indicates an expected call of ListIds.
func (mr *MockObjectStoreMockRecorder) ListIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIds", reflect.TypeOf((*MockObjectStore)(nil).ListIds))
}

// Name mocks base method.
func (m *MockObjectStore) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockObjectStoreMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockObjectStore)(nil).Name))
}

// Query mocks base method.
func (m *MockObjectStore) Query(arg0 database.Query) ([]database.Record, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query.
func (mr *MockObjectStoreMockRecorder) Query(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockObjectStore)(nil).Query), arg0)
}

// QueryByID mocks base method.
func (m *MockObjectStore) QueryByID(arg0 []string) ([]database.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryByID", arg0)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryByID indicates an expected call of QueryByID.
func (mr *MockObjectStoreMockRecorder) QueryByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByID", reflect.TypeOf((*MockObjectStore)(nil).QueryByID), arg0)
}

// QueryByIDAndSubscribeForChanges mocks base method.
func (m *MockObjectStore) QueryByIDAndSubscribeForChanges(arg0 []string, arg1 database.Subscription) ([]database.Record, func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryByIDAndSubscribeForChanges", arg0, arg1)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(func())
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryByIDAndSubscribeForChanges indicates an expected call of QueryByIDAndSubscribeForChanges.
func (mr *MockObjectStoreMockRecorder) QueryByIDAndSubscribeForChanges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByIDAndSubscribeForChanges", reflect.TypeOf((*MockObjectStore)(nil).QueryByIDAndSubscribeForChanges), arg0, arg1)
}

// QueryObjectIDs mocks base method.
func (m *MockObjectStore) QueryObjectIDs(arg0 database.Query, arg1 []smartblock.SmartBlockType) ([]string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryObjectIDs", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryObjectIDs indicates an expected call of QueryObjectIDs.
func (mr *MockObjectStoreMockRecorder) QueryObjectIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryObjectIDs", reflect.TypeOf((*MockObjectStore)(nil).QueryObjectIDs), arg0, arg1)
}

// QueryRaw mocks base method.
func (m *MockObjectStore) QueryRaw(arg0 *database.Filters, arg1, arg2 int) ([]database.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRaw", arg0, arg1, arg2)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRaw indicates an expected call of QueryRaw.
func (mr *MockObjectStoreMockRecorder) QueryRaw(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRaw", reflect.TypeOf((*MockObjectStore)(nil).QueryRaw), arg0, arg1, arg2)
}

// RemoveCurrentWorkspaceID mocks base method.
func (m *MockObjectStore) RemoveCurrentWorkspaceID() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCurrentWorkspaceID")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCurrentWorkspaceID indicates an expected call of RemoveCurrentWorkspaceID.
func (mr *MockObjectStoreMockRecorder) RemoveCurrentWorkspaceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCurrentWorkspaceID", reflect.TypeOf((*MockObjectStore)(nil).RemoveCurrentWorkspaceID))
}

// RemoveIDsFromFullTextQueue mocks base method.
func (m *MockObjectStore) RemoveIDsFromFullTextQueue(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveIDsFromFullTextQueue", arg0)
}

// RemoveIDsFromFullTextQueue indicates an expected call of RemoveIDsFromFullTextQueue.
func (mr *MockObjectStoreMockRecorder) RemoveIDsFromFullTextQueue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIDsFromFullTextQueue", reflect.TypeOf((*MockObjectStore)(nil).RemoveIDsFromFullTextQueue), arg0)
}

// ResolveSpaceID mocks base method.
func (m *MockObjectStore) ResolveSpaceID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveSpaceID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveSpaceID indicates an expected call of ResolveSpaceID.
func (mr *MockObjectStoreMockRecorder) ResolveSpaceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveSpaceID", reflect.TypeOf((*MockObjectStore)(nil).ResolveSpaceID), arg0)
}

// Run mocks base method.
func (m *MockObjectStore) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockObjectStoreMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockObjectStore)(nil).Run), arg0)
}

// SaveAccountStatus mocks base method.
func (m *MockObjectStore) SaveAccountStatus(arg0 *coordinatorproto.SpaceStatusPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAccountStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAccountStatus indicates an expected call of SaveAccountStatus.
func (mr *MockObjectStoreMockRecorder) SaveAccountStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAccountStatus", reflect.TypeOf((*MockObjectStore)(nil).SaveAccountStatus), arg0)
}

// SaveChecksums mocks base method.
func (m *MockObjectStore) SaveChecksums(arg0 *model.ObjectStoreChecksums) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveChecksums", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveChecksums indicates an expected call of SaveChecksums.
func (mr *MockObjectStoreMockRecorder) SaveChecksums(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveChecksums", reflect.TypeOf((*MockObjectStore)(nil).SaveChecksums), arg0)
}

// SaveLastIndexedHeadsHash mocks base method.
func (m *MockObjectStore) SaveLastIndexedHeadsHash(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLastIndexedHeadsHash", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLastIndexedHeadsHash indicates an expected call of SaveLastIndexedHeadsHash.
func (mr *MockObjectStoreMockRecorder) SaveLastIndexedHeadsHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLastIndexedHeadsHash", reflect.TypeOf((*MockObjectStore)(nil).SaveLastIndexedHeadsHash), arg0, arg1)
}

// SetCurrentWorkspaceID mocks base method.
func (m *MockObjectStore) SetCurrentWorkspaceID(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurrentWorkspaceID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCurrentWorkspaceID indicates an expected call of SetCurrentWorkspaceID.
func (mr *MockObjectStoreMockRecorder) SetCurrentWorkspaceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentWorkspaceID", reflect.TypeOf((*MockObjectStore)(nil).SetCurrentWorkspaceID), arg0)
}

// StoreSpaceID mocks base method.
func (m *MockObjectStore) StoreSpaceID(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSpaceID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreSpaceID indicates an expected call of StoreSpaceID.
func (mr *MockObjectStoreMockRecorder) StoreSpaceID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSpaceID", reflect.TypeOf((*MockObjectStore)(nil).StoreSpaceID), arg0, arg1)
}

// SubscribeForAll mocks base method.
func (m *MockObjectStore) SubscribeForAll(arg0 func(database.Record)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubscribeForAll", arg0)
}

// SubscribeForAll indicates an expected call of SubscribeForAll.
func (mr *MockObjectStoreMockRecorder) SubscribeForAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeForAll", reflect.TypeOf((*MockObjectStore)(nil).SubscribeForAll), arg0)
}

// UpdateObjectDetails mocks base method.
func (m *MockObjectStore) UpdateObjectDetails(arg0 string, arg1 *types.Struct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectDetails", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateObjectDetails indicates an expected call of UpdateObjectDetails.
func (mr *MockObjectStoreMockRecorder) UpdateObjectDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectDetails", reflect.TypeOf((*MockObjectStore)(nil).UpdateObjectDetails), arg0, arg1)
}

// UpdateObjectLinks mocks base method.
func (m *MockObjectStore) UpdateObjectLinks(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectLinks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateObjectLinks indicates an expected call of UpdateObjectLinks.
func (mr *MockObjectStoreMockRecorder) UpdateObjectLinks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectLinks", reflect.TypeOf((*MockObjectStore)(nil).UpdateObjectLinks), arg0, arg1)
}

// UpdateObjectSnippet mocks base method.
func (m *MockObjectStore) UpdateObjectSnippet(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectSnippet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateObjectSnippet indicates an expected call of UpdateObjectSnippet.
func (mr *MockObjectStoreMockRecorder) UpdateObjectSnippet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectSnippet", reflect.TypeOf((*MockObjectStore)(nil).UpdateObjectSnippet), arg0, arg1)
}

// UpdatePendingLocalDetails mocks base method.
func (m *MockObjectStore) UpdatePendingLocalDetails(arg0 string, arg1 func(*types.Struct) (*types.Struct, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePendingLocalDetails", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePendingLocalDetails indicates an expected call of UpdatePendingLocalDetails.
func (mr *MockObjectStoreMockRecorder) UpdatePendingLocalDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePendingLocalDetails", reflect.TypeOf((*MockObjectStore)(nil).UpdatePendingLocalDetails), arg0, arg1)
}
