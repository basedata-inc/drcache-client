package main

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"time"
)

func test_add_memcache(ctx context.Context, client memcache.Client, item *memcache.Item) {
	start := time.Now()
	_ = client.Add(item)
	elapsed := time.Since(start)
	log.Printf("client %v, itemkey = %v", client.MaxIdleConns, item.Value)
	log.Printf("add memcache took %s", elapsed)
}

func test_set_memcache(ctx context.Context, client memcache.Client, item *memcache.Item) {
	start := time.Now()
	_ = client.Set(item)
	elapsed := time.Since(start)
	log.Printf("set memcache took %s", elapsed)
}

func test_get_memcache(ctx context.Context, client memcache.Client, key string) {
	start := time.Now()
	_, _ = client.Get(key)
	elapsed := time.Since(start)
	log.Printf("client %v", client.MaxIdleConns)
	log.Printf("get memcache took %s", elapsed)
}

func test_delete_memcache(ctx context.Context, client memcache.Client, key string) {
	start := time.Now()
	_ = client.Delete(key)
	elapsed := time.Since(start)
	log.Printf("delete memcache took %s", elapsed)
}

func test_flush_memcache(ctx context.Context, client memcache.Client) {
	start := time.Now()
	_ = client.FlushAll()
	elapsed := time.Since(start)
	log.Printf("flush memcache took %s", elapsed)
}

func main() {
	/*test case 1: 10000 read
	 */
	var clients [10]*memcache.Client
	var items [10000]*memcache.Item
	for i := 0; i < 10; i++ {
		clients[i] = memcache.New("localhost:11211")
	}
	for i := 0; i < 10000; i++ {
		items[i] = &memcache.Item{
			Key:        string(i),
			Value:      []byte(string(i % 4)),
			Flags:      0,
			Expiration: 0,
		}
	}

	start := time.Now()
	for i := 0; i < 10000; i++ {

		test_get_memcache(context.Background(), *clients[(i+1000)/1000-1], string(i))
	}
	elapsed := time.Since(start)
	log.Printf("Elapsed took %s", elapsed)
	log.Printf("avarage get: %v", elapsed.Seconds()/10000)
}
