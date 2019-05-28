package main

import (
	"drcache-client/consistent_hashing"
	"github.com/bradfitz/gomemcache/memcache"
	"google.golang.org/grpc"
	"log"
	"testing"

	pb "drcache-client/grpc"
)

func BenchmarkAdd(b *testing.B) {

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

	for n := 0; n < b.N; n++ {
		item1 := pb.Item{Key: string(n + 100), Value: []byte("11331"), LastUpdate: 1, Expiration: 100}
		_, err := add(clients, item1, ring)

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		//log.Printf("key: %s", r.Message)
	}
}

//qq
func BenchmarkGet(b *testing.B) {
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
	for n := 0; n < b.N; n++ {
		go func() {
			r, err := get(clients, string(n+100), ring)

			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("key: %s", r.Message)
		}()
	}
}

func BenchmarkGetMemcache(b *testing.B) {
	var clients [10]*memcache.Client
	for i := 0; i < 10; i++ {
		clients[i] = memcache.New("localhost:11211")
	}
	for n := 0; n < b.N; n++ {
		go func() {
			_, _ = clients[n%10].Get(string(n))
		}()
	}
}

func BenchmarkAddMemcache(b *testing.B) {
	var clients [10]*memcache.Client
	for i := 0; i < 10; i++ {
		clients[i] = memcache.New("localhost:11211")
	}
	var items [100000]*memcache.Item
	for i := 0; i < 100000; i++ {
		items[i] = &memcache.Item{
			Key:        string(i),
			Value:      []byte(string(i % 4)),
			Flags:      0,
			Expiration: 0,
		}
	}
	for n := 0; n < b.N; n++ {
		go func() {
			_ = clients[n%10].Add(&memcache.Item{
				Key:        "id",
				Value:      nil,
				Flags:      0,
				Expiration: 0,
			})
		}()
	}
}
