// Code generated by MockGen. DO NOT EDIT.
// Source: ./memgopb/memgo_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=./memgopb/memgo_grpc.pb.go -destination=./memgopb/mocks/memgo_grpc.pb.mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	memgopb "memgo/memgopb"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockMemgoServiceClient is a mock of MemgoServiceClient interface.
type MockMemgoServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMemgoServiceClientMockRecorder
}

// MockMemgoServiceClientMockRecorder is the mock recorder for MockMemgoServiceClient.
type MockMemgoServiceClientMockRecorder struct {
	mock *MockMemgoServiceClient
}

// NewMockMemgoServiceClient creates a new mock instance.
func NewMockMemgoServiceClient(ctrl *gomock.Controller) *MockMemgoServiceClient {
	mock := &MockMemgoServiceClient{ctrl: ctrl}
	mock.recorder = &MockMemgoServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemgoServiceClient) EXPECT() *MockMemgoServiceClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMemgoServiceClient) Delete(ctx context.Context, in *memgopb.DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockMemgoServiceClientMockRecorder) Delete(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemgoServiceClient)(nil).Delete), varargs...)
}

// Read mocks base method.
func (m *MockMemgoServiceClient) Read(ctx context.Context, in *memgopb.ReadRequest, opts ...grpc.CallOption) (*memgopb.ReadResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*memgopb.ReadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockMemgoServiceClientMockRecorder) Read(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockMemgoServiceClient)(nil).Read), varargs...)
}

// Write mocks base method.
func (m *MockMemgoServiceClient) Write(ctx context.Context, in *memgopb.WriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Write", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockMemgoServiceClientMockRecorder) Write(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockMemgoServiceClient)(nil).Write), varargs...)
}

// MockMemgoServiceServer is a mock of MemgoServiceServer interface.
type MockMemgoServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockMemgoServiceServerMockRecorder
}

// MockMemgoServiceServerMockRecorder is the mock recorder for MockMemgoServiceServer.
type MockMemgoServiceServerMockRecorder struct {
	mock *MockMemgoServiceServer
}

// NewMockMemgoServiceServer creates a new mock instance.
func NewMockMemgoServiceServer(ctrl *gomock.Controller) *MockMemgoServiceServer {
	mock := &MockMemgoServiceServer{ctrl: ctrl}
	mock.recorder = &MockMemgoServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemgoServiceServer) EXPECT() *MockMemgoServiceServerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMemgoServiceServer) Delete(arg0 context.Context, arg1 *memgopb.DeleteRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockMemgoServiceServerMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemgoServiceServer)(nil).Delete), arg0, arg1)
}

// Read mocks base method.
func (m *MockMemgoServiceServer) Read(arg0 context.Context, arg1 *memgopb.ReadRequest) (*memgopb.ReadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*memgopb.ReadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockMemgoServiceServerMockRecorder) Read(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockMemgoServiceServer)(nil).Read), arg0, arg1)
}

// Write mocks base method.
func (m *MockMemgoServiceServer) Write(arg0 context.Context, arg1 *memgopb.WriteRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockMemgoServiceServerMockRecorder) Write(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockMemgoServiceServer)(nil).Write), arg0, arg1)
}

// mustEmbedUnimplementedMemgoServiceServer mocks base method.
func (m *MockMemgoServiceServer) mustEmbedUnimplementedMemgoServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMemgoServiceServer")
}

// mustEmbedUnimplementedMemgoServiceServer indicates an expected call of mustEmbedUnimplementedMemgoServiceServer.
func (mr *MockMemgoServiceServerMockRecorder) mustEmbedUnimplementedMemgoServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMemgoServiceServer", reflect.TypeOf((*MockMemgoServiceServer)(nil).mustEmbedUnimplementedMemgoServiceServer))
}

// MockUnsafeMemgoServiceServer is a mock of UnsafeMemgoServiceServer interface.
type MockUnsafeMemgoServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeMemgoServiceServerMockRecorder
}

// MockUnsafeMemgoServiceServerMockRecorder is the mock recorder for MockUnsafeMemgoServiceServer.
type MockUnsafeMemgoServiceServerMockRecorder struct {
	mock *MockUnsafeMemgoServiceServer
}

// NewMockUnsafeMemgoServiceServer creates a new mock instance.
func NewMockUnsafeMemgoServiceServer(ctrl *gomock.Controller) *MockUnsafeMemgoServiceServer {
	mock := &MockUnsafeMemgoServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeMemgoServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeMemgoServiceServer) EXPECT() *MockUnsafeMemgoServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedMemgoServiceServer mocks base method.
func (m *MockUnsafeMemgoServiceServer) mustEmbedUnimplementedMemgoServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMemgoServiceServer")
}

// mustEmbedUnimplementedMemgoServiceServer indicates an expected call of mustEmbedUnimplementedMemgoServiceServer.
func (mr *MockUnsafeMemgoServiceServerMockRecorder) mustEmbedUnimplementedMemgoServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMemgoServiceServer", reflect.TypeOf((*MockUnsafeMemgoServiceServer)(nil).mustEmbedUnimplementedMemgoServiceServer))
}