package arithmetic

import (
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedArithmeticServer
}

func (s *Server) Operation(ctx context.Context, req *OperationRequest) (*OperationResponse, error) {
	var operationResult int64 = 0
	switch req.OperationType {
	case OperationType_ADD:
		operationResult = req.OperandOne + req.OperandTwo
	case OperationType_SUB:
		operationResult = req.OperandOne - req.OperandTwo
	}
	return &OperationResponse{Result: operationResult}, nil
}
