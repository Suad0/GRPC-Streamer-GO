package service

import (
	"io"
	"os"
	"time"

	pb "github.com/Suad0/GrpcStreamer/api/proto/generated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VideoStreamingService struct {
	pb.UnimplementedVideoStreamingServer
}

func (s *VideoStreamingService) StreamVideo(req *pb.VideoRequest, stream pb.VideoStreaming_StreamVideoServer) error {
	videoPath := "internal/video/" + req.VideoId + ".mp4"

	print("StreamVideo wird aufgerufen")
	print(req.VideoId)

	file, err := os.Open(videoPath)
	if err != nil {
		return status.Errorf(codes.NotFound, "Video not found: %v", err)
	}
	defer file.Close()

	buffer := make([]byte, 1024) // 1 KB buffer for video chunks
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "Error reading video: %v", err)
		}

		chunk := &pb.VideoChunk{
			Data:      buffer[:n],
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(chunk); err != nil {
			return status.Errorf(codes.Internal, "Error sending chunk: %v", err)
		}
	}

	return nil
}
