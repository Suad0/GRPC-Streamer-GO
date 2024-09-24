package api

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// VideoStreamingClient defines the client API for the VideoStreaming service.
type VideoStreamingClient interface {
	StreamVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoStreaming_StreamVideoClient, error)
}

// VideoStreamingServer defines the server API for the VideoStreaming service.
type VideoStreamingServer interface {
	StreamVideo(*VideoRequest, VideoStreaming_StreamVideoServer) error
}

// UnimplementedVideoStreamingServer should be embedded in your server implementation for forward compatibility.
type UnimplementedVideoStreamingServer struct{}

func (*UnimplementedVideoStreamingServer) StreamVideo(req *VideoRequest, srv VideoStreaming_StreamVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamVideo not implemented")
}

// RegisterVideoStreamingServer registers the VideoStreamingServer to a gRPC server.
func RegisterVideoStreamingServer(s *grpc.Server, srv VideoStreamingServer) {
	s.RegisterService(&_VideoStreaming_serviceDesc, srv)
}

// VideoStreaming_StreamVideoServer is the server-side stream interface.
type VideoStreaming_StreamVideoServer interface {
	Send(*VideoChunk) error
	grpc.ServerStream
}

// VideoStreaming_StreamVideoClient is the client-side stream interface.
type VideoStreaming_StreamVideoClient interface {
	Recv() (*VideoChunk, error)
	grpc.ClientStream
}

// Implementation for VideoStreaming client
type videoStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStreamingClient(cc grpc.ClientConnInterface) VideoStreamingClient {
	return &videoStreamingClient{cc}
}

func (c *videoStreamingClient) StreamVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoStreaming_StreamVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VideoStreaming_serviceDesc.Streams[0], "/video.VideoStreaming/StreamVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoStreamingStreamVideoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type videoStreamingStreamVideoClient struct {
	grpc.ClientStream
}

func (x *videoStreamingStreamVideoClient) Recv() (*VideoChunk, error) {
	m := new(VideoChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Service descriptor
var _VideoStreaming_serviceDesc = grpc.ServiceDesc{
	ServiceName: "video.VideoStreaming",
	HandlerType: (*VideoStreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamVideo",
			Handler:       _VideoStreaming_StreamVideo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "video_streaming.proto",
}

// Handler for StreamVideo RPC
func _VideoStreaming_StreamVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VideoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoStreamingServer).StreamVideo(m, &videoStreamingStreamVideoServer{stream})
}

type videoStreamingStreamVideoServer struct {
	grpc.ServerStream
}

func (x *videoStreamingStreamVideoServer) Send(m *VideoChunk) error {
	return x.ServerStream.SendMsg(m)
}
