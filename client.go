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
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{Item: &item})
	return r, err
}

func get(c pb.DrcacheClient, key string) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.GetRequest{Key: key})
	return r, err
}

func delete(c pb.DrcacheClient, key string) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Delete(ctx, &pb.DeleteRequest{Key: key})
	return r, err
}

func deleteAll(c pb.DrcacheClient) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DeleteAll(ctx, &pb.DeleteAllRequest{})
	return r, err
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	c := pb.NewDrcacheClient(conn)

	item1 := pb.Item{Key: "qwer", Value: []byte("11199331"), LastUpdate: 1, Expiration: 100}
	r, err := add(c, item1)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("key: %s", r.Message)

	item2 := pb.Item{Key: "qwert", Value: []byte("111331"), LastUpdate: 1, Expiration: 100}
	r2, err2 := add(c, item2)

	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Printf("key: %s", r2.Message)

	//add(c, item2)
	r3, err3 := deleteAll(c)

	if err3 != nil {
		log.Fatalf("could not greet: %v", err3)
	}
	log.Printf("Greeting: %s", r3)

	r1, err1 := get(c, "qwer")

	if err1 != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r1.Item.Value)
}
