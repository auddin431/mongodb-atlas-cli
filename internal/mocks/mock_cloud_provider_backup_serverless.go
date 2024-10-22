// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: ServerlessSnapshotsLister,ServerlessSnapshotsDescriber,ServerlessRestoreJobsLister,ServerlessRestoreJobsDescriber,ServerlessRestoreJobsCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20240805005/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockServerlessSnapshotsLister is a mock of ServerlessSnapshotsLister interface.
type MockServerlessSnapshotsLister struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessSnapshotsListerMockRecorder
}

// MockServerlessSnapshotsListerMockRecorder is the mock recorder for MockServerlessSnapshotsLister.
type MockServerlessSnapshotsListerMockRecorder struct {
	mock *MockServerlessSnapshotsLister
}

// NewMockServerlessSnapshotsLister creates a new mock instance.
func NewMockServerlessSnapshotsLister(ctrl *gomock.Controller) *MockServerlessSnapshotsLister {
	mock := &MockServerlessSnapshotsLister{ctrl: ctrl}
	mock.recorder = &MockServerlessSnapshotsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessSnapshotsLister) EXPECT() *MockServerlessSnapshotsListerMockRecorder {
	return m.recorder
}

// ServerlessSnapshots mocks base method.
func (m *MockServerlessSnapshotsLister) ServerlessSnapshots(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*admin.PaginatedApiAtlasServerlessBackupSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessSnapshots", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasServerlessBackupSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessSnapshots indicates an expected call of ServerlessSnapshots.
func (mr *MockServerlessSnapshotsListerMockRecorder) ServerlessSnapshots(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessSnapshots", reflect.TypeOf((*MockServerlessSnapshotsLister)(nil).ServerlessSnapshots), arg0, arg1, arg2)
}

// MockServerlessSnapshotsDescriber is a mock of ServerlessSnapshotsDescriber interface.
type MockServerlessSnapshotsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessSnapshotsDescriberMockRecorder
}

// MockServerlessSnapshotsDescriberMockRecorder is the mock recorder for MockServerlessSnapshotsDescriber.
type MockServerlessSnapshotsDescriberMockRecorder struct {
	mock *MockServerlessSnapshotsDescriber
}

// NewMockServerlessSnapshotsDescriber creates a new mock instance.
func NewMockServerlessSnapshotsDescriber(ctrl *gomock.Controller) *MockServerlessSnapshotsDescriber {
	mock := &MockServerlessSnapshotsDescriber{ctrl: ctrl}
	mock.recorder = &MockServerlessSnapshotsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessSnapshotsDescriber) EXPECT() *MockServerlessSnapshotsDescriberMockRecorder {
	return m.recorder
}

// ServerlessSnapshot mocks base method.
func (m *MockServerlessSnapshotsDescriber) ServerlessSnapshot(arg0, arg1, arg2 string) (*admin.ServerlessBackupSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ServerlessBackupSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessSnapshot indicates an expected call of ServerlessSnapshot.
func (mr *MockServerlessSnapshotsDescriberMockRecorder) ServerlessSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessSnapshot", reflect.TypeOf((*MockServerlessSnapshotsDescriber)(nil).ServerlessSnapshot), arg0, arg1, arg2)
}

// MockServerlessRestoreJobsLister is a mock of ServerlessRestoreJobsLister interface.
type MockServerlessRestoreJobsLister struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessRestoreJobsListerMockRecorder
}

// MockServerlessRestoreJobsListerMockRecorder is the mock recorder for MockServerlessRestoreJobsLister.
type MockServerlessRestoreJobsListerMockRecorder struct {
	mock *MockServerlessRestoreJobsLister
}

// NewMockServerlessRestoreJobsLister creates a new mock instance.
func NewMockServerlessRestoreJobsLister(ctrl *gomock.Controller) *MockServerlessRestoreJobsLister {
	mock := &MockServerlessRestoreJobsLister{ctrl: ctrl}
	mock.recorder = &MockServerlessRestoreJobsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessRestoreJobsLister) EXPECT() *MockServerlessRestoreJobsListerMockRecorder {
	return m.recorder
}

