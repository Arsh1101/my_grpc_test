package main

import (
	"fmt"
	"io"
	"log"
	"net"

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
	var data []byte
	for {
		chunk, err := stream.Recv()
		fmt.Println(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		data = append(data, chunk.Data...)
	}

	// TODO: Save the received file to disk

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
