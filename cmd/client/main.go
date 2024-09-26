package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"

	pb "github.com/Suad0/GrpcStreamer/api/proto/generated"
)

func main() {

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	/*

		conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}



		defer conn.Close()

	*/

	client := pb.NewVideoStreamingClient(conn)

	req := &pb.VideoRequest{VideoId: "example_video"}

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
		if err == io.EOF {
			log.Println("Stream ended")
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive chunk: %v", err)
		}

		_, err = file.Write(chunk.Data)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
	}

	log.Println("Video received successfully")
}
