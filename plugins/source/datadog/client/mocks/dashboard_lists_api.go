// Code generated by MockGen. DO NOT EDIT.
// Source: dashboard_lists_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	datadogV1 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	gomock "github.com/golang/mock/gomock"
)

// MockDashboardListsAPIClient is a mock of DashboardListsAPIClient interface.
type MockDashboardListsAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockDashboardListsAPIClientMockRecorder
}

// MockDashboardListsAPIClientMockRecorder is the mock recorder for MockDashboardListsAPIClient.
type MockDashboardListsAPIClientMockRecorder struct {
	mock *MockDashboardListsAPIClient
}

// NewMockDashboardListsAPIClient creates a new mock instance.
func NewMockDashboardListsAPIClient(ctrl *gomock.Controller) *MockDashboardListsAPIClient {
	mock := &MockDashboardListsAPIClient{ctrl: ctrl}
	mock.recorder = &MockDashboardListsAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDashboardListsAPIClient) EXPECT() *MockDashboardListsAPIClientMockRecorder {
	return m.recorder
}

// ListDashboardLists mocks base method.
func (m *MockDashboardListsAPIClient) ListDashboardLists(arg0 context.Context) (datadogV1.DashboardListListResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDashboardLists", arg0)
	ret0, _ := ret[0].(datadogV1.DashboardListListResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDashboardLists indicates an expected call of ListDashboardLists.
func (mr *MockDashboardListsAPIClientMockRecorder) ListDashboardLists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDashboardLists", reflect.TypeOf((*MockDashboardListsAPIClient)(nil).ListDashboardLists), arg0)
}
