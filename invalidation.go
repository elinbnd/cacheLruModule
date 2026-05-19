package cache

import (
	"regexp"
)

func (cache *Cache) Delete(key string) {
	cache.mutex.Lock()
	delete(cache.data, key)
	cache.mutex.Unlock()
}
func (cache *Cache) DeletePrefix(prefix string) {
	cache.mutex.Lock()
	for k := range cache.data {
		if len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			delete(cache.data, k)
		}
	}
	cache.mutex.Unlock()
}
func (lru *LRU) DeleteRegex(regex string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	for k := range lru.info {
		ok, err := regexp.MatchString(regex, k)
		if err != nil {
			return
		}
		if ok {
			delete(lru.info, k)
		}
	}
}
func (cache *Cache) DeleteRegex(pattern string) {
	re := regexp.MustCompile(pattern)

	cache.mutex.Lock()
	for k := range cache.data {
		if re.MatchString(k) {
			delete(cache.data, k)
		}
	}
	cache.mutex.Unlock()
}
func (cache *Cache) DeleteByTag(tag string) {
	cache.mutex.Lock()
	for k, v := range cache.data {
		for _, t := range v.Tags {
			if t == tag {
				delete(cache.data, k)
			}
		}
	}
	cache.mutex.Unlock()
}
func (cache *Cache) Clear() {
	cache.mutex.Lock()
	cache.data = make(map[string]Item)
	cache.mutex.Unlock()
}
