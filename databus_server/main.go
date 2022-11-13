package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/AManOfHonestFate/databus/databus"
	"google.golang.org/grpc"
)

// Databus server
type server struct {
	pb.UnimplementedDatabusServiceServer
}

func (s* server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	a, b := in.GetPrm1(), in.GetPrm2()
	var res float32

	switch os.Args[2] {
		case "add":
			res = a + b
			fmt.Printf("%v + %v = %v", a, b, res)
		case "sub":
			res = a - b
			fmt.Printf("%v - %v = %v", a, b, res)
		case "mul":
			res = a * b
			fmt.Printf("%v * %v = %v", a, b, res)
		case "div":
			res = a / b
			fmt.Printf("%v / %v = %v", a, b, res)
		default:
			fmt.Printf("wrong operation: %s", os.Args[2])
			return nil, errors.New("Wrong operation")
	}

	return &pb.SendResponse{Result: res}, nil
}

func main() {
	// Listening on port given as the first argument
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Args[1]))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creating new grpc server
	s := grpc.NewServer()
	// Registering databus server
	pb.RegisterDatabusServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

