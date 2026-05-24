package cache

import (
	"regexp"
)

func Delete(cache *Cache, key string) {
	cache.mutex.Lock()
	delete(cache.data, key)
	cache.mutex.Unlock()
}
func DeletePrefix(cache *Cache, prefix string) {
	cache.mutex.Lock()
	for k := range cache.data {
		if len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			delete(cache.data, k)
		}
	}
	cache.mutex.Unlock()
}
func DeleteRegex(cache *Cache, regex string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for k := range cache.data {
		ok, err := regexp.MatchString(regex, k)
		if err != nil {
			return
		}
		if ok {
			delete(cache.data, k)
		}
	}
}

func DeleteByTag(cache *Cache, tag string) {
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
