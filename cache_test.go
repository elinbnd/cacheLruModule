package cache

import (
	"testing"
	"time"
)

func TestCacheBasic(test *testing.T) {
	c := NewCache(10, 100)
	c.Set("key", Item{Status: 1}, 1)
	item, ok := c.Get("key")
	if !ok || item.Status != 1 {
		test.Error("cache fail")
	}
}
func TestCacheExpire(test *testing.T) {
	c := NewCache(1, 100)
	c.Set("key", Item{Status: 1}, 1)
	time.Sleep(2 * time.Second)
	_, ok := c.Get("key")
	if ok {
		test.Error("should expire")
	}
}
func TestDelete(test *testing.T) {
	c := NewCache(10, 100)
	c.Set("k1", Item{}, 10)
	c.Delete("k1")
	_, ok := c.Get("k1")
	if ok {
		test.Error("delete fail")
	}
}
func TestCacheMaxItems(test *testing.T) {
	c := NewCache(10, 1)
	c.Set("k1", Item{}, 10)
	c.Set("k2", Item{}, 10)
	res, ok := c.Get("k2")
	if ok {
		test.Fatal("max items limit broken (TestCacheMaxItems) ", res)
	}
}
func TestDeleteCache(test *testing.T) {
	c := NewCache(10, 10)
	c.Set("k1", Item{}, 10)
	c.DeleteCache("k1")
	res, ok := c.Get("k1")
	if ok {
		test.Fatal("delete cache fail (TestDeleteCache) ", res)
	}
}
func TestClearCache(test *testing.T) {
	c := NewCache(10, 10)
	c.Set("k1", Item{}, 10)
	c.ClearCache()
	res, ok := c.Get("k1")
	if ok {
		test.Fatal("clear cache fail (TestClearCache) ", res)
	}
}
func TestFindChangesDelete(test *testing.T) {
	lru := CreateLRUCache(10)

	lru.PutLRU("k1", Item{}, 10)

	FindChanges(lru, "k1", "old", "new")

	res, ok := lru.GetCache("k1")
	if ok {
		test.Fatal("should delete changed item (TestFindChangesDelete) ", res)
	}
}
func TestFindChangesNoDelete(test *testing.T) {
	lru := CreateLRUCache(10)
	lru.PutLRU("k1", Item{}, 10)
	FindChanges(lru, "k1", "same", "same")
	res, ok := lru.GetCache("k1")
	if !ok {
		test.Fatal("must not delete (TestFindChangesNoDelete) ", res)
	}
}
func TestCascade(test *testing.T) {
	lru := CreateLRUCache(10)
	lru.PutLRU("k1", Item{}, 10)
	lru.PutLRU("k2", Item{}, 10)
	Cascade(lru, []string{"k1", "k2"})
	res, ok := lru.GetCache("k1")
	if ok {
		test.Fatal("cascade failed (TestCascade) ", res)
	}
}
func TestPrefixDeleteWrapper(t *testing.T) {
	lru := CreateLRUCache(10)
	lru.PutLRU("api:1", Item{}, 10)
	PrefixDelete(lru, "api")
	res, ok := lru.GetCache("api:1")
	if ok {
		t.Fatal("prefix wrapper fail (TestPrefixDeleteWrapper) ", res)
	}
}
