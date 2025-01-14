package grpc

import (
	"gc2-p3-xiowel/internal/handler"
	internal "gc2-p3-xiowel/internal/middleware"
	"gc2-p3-xiowel/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGRPCServer() {
	// Membuat server GRPC
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(internal.JwtUnaryInterceptor),
	)

	// Mendaftarkan service ke server GRPC
	pb.RegisterBookServiceServer(grpcServer, &handler.BookServiceServer{})

	// Membuka listener di port 50051
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

	// Jalankan server GRPC
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
