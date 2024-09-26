package main

import (
	"log"
	"net"

	pb "github.com/Suad0/GrpcStreamer/api/proto/generated"
	"github.com/Suad0/GrpcStreamer/internal/service"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	videoService := &service.VideoStreamingService{}
	pb.RegisterVideoStreamingServer(s, videoService)

	log.Printf("gRPC server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
