package main

import (
	"context"
	"log"
	"os"

	pb "github.com/Suad0/GrpcStreamer/api/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewVideoStreamingClient(conn)
	req := &pb.VideoRequest{VideoId: "example_video.mp4"}

	stream, err := client.StreamVideo(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to start streaming: %v", err)
	}

	file, err := os.Create("received_video.mp4")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive chunk: %v", err)
		}
		if len(chunk.Data) == 0 {
			break
		}
		_, err = file.Write(chunk.Data)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
	}

	log.Println("Video received successfully")
}
