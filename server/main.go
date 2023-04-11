package main

import (
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "example.com/go-grpc/filetransfer"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedFileTransferServer
}

func (s *server) SendFile(stream pb.FileTransfer_SendFileServer) error {
	// Receive the file chunks from the client
	var filename string
	var data []byte
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		filename = chunk.Filename
		data = append(data, chunk.Data...)
	}

	// Save the received file to disk with the exact filename
	file, err := os.Create("./logs/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	// Send the status back to the client
	return stream.SendAndClose(&pb.SendStatus{
		Success: true,
		Message: "File received successfully",
	})
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFileTransferServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
