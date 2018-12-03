package main

import (
	"context"
	"fmt"
	"io"
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

	// doUnary(c)

	doServerStreaming(c)
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

func doServerStreaming(cc calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do server streaming RPC...")

	req := &calcpb.PrimeNumberRequest{
		PrimeNumber: &calcpb.PrimeNumber{
			Number: 120,
		},
	}

	resStream, err := cc.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumberDecomposition: %v", msg.GetResult())
	}
}
