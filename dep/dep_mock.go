// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=dep_mock.go -package=dep
//

// Package dep is a generated GoMock package.
package dep

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDepService is a mock of DepService interface.
type MockDepService struct {
	ctrl     *gomock.Controller
	recorder *MockDepServiceMockRecorder
	isgomock struct{}
}

// MockDepServiceMockRecorder is the mock recorder for MockDepService.
type MockDepServiceMockRecorder struct {
	mock *MockDepService
}

// NewMockDepService creates a new mock instance.
func NewMockDepService(ctrl *gomock.Controller) *MockDepService {
	mock := &MockDepService{ctrl: ctrl}
	mock.recorder = &MockDepServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepService) EXPECT() *MockDepServiceMockRecorder {
	return m.recorder
}

// GetUserAttendanceInfo mocks base method.
func (m *MockDepService) GetUserAttendanceInfo(userId int) (*UserAttendanceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAttendanceInfo", userId)
	ret0, _ := ret[0].(*UserAttendanceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAttendanceInfo indicates an expected call of GetUserAttendanceInfo.
func (mr *MockDepServiceMockRecorder) GetUserAttendanceInfo(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAttendanceInfo", reflect.TypeOf((*MockDepService)(nil).GetUserAttendanceInfo), userId)
}

// ProcessOrder mocks base method.
func (m *MockDepService) ProcessOrder(orderId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessOrder", orderId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessOrder indicates an expected call of ProcessOrder.
func (mr *MockDepServiceMockRecorder) ProcessOrder(orderId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessOrder", reflect.TypeOf((*MockDepService)(nil).ProcessOrder), orderId)
}