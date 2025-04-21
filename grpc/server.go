package main

import (
	"log"
	"net"

	"github.com/im-vishalanand/paper-social-backend/grpc/proto/postpb"
	"github.com/im-vishalanand/paper-social-backend/grpc/service"
	"google.golang.org/grpc"
)


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	postpb.RegisterPostServiceServer(grpcServer, &service.PostServiceServerImpl{})

	log.Println("✅ gRPC PostService is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
