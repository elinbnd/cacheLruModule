package tests
import (
	"testing"
	"time"
	"github.com/elinbnd/cacheLruModule/cache"
)
func TestLRUPutAndGet(test *testing.T) {
	lru := cache.CreateLRUCache(2)
	lru.PutLRU("k1", cache.Item{Status: 200}, 10)
	item, ok := lru.GetCache("k1")
	if !ok {
		test.Fatal("404 item not found (TestLRUPutAndGet)")
	}

	if item.Status != 200 {
		test.Fatal("wrong status error (TestLRUPutAndGet)")
	}
}
func TestLRUExpire(test *testing.T) {
	lru := cache.CreateLRUCache(2)
	lru.PutLRU("k1", cache.Item{Status: 200}, 1)
	time.Sleep(2 * time.Second)
	value, ok := lru.GetCache("k1")
	if ok {
		test.Fatal("must be exp (TestLRUExpire)", value)
	}
}
func TestLRUCapacity(test *testing.T) {
	lru := cache.CreateLRUCache(2)
	lru.PutLRU("k1", cache.Item{}, 10)
	lru.PutLRU("k2", cache.Item{}, 10)
	lru.PutLRU("k3", cache.Item{}, 10)
	res, ok := lru.GetCache("k1")
	if ok {
		test.Fatal("oldest item need deleted (TestLRUExpire)", res)
	}
}
func TestDeleteFromLRU(test *testing.T) {
	lru := cache.CreateLRUCache(2)
	lru.PutLRU("k1", cache.Item{}, 10)
	lru.DeleteFromLRU("k1")
	res, ok := lru.GetCache("k1")
	if ok {
		test.Fatal("should delete", res)
	}
}
func TestClearLRU(test *testing.T) {
	lru := cache.CreateLRUCache(2)
	lru.PutLRU("k1", cache.Item{}, 10)
	lru.ClearCache()
	res, ok := lru.GetCache("k1")
	if ok {
		test.Fatal("should be clear (TestClearLRU)", res)
	}
}