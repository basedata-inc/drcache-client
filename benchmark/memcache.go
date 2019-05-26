package benchmark

import (
	"context"
	"google.golang.org/appengine/memcache"
	"log"
	"time"
)

func test_add_memcache(ctx context.Context, item memcache.Item) {
	start := time.Now()
	var _ = memcache.Add(ctx, &item)
	elapsed := time.Since(start)
	log.Printf("get memcache took %s", elapsed)
}

func test_set_memcache(ctx context.Context, item memcache.Item) {
	start := time.Now()
	var _ = memcache.Set(ctx, &item)
	elapsed := time.Since(start)
	log.Printf("get memcache took %s", elapsed)
}

func test_get_memcache(ctx context.Context, key string) {
	start := time.Now()
	var _, _ = memcache.Get(ctx, key)
	elapsed := time.Since(start)
	log.Printf("get memcache took %s", elapsed)
}

func test_delete_memcache(ctx context.Context, key string) {
	start := time.Now()
	var _ = memcache.Delete(ctx, key)
	elapsed := time.Since(start)
	log.Printf("delete memcache took %s", elapsed)
}

func test_flush_memcache(ctx context.Context) {
	start := time.Now()
	var _ = memcache.Flush(ctx)
	elapsed := time.Since(start)
	log.Printf("flush memcache took %s", elapsed)
}

func main() {

}
