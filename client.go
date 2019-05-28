package main

import (
	"context"
	"drcache-client/consistent_hashing"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "drcache-client/grpc"
)

var (
	servers = map[string]struct{}{"localhost:50051": {}}
)

func add(c map[string]pb.DrcacheClient, item pb.Item, ring *consistent_hashing.Ring) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client := c[ring.Get(item.Key)]
	r, err := client.Add(ctx, &pb.AddRequest{Item: &item})
	return r, err
}

func get(c map[string]pb.DrcacheClient, key string, ring *consistent_hashing.Ring) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c[ring.Get(key)].Get(ctx, &pb.GetRequest{Key: key})
	return r, err
}

func delete(c pb.DrcacheClient, key string) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.Delete(ctx, &pb.DeleteRequest{Key: key})
	return r, err
}

func deleteAll(c pb.DrcacheClient) (*pb.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.DeleteAll(ctx, &pb.DeleteAllRequest{})
	return r, err
}

func main() {

	ring := consistent_hashing.NewRing(servers)
	clients := make(map[string]pb.DrcacheClient)

	for address := range servers {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		c := pb.NewDrcacheClient(conn)
		clients[address] = c
	}

	item1 := pb.Item{Key: "demokey", Value: []byte("demovalue"), LastUpdate: 1, Expiration: 100}
	r, err := add(clients, item1, ring)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("key: %s", r.Message)

	r1, err1 := get(clients, "demokey", ring)
	if err != nil {
		log.Fatalf("Error %v", err1)
	}
	log.Printf("Item={ %s: %s}", r1.Item.Key, r1.Item.Value)

}
