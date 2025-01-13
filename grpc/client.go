package grpc

import (
	"gc2-p3-xiowel/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// ConnectGRPCClient untuk menginisialisasi dan mengembalikan client GRPC
func ConnectGRPCClient() pb.BookServiceClient {
	// Membuka koneksi ke GRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to GRPC server: %v", err)
	}
	// Mengembalikan client untuk service BookService
	return pb.NewBookServiceClient(conn)
}
