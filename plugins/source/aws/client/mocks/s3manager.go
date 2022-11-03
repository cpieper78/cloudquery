// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client/services (interfaces: S3managerClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

// MockS3managerClient is a mock of S3managerClient interface.
type MockS3managerClient struct {
	ctrl     *gomock.Controller
	recorder *MockS3managerClientMockRecorder
}

// MockS3managerClientMockRecorder is the mock recorder for MockS3managerClient.
type MockS3managerClientMockRecorder struct {
	mock *MockS3managerClient
}

// NewMockS3managerClient creates a new mock instance.
func NewMockS3managerClient(ctrl *gomock.Controller) *MockS3managerClient {
	mock := &MockS3managerClient{ctrl: ctrl}
	mock.recorder = &MockS3managerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3managerClient) EXPECT() *MockS3managerClientMockRecorder {
	return m.recorder
}

// GetBucketRegion mocks base method.
func (m *MockS3managerClient) GetBucketRegion(arg0 context.Context, arg1 string, arg2 ...func(*s3.Options)) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketRegion", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketRegion indicates an expected call of GetBucketRegion.
func (mr *MockS3managerClientMockRecorder) GetBucketRegion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketRegion", reflect.TypeOf((*MockS3managerClient)(nil).GetBucketRegion), varargs...)
}