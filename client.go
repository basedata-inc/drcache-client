package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "drcacheClient/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func add(c pb.DrcacheClient) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{Item: &pb.Item{Key: "asd", Value: []byte("1111"), LastUpdate: 1, Expiration: 9}})
	return r, err
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	c := pb.NewDrcacheClient(conn)
	r, err := add(c)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}
