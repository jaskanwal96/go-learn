package main

import (
	"context"
	"log"
	"net"

	hellopb "challenge1_2/pb"

	"google.golang.org/grpc"
)

// Server struct
type server struct {
	hellopb.UnimplementedGreeterServer
}

// Implement the RPC
func (s *server) SayHello(
	ctx context.Context,
	req *hellopb.HelloRequest,
) (*hellopb.HelloReply, error) {

	return &hellopb.HelloReply{
		Message: "Hello, " + req.Name + "!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	hellopb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server running on :50051")
	grpcServer.Serve(lis)
}
