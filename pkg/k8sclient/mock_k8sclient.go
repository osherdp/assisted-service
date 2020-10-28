// Code generated by MockGen. DO NOT EDIT.
// Source: k8sclient.go

// Package k8sclient is a generated GoMock package.
package k8sclient

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/config/v1"
	v10 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockK8SClient is a mock of K8SClient interface
type MockK8SClient struct {
	ctrl     *gomock.Controller
	recorder *MockK8SClientMockRecorder
}

// MockK8SClientMockRecorder is the mock recorder for MockK8SClient
type MockK8SClientMockRecorder struct {
	mock *MockK8SClient
}

// NewMockK8SClient creates a new mock instance
func NewMockK8SClient(ctrl *gomock.Controller) *MockK8SClient {
	mock := &MockK8SClient{ctrl: ctrl}
	mock.recorder = &MockK8SClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockK8SClient) EXPECT() *MockK8SClientMockRecorder {
	return m.recorder
}

// GetConfigMap mocks base method
func (m *MockK8SClient) GetConfigMap(namespace, name string) (*v10.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", namespace, name)
	ret0, _ := ret[0].(*v10.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap
func (mr *MockK8SClientMockRecorder) GetConfigMap(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockK8SClient)(nil).GetConfigMap), namespace, name)
}

// GetClusterVersion mocks base method
func (m *MockK8SClient) GetClusterVersion(name string) (*v1.ClusterVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterVersion", name)
	ret0, _ := ret[0].(*v1.ClusterVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterVersion indicates an expected call of GetClusterVersion
func (mr *MockK8SClientMockRecorder) GetClusterVersion(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterVersion", reflect.TypeOf((*MockK8SClient)(nil).GetClusterVersion), name)
}

// ListNodes mocks base method
func (m *MockK8SClient) ListNodes() (*v10.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes")
	ret0, _ := ret[0].(*v10.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes
func (mr *MockK8SClientMockRecorder) ListNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockK8SClient)(nil).ListNodes))
}
