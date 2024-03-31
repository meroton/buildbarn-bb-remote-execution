// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: pkg/proto/completedactionlogger/completed_action_logger.proto

package completedactionlogger

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CompletedActionLogger_LogCompletedActions_FullMethodName = "/buildbarn.completedactionlogger.CompletedActionLogger/LogCompletedActions"
)

// CompletedActionLoggerClient is the client API for CompletedActionLogger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompletedActionLoggerClient interface {
	LogCompletedActions(ctx context.Context, opts ...grpc.CallOption) (CompletedActionLogger_LogCompletedActionsClient, error)
}

type completedActionLoggerClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletedActionLoggerClient(cc grpc.ClientConnInterface) CompletedActionLoggerClient {
	return &completedActionLoggerClient{cc}
}

func (c *completedActionLoggerClient) LogCompletedActions(ctx context.Context, opts ...grpc.CallOption) (CompletedActionLogger_LogCompletedActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompletedActionLogger_ServiceDesc.Streams[0], CompletedActionLogger_LogCompletedActions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &completedActionLoggerLogCompletedActionsClient{stream}
	return x, nil
}

type CompletedActionLogger_LogCompletedActionsClient interface {
	Send(*CompletedAction) error
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type completedActionLoggerLogCompletedActionsClient struct {
	grpc.ClientStream
}

func (x *completedActionLoggerLogCompletedActionsClient) Send(m *CompletedAction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *completedActionLoggerLogCompletedActionsClient) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompletedActionLoggerServer is the server API for CompletedActionLogger service.
// All implementations should embed UnimplementedCompletedActionLoggerServer
// for forward compatibility
type CompletedActionLoggerServer interface {
	LogCompletedActions(CompletedActionLogger_LogCompletedActionsServer) error
}

// UnimplementedCompletedActionLoggerServer should be embedded to have forward compatible implementations.
type UnimplementedCompletedActionLoggerServer struct {
}

func (UnimplementedCompletedActionLoggerServer) LogCompletedActions(CompletedActionLogger_LogCompletedActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method LogCompletedActions not implemented")
}

// UnsafeCompletedActionLoggerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompletedActionLoggerServer will
// result in compilation errors.
type UnsafeCompletedActionLoggerServer interface {
	mustEmbedUnimplementedCompletedActionLoggerServer()
}

func RegisterCompletedActionLoggerServer(s grpc.ServiceRegistrar, srv CompletedActionLoggerServer) {
	s.RegisterService(&CompletedActionLogger_ServiceDesc, srv)
}

func _CompletedActionLogger_LogCompletedActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CompletedActionLoggerServer).LogCompletedActions(&completedActionLoggerLogCompletedActionsServer{stream})
}

type CompletedActionLogger_LogCompletedActionsServer interface {
	Send(*emptypb.Empty) error
	Recv() (*CompletedAction, error)
	grpc.ServerStream
}

type completedActionLoggerLogCompletedActionsServer struct {
	grpc.ServerStream
}

func (x *completedActionLoggerLogCompletedActionsServer) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *completedActionLoggerLogCompletedActionsServer) Recv() (*CompletedAction, error) {
	m := new(CompletedAction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompletedActionLogger_ServiceDesc is the grpc.ServiceDesc for CompletedActionLogger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompletedActionLogger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.completedactionlogger.CompletedActionLogger",
	HandlerType: (*CompletedActionLoggerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LogCompletedActions",
			Handler:       _CompletedActionLogger_LogCompletedActions_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/proto/completedactionlogger/completed_action_logger.proto",
}
