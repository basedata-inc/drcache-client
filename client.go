package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "drcache-client/grpc"
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

func get(c pb.DrcacheClient, key string) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.GetRequest{Key: key})
	return r, err
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	c := pb.NewDrcacheClient(conn)

	item1 := pb.Item{Key: "q", Value: []byte("111331"), LastUpdate: 1, Expiration: 100}
	r, err := add(c, item1)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("key: %s", r.Message)

	//add(c, item2)
	r1, err1 := get(c, "q")

	if err1 != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r1.Item.Value)

}
