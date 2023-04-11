// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: filestream.proto

package filestream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	SendFile(ctx context.Context, opts ...grpc.CallOption) (FileService_SendFileClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) SendFile(ctx context.Context, opts ...grpc.CallOption) (FileService_SendFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], "/filestream.FileService/SendFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceSendFileClient{stream}
	return x, nil
}

type FileService_SendFileClient interface {
	Send(*FileChunk) error
	CloseAndRecv() (*FileUploadStatus, error)
	grpc.ClientStream
}

type fileServiceSendFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceSendFileClient) Send(m *FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceSendFileClient) CloseAndRecv() (*FileUploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	SendFile(FileService_SendFileServer) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) SendFile(FileService_SendFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SendFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_SendFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).SendFile(&fileServiceSendFileServer{stream})
}

type FileService_SendFileServer interface {
	SendAndClose(*FileUploadStatus) error
	Recv() (*FileChunk, error)
	grpc.ServerStream
}

type fileServiceSendFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceSendFileServer) SendAndClose(m *FileUploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceSendFileServer) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filestream.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendFile",
			Handler:       _FileService_SendFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "filestream.proto",
}
