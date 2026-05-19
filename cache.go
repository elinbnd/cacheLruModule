package cache
import (
	"sync"
	"time"
)
type Item struct {
	Status int
	Body []byte
	Headers map[string]string
	Expire int64
	Tags []string
}
type Cache struct {
	data map[string]Item
	mutex sync.RWMutex
    ttl int
	maxPutItems int 
}
func NewCache(ttl int, maxPutItems int) *Cache {
	return &Cache{
		data: make(map[string]Item),
        ttl: ttl,
		maxPutItems: maxPutItems,
	}
}
func (cache *Cache) Get(key string) (Item, bool) {
	cache.mutex.RLock()
	item, ok := cache.data[key]
	cache.mutex.RUnlock()
	if !ok {
		return Item{}, false
	}
	if item.Expire < time.Now().Unix() {
		cache.mutex.Lock()
		delete(cache.data, key)
		cache.mutex.Unlock()
		return Item{}, false
	}
	return item, true
}
func (cache *Cache) Set(key string, item Item, ttl int) {
    if ttl == 0 {
        ttl = cache.ttl
    }
	if len(cache.data) >= cache.maxPutItems {
		return
	}
    item.Expire = time.Now().Unix() + int64(ttl)
    cache.mutex.Lock()
    cache.data[key] = item
    cache.mutex.Unlock()
}





func (cache *Cache) DeleteCache(k string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	delete(cache.data, k)
}
func(cache *Cache) ClearCache() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.data = make(map[string]Item)
}
func (cache *Cache) DeleteExpired() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	nowTime := time.Now().Unix()
	for key, item := range cache.data {
		if item.Expire < nowTime {
			delete(cache.data, key)
		}
	}
}