package main

import (
	"google.golang.org/grpc"
	"log"
	"testing"

	pb "drcache-client/grpc"
)


func BenchmarkAdd(b *testing.B) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//defer conn.Close()

	c := pb.NewDrcacheClient(conn)
	for n := 0; n < b.N; n++ {
		item1 := pb.Item{Key: string(n+100), Value: []byte("11331"), LastUpdate: 1, Expiration: 100}
		r, err := add(c, item1)

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("key: %s", r.Message)
	}
}
//qq
func BenchmarkGet(b *testing.B) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//defer conn.Close()

	c := pb.NewDrcacheClient(conn)
	for n := 0; n < b.N; n++ {
		go func() {
			r, err := get(c, string(n+100))

			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("key: %s", r.Message)
		}()
	}
}

//func TestAdd(t *testing.T) {go BenchmarkAdd()}