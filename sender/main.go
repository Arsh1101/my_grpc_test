package main

import (
	"context"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"

	pb "example.com/go-grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	file, err := os.Open("path/to/your/file")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stream, err := client.SendFile(context.Background())
	if err != nil {
		log.Fatalf("Failed to send file: %v", err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		if err := stream.Send(&pb.FileChunk{Chunk: buf[:n]}); err != nil {
			log.Fatalf("Failed to send chunk: %v", err)
		}
	}

	status, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive status: %v", err)
	}
	log.Printf("Upload complete: %v", status.Success)
}
