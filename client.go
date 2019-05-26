package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "drcacheClient/grpc"
)

const (
	address = "localhost:50051"
)

func add(c pb.DrcacheClient, item pb.Item) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{Item: &item})
	return r, err
}

func get(c pb.DrcacheClient) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.GetRequest{})
	return r, err
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	c := pb.NewDrcacheClient(conn)

	item := pb.Item{Key: "qwe", Value: []byte("111331"), LastUpdate: 1, Expiration: 9}
	r, err := add(c, item)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}
