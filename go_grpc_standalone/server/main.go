package main

import (
	"log"
	"net"
	chat "./chat"
	"google.golang.org/grpc"
)
type Server struct {
}


func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterHelloServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
