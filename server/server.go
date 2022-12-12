package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/devkolasani/go-grpc-example/arithmetic"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v \n", err)
	}

	s := arithmetic.Server{}

	grpcServerInstance := grpc.NewServer()

	arithmetic.RegisterArithmeticServer(grpcServerInstance, &s)

	if err := grpcServerInstance.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC over port 9000: %v \n", err)
	}
}