// ServerlessRestoreJobs mocks base method.
func (m *MockServerlessRestoreJobsLister) ServerlessRestoreJobs(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*admin.PaginatedApiAtlasServerlessBackupRestoreJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessRestoreJobs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasServerlessBackupRestoreJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessRestoreJobs indicates an expected call of ServerlessRestoreJobs.
func (mr *MockServerlessRestoreJobsListerMockRecorder) ServerlessRestoreJobs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessRestoreJobs", reflect.TypeOf((*MockServerlessRestoreJobsLister)(nil).ServerlessRestoreJobs), arg0, arg1, arg2)
}

// MockServerlessRestoreJobsDescriber is a mock of ServerlessRestoreJobsDescriber interface.
type MockServerlessRestoreJobsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessRestoreJobsDescriberMockRecorder
}

// MockServerlessRestoreJobsDescriberMockRecorder is the mock recorder for MockServerlessRestoreJobsDescriber.
type MockServerlessRestoreJobsDescriberMockRecorder struct {
	mock *MockServerlessRestoreJobsDescriber
}

// NewMockServerlessRestoreJobsDescriber creates a new mock instance.
func NewMockServerlessRestoreJobsDescriber(ctrl *gomock.Controller) *MockServerlessRestoreJobsDescriber {
	mock := &MockServerlessRestoreJobsDescriber{ctrl: ctrl}
	mock.recorder = &MockServerlessRestoreJobsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessRestoreJobsDescriber) EXPECT() *MockServerlessRestoreJobsDescriberMockRecorder {
	return m.recorder
}

// ServerlessRestoreJob mocks base method.
func (m *MockServerlessRestoreJobsDescriber) ServerlessRestoreJob(arg0, arg1, arg2 string) (*admin.ServerlessBackupRestoreJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessRestoreJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ServerlessBackupRestoreJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessRestoreJob indicates an expected call of ServerlessRestoreJob.
func (mr *MockServerlessRestoreJobsDescriberMockRecorder) ServerlessRestoreJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessRestoreJob", reflect.TypeOf((*MockServerlessRestoreJobsDescriber)(nil).ServerlessRestoreJob), arg0, arg1, arg2)
}

// MockServerlessRestoreJobsCreator is a mock of ServerlessRestoreJobsCreator interface.
type MockServerlessRestoreJobsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessRestoreJobsCreatorMockRecorder
}

// MockServerlessRestoreJobsCreatorMockRecorder is the mock recorder for MockServerlessRestoreJobsCreator.
type MockServerlessRestoreJobsCreatorMockRecorder struct {
	mock *MockServerlessRestoreJobsCreator
}

// NewMockServerlessRestoreJobsCreator creates a new mock instance.
func NewMockServerlessRestoreJobsCreator(ctrl *gomock.Controller) *MockServerlessRestoreJobsCreator {
	mock := &MockServerlessRestoreJobsCreator{ctrl: ctrl}
	mock.recorder = &MockServerlessRestoreJobsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessRestoreJobsCreator) EXPECT() *MockServerlessRestoreJobsCreatorMockRecorder {
	return m.recorder
}

// ServerlessCreateRestoreJobs mocks base method.
func (m *MockServerlessRestoreJobsCreator) ServerlessCreateRestoreJobs(arg0, arg1 string, arg2 *admin.ServerlessBackupRestoreJob) (*admin.ServerlessBackupRestoreJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessCreateRestoreJobs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ServerlessBackupRestoreJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessCreateRestoreJobs indicates an expected call of ServerlessCreateRestoreJobs.
func (mr *MockServerlessRestoreJobsCreatorMockRecorder) ServerlessCreateRestoreJobs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessCreateRestoreJobs", reflect.TypeOf((*MockServerlessRestoreJobsCreator)(nil).ServerlessCreateRestoreJobs), arg0, arg1, arg2)
}
