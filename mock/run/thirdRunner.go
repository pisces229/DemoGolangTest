// Code generated by MockGen. DO NOT EDIT.
// Source: demo.golang.test/run (interfaces: ThirdRunner)

// Package mock_run is a generated GoMock package.
package mock_run

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockThirdRunner is a mock of ThirdRunner interface.
type MockThirdRunner struct {
	ctrl     *gomock.Controller
	recorder *MockThirdRunnerMockRecorder
}

// MockThirdRunnerMockRecorder is the mock recorder for MockThirdRunner.
type MockThirdRunnerMockRecorder struct {
	mock *MockThirdRunner
}

// NewMockThirdRunner creates a new mock instance.
func NewMockThirdRunner(ctrl *gomock.Controller) *MockThirdRunner {
	mock := &MockThirdRunner{ctrl: ctrl}
	mock.recorder = &MockThirdRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThirdRunner) EXPECT() *MockThirdRunnerMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockThirdRunner) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockThirdRunnerMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockThirdRunner)(nil).GetName))
}

// ThirdSpeedUp mocks base method.
func (m *MockThirdRunner) ThirdSpeedUp(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ThirdSpeedUp", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// ThirdSpeedUp indicates an expected call of ThirdSpeedUp.
func (mr *MockThirdRunnerMockRecorder) ThirdSpeedUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ThirdSpeedUp", reflect.TypeOf((*MockThirdRunner)(nil).ThirdSpeedUp), arg0)
}
