package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/AManOfHonestFate/databus/databus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// requires 3 arguments: target, param1, param2
	if len(os.Args) != 4 {
		log.Fatalf("requires 3 arguments, given %v", len(os.Args) - 1)
	}
	
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDatabusServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	// Parsing params
	prm1, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		log.Fatalf("didn't parse param1: %v", err)
	}
	prm2, err := strconv.ParseFloat(os.Args[3], 32)
	if err != nil {
		log.Fatalf("didn't parse param2: %v", err)
	}
	
	// RPC call
	r, err := c.Send(ctx, &pb.SendRequest{Prm1: float32(prm1), Prm2: float32(prm2)})
	if err != nil {
		log.Fatalf("RPC call failed: %v", err)
	}
	log.Println(r.GetResult())
}
