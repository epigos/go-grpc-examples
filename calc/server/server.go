package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/epigos/grpc-course/calc/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calcpb.CalculatorRequest) (*calcpb.CalculatorResponse, error) {
	log.Printf("Greet function was invoked with %v\n", req)
	op1 := req.GetCalculate().GetOperandOne()
	op2 := req.GetCalculate().GetOperandTwo()
	result := op1 + op2
	res := &calcpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Starting calculator service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
