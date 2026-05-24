package invalidation

import "github.com/elinbnd/cacheLruModule/cache"

func Cascade(cache *cache.LRU, keys []string) {
	for i := 0; i < len(keys); i++ {
		cache.DeleteFromLRU(keys[i])
	}
}