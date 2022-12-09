// Code generated by MockGen. DO NOT EDIT.
// Source: util/init.go

// Package mockUtil is a generated GoMock package.
package mockUtil

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUtil is a mock of Util interface.
type MockUtil struct {
	ctrl     *gomock.Controller
	recorder *MockUtilMockRecorder
}

// MockUtilMockRecorder is the mock recorder for MockUtil.
type MockUtilMockRecorder struct {
	mock *MockUtil
}

// NewMockUtil creates a new mock instance.
func NewMockUtil(ctrl *gomock.Controller) *MockUtil {
	mock := &MockUtil{ctrl: ctrl}
	mock.recorder = &MockUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUtil) EXPECT() *MockUtilMockRecorder {
	return m.recorder
}

// ReadUnmarshalJSON mocks base method.
func (m *MockUtil) ReadUnmarshalJSON(srcPath string, v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUnmarshalJSON", srcPath, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadUnmarshalJSON indicates an expected call of ReadUnmarshalJSON.
func (mr *MockUtilMockRecorder) ReadUnmarshalJSON(srcPath, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUnmarshalJSON", reflect.TypeOf((*MockUtil)(nil).ReadUnmarshalJSON), srcPath, v)
}

// ReadUnmarshalYAML mocks base method.
func (m *MockUtil) ReadUnmarshalYAML(srcPath string, v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUnmarshalYAML", srcPath, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadUnmarshalYAML indicates an expected call of ReadUnmarshalYAML.
func (mr *MockUtilMockRecorder) ReadUnmarshalYAML(srcPath, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUnmarshalYAML", reflect.TypeOf((*MockUtil)(nil).ReadUnmarshalYAML), srcPath, v)
}

// WriteMarshalJSON mocks base method.
func (m *MockUtil) WriteMarshalJSON(destPath string, v interface{}, prettyFormat bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMarshalJSON", destPath, v, prettyFormat)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMarshalJSON indicates an expected call of WriteMarshalJSON.
func (mr *MockUtilMockRecorder) WriteMarshalJSON(destPath, v, prettyFormat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMarshalJSON", reflect.TypeOf((*MockUtil)(nil).WriteMarshalJSON), destPath, v, prettyFormat)
}

// WriteMarshalYAML mocks base method.
func (m *MockUtil) WriteMarshalYAML(destPath string, v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMarshalYAML", destPath, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMarshalYAML indicates an expected call of WriteMarshalYAML.
func (mr *MockUtilMockRecorder) WriteMarshalYAML(destPath, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMarshalYAML", reflect.TypeOf((*MockUtil)(nil).WriteMarshalYAML), destPath, v)
}
