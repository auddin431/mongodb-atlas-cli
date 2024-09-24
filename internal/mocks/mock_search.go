// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: SearchIndexLister,SearchIndexCreator,SearchIndexDescriber,SearchIndexUpdater,SearchIndexDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20240805001/admin"
)

// MockSearchIndexLister is a mock of SearchIndexLister interface.
type MockSearchIndexLister struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexListerMockRecorder
}

// MockSearchIndexListerMockRecorder is the mock recorder for MockSearchIndexLister.
type MockSearchIndexListerMockRecorder struct {
	mock *MockSearchIndexLister
}

// NewMockSearchIndexLister creates a new mock instance.
func NewMockSearchIndexLister(ctrl *gomock.Controller) *MockSearchIndexLister {
	mock := &MockSearchIndexLister{ctrl: ctrl}
	mock.recorder = &MockSearchIndexListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchIndexLister) EXPECT() *MockSearchIndexListerMockRecorder {
	return m.recorder
}

// SearchIndexes mocks base method.
func (m *MockSearchIndexLister) SearchIndexes(arg0, arg1, arg2, arg3 string) ([]admin.ClusterSearchIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIndexes", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]admin.ClusterSearchIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIndexes indicates an expected call of SearchIndexes.
func (mr *MockSearchIndexListerMockRecorder) SearchIndexes(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIndexes", reflect.TypeOf((*MockSearchIndexLister)(nil).SearchIndexes), arg0, arg1, arg2, arg3)
}

// MockSearchIndexCreator is a mock of SearchIndexCreator interface.
type MockSearchIndexCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexCreatorMockRecorder
}

// MockSearchIndexCreatorMockRecorder is the mock recorder for MockSearchIndexCreator.
type MockSearchIndexCreatorMockRecorder struct {
	mock *MockSearchIndexCreator
}

// NewMockSearchIndexCreator creates a new mock instance.
func NewMockSearchIndexCreator(ctrl *gomock.Controller) *MockSearchIndexCreator {
	mock := &MockSearchIndexCreator{ctrl: ctrl}
	mock.recorder = &MockSearchIndexCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchIndexCreator) EXPECT() *MockSearchIndexCreatorMockRecorder {
	return m.recorder
}

// CreateSearchIndexes mocks base method.
func (m *MockSearchIndexCreator) CreateSearchIndexes(arg0, arg1 string, arg2 *admin.ClusterSearchIndex) (*admin.ClusterSearchIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSearchIndexes", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ClusterSearchIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSearchIndexes indicates an expected call of CreateSearchIndexes.
func (mr *MockSearchIndexCreatorMockRecorder) CreateSearchIndexes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSearchIndexes", reflect.TypeOf((*MockSearchIndexCreator)(nil).CreateSearchIndexes), arg0, arg1, arg2)
}

// SearchIndex mocks base method.
func (m *MockSearchIndexCreator) SearchIndex(arg0, arg1, arg2 string) (*admin.ClusterSearchIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ClusterSearchIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIndex indicates an expected call of SearchIndex.
func (mr *MockSearchIndexCreatorMockRecorder) SearchIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIndex", reflect.TypeOf((*MockSearchIndexCreator)(nil).SearchIndex), arg0, arg1, arg2)
}

// MockSearchIndexDescriber is a mock of SearchIndexDescriber interface.
type MockSearchIndexDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexDescriberMockRecorder
}

// MockSearchIndexDescriberMockRecorder is the mock recorder for MockSearchIndexDescriber.
type MockSearchIndexDescriberMockRecorder struct {
	mock *MockSearchIndexDescriber
}

// NewMockSearchIndexDescriber creates a new mock instance.
func NewMockSearchIndexDescriber(ctrl *gomock.Controller) *MockSearchIndexDescriber {
	mock := &MockSearchIndexDescriber{ctrl: ctrl}
	mock.recorder = &MockSearchIndexDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchIndexDescriber) EXPECT() *MockSearchIndexDescriberMockRecorder {
	return m.recorder
}

// SearchIndex mocks base method.
func (m *MockSearchIndexDescriber) SearchIndex(arg0, arg1, arg2 string) (*admin.ClusterSearchIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ClusterSearchIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIndex indicates an expected call of SearchIndex.
func (mr *MockSearchIndexDescriberMockRecorder) SearchIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIndex", reflect.TypeOf((*MockSearchIndexDescriber)(nil).SearchIndex), arg0, arg1, arg2)
}

// MockSearchIndexUpdater is a mock of SearchIndexUpdater interface.
type MockSearchIndexUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexUpdaterMockRecorder
}

// MockSearchIndexUpdaterMockRecorder is the mock recorder for MockSearchIndexUpdater.
type MockSearchIndexUpdaterMockRecorder struct {
	mock *MockSearchIndexUpdater
}

// NewMockSearchIndexUpdater creates a new mock instance.
func NewMockSearchIndexUpdater(ctrl *gomock.Controller) *MockSearchIndexUpdater {
	mock := &MockSearchIndexUpdater{ctrl: ctrl}
	mock.recorder = &MockSearchIndexUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchIndexUpdater) EXPECT() *MockSearchIndexUpdaterMockRecorder {
	return m.recorder
}

// UpdateSearchIndexes mocks base method.
func (m *MockSearchIndexUpdater) UpdateSearchIndexes(arg0, arg1, arg2 string, arg3 *admin.ClusterSearchIndex) (*admin.ClusterSearchIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSearchIndexes", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.ClusterSearchIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSearchIndexes indicates an expected call of UpdateSearchIndexes.
func (mr *MockSearchIndexUpdaterMockRecorder) UpdateSearchIndexes(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSearchIndexes", reflect.TypeOf((*MockSearchIndexUpdater)(nil).UpdateSearchIndexes), arg0, arg1, arg2, arg3)
}

// MockSearchIndexDeleter is a mock of SearchIndexDeleter interface.
type MockSearchIndexDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexDeleterMockRecorder
}

// MockSearchIndexDeleterMockRecorder is the mock recorder for MockSearchIndexDeleter.
type MockSearchIndexDeleterMockRecorder struct {
	mock *MockSearchIndexDeleter
}

// NewMockSearchIndexDeleter creates a new mock instance.
func NewMockSearchIndexDeleter(ctrl *gomock.Controller) *MockSearchIndexDeleter {
	mock := &MockSearchIndexDeleter{ctrl: ctrl}
	mock.recorder = &MockSearchIndexDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchIndexDeleter) EXPECT() *MockSearchIndexDeleterMockRecorder {
	return m.recorder
}

// DeleteSearchIndex mocks base method.
func (m *MockSearchIndexDeleter) DeleteSearchIndex(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSearchIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSearchIndex indicates an expected call of DeleteSearchIndex.
func (mr *MockSearchIndexDeleterMockRecorder) DeleteSearchIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSearchIndex", reflect.TypeOf((*MockSearchIndexDeleter)(nil).DeleteSearchIndex), arg0, arg1, arg2)
}
