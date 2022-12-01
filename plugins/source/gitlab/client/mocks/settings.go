// Code generated by MockGen. DO NOT EDIT.
// Source: settings.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gitlab "github.com/xanzy/go-gitlab"
)

// MockSettingsClient is a mock of SettingsClient interface.
type MockSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsClientMockRecorder
}

// MockSettingsClientMockRecorder is the mock recorder for MockSettingsClient.
type MockSettingsClientMockRecorder struct {
	mock *MockSettingsClient
}

// NewMockSettingsClient creates a new mock instance.
func NewMockSettingsClient(ctrl *gomock.Controller) *MockSettingsClient {
	mock := &MockSettingsClient{ctrl: ctrl}
	mock.recorder = &MockSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsClient) EXPECT() *MockSettingsClientMockRecorder {
	return m.recorder
}

// GetSettings mocks base method.
func (m *MockSettingsClient) GetSettings(options ...gitlab.RequestOptionFunc) (*gitlab.Settings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSettings", varargs...)
	ret0, _ := ret[0].(*gitlab.Settings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockSettingsClientMockRecorder) GetSettings(options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockSettingsClient)(nil).GetSettings), options...)
}