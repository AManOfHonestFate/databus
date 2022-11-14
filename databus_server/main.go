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

// does given operation on 2 floats
func operation(op string, a float32, b float32) (float32, error) {
	var res float32	
	
	switch op {
		case "add":
			res = a + b
			fmt.Printf("%v + %v = %v\n", a, b, res)
		case "sub":
			res = a - b
			fmt.Printf("%v - %v = %v\n", a, b, res)
		case "mul":
			res = a * b
			fmt.Printf("%v * %v = %v\n", a, b, res)
		case "div":
			res = a / b
			fmt.Printf("%v / %v = %v\n", a, b, res)
		default:
			return 0, errors.New("unsupported operation")
	}

	return res, nil
}

// Send implementation
func (s* server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	a, b := in.GetPrm1(), in.GetPrm2()

	res, err := operation(os.Args[2], a, b)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &pb.SendResponse{Result: res}, nil
}

func main() {
	// requires 2 arguments: port, operation
	if len(os.Args) != 3 {
		log.Fatalf("requires 2 arguments, given %v", len(os.Args))
	}
	
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

