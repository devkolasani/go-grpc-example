package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/devkolasani/go-grpc-example/arithmetic"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := arithmetic.NewArithmeticClient(conn)

	newOperationRequest := &arithmetic.OperationRequest{
		OperandOne:    64,
		OperandTwo:    36,
		OperationType: arithmetic.OperationType_ADD,
	}
	log.Printf("Calling Arithmetic service with request: %s", newOperationRequest)
	response, err := c.Operation(context.Background(), newOperationRequest)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %v", response.Result)

}
