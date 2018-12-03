package main

import (
	"context"
	"fmt"
	"log"

	"github.com/epigos/grpc-course/calc/calcpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()
	c := calcpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(cc calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC...")

	// reader := bufio.NewReader(os.Stdin)
	var op1, op2 int32
	fmt.Print("Enter first operand: ")
	fmt.Scanf("%d", &op1)
	fmt.Print("Enter second operand: ")
	fmt.Scanf("%d", &op2)

	req := &calcpb.CalculatorRequest{
		Calculate: &calcpb.Calculator{
			OperandOne: op1,
			OperandTwo: op2,
		},
	}

	res, err := cc.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Result)
}
