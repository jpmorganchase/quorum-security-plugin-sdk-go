// Code generated by MockGen. DO NOT EDIT.
// Source: proto/init.pb.go

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	grpc "google.golang.org/grpc"
)

// MockPluginInitializerClient is a mock of PluginInitializerClient interface
type MockPluginInitializerClient struct {
	ctrl     *gomock.Controller
	recorder *MockPluginInitializerClientMockRecorder
}

// MockPluginInitializerClientMockRecorder is the mock recorder for MockPluginInitializerClient
type MockPluginInitializerClientMockRecorder struct {
	mock *MockPluginInitializerClient
}

// NewMockPluginInitializerClient creates a new mock instance
func NewMockPluginInitializerClient(ctrl *gomock.Controller) *MockPluginInitializerClient {
	mock := &MockPluginInitializerClient{ctrl: ctrl}
	mock.recorder = &MockPluginInitializerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPluginInitializerClient) EXPECT() *MockPluginInitializerClientMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockPluginInitializerClient) Init(ctx context.Context, in *proto.PluginInitialization_Request, opts ...grpc.CallOption) (*proto.PluginInitialization_Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Init", varargs...)
	ret0, _ := ret[0].(*proto.PluginInitialization_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (mr *MockPluginInitializerClientMockRecorder) Init(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockPluginInitializerClient)(nil).Init), varargs...)
}

// MockPluginInitializerServer is a mock of PluginInitializerServer interface
type MockPluginInitializerServer struct {
	ctrl     *gomock.Controller
	recorder *MockPluginInitializerServerMockRecorder
}

// MockPluginInitializerServerMockRecorder is the mock recorder for MockPluginInitializerServer
type MockPluginInitializerServerMockRecorder struct {
	mock *MockPluginInitializerServer
}

// NewMockPluginInitializerServer creates a new mock instance
func NewMockPluginInitializerServer(ctrl *gomock.Controller) *MockPluginInitializerServer {
	mock := &MockPluginInitializerServer{ctrl: ctrl}
	mock.recorder = &MockPluginInitializerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPluginInitializerServer) EXPECT() *MockPluginInitializerServerMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockPluginInitializerServer) Init(arg0 context.Context, arg1 *proto.PluginInitialization_Request) (*proto.PluginInitialization_Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(*proto.PluginInitialization_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (mr *MockPluginInitializerServerMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockPluginInitializerServer)(nil).Init), arg0, arg1)
}
