package api

import (
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// VideoRequest is the request sent by the client to start a video stream.
type VideoRequest struct {
	VideoId string
}

// VideoChunk represents a chunk of video data sent from the server to the client.
type VideoChunk struct {
	Data      []byte // Video data
	Timestamp int64  // Unix timestamp
}

// Methods to implement `proto.Message` interface for VideoRequest
func (vr *VideoRequest) Reset() {
	*vr = VideoRequest{}
}

func (vr *VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(vr)
}

func (*VideoRequest) ProtoMessage() {}

func (*VideoRequest) ProtoReflect() protoreflect.Message {
	return nil
}

// Methods to implement `proto.Message` interface for VideoChunk
func (vc *VideoChunk) Reset() {
	*vc = VideoChunk{}
}

func (vc *VideoChunk) String() string {
	return protoimpl.X.MessageStringOf(vc)
}

func (*VideoChunk) ProtoMessage() {}

func (*VideoChunk) ProtoReflect() protoreflect.Message {
	return nil
}

// Exporting these variables to make sure they are initialized once per type
var (
	videoRequestProtoOnce sync.Once
	videoChunkProtoOnce   sync.Once
)
