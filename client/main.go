package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/go-grpc/filetransfer"
)

const (
	address = "localhost:50051"
	thePath = "/var/log/vigilant-guard/"
)

func fileStreaming(filePath string, filename string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileTransferClient(conn)
	// Open the file to send.
	file, err := os.Open(filePath + filename)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	// Create a stream to send the file chunks.
	stream, err := c.SendFile(context.Background())
	if err != nil {
		log.Fatalf("could not send file: %v", err)
	}

	// Send the filename and file data to the server.
	chunkSize := 1024 // 1KB
	buf := make([]byte, chunkSize)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file: %v", err)
		}
		if err := stream.Send(&pb.SendFileRequest{Filename: filename, Data: buf[:n]}); err != nil {
			log.Fatalf("could not send chunk: %v", err)
		}
	}

	// Close the stream and wait for the response from the server.
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}
	log.Printf("Response: %v", resp)
}

func main() {
	//Get Logs:
	files, err := os.ReadDir(thePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		fileStreaming(thePath, file.Name())
	}
	fmt.Println("Done!")
}
