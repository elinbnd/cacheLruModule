package invalidation
import "github.com/elinbnd/cacheLruModule/cache"
func FindChanges(cache *cache.LRU, key string, old string, fix string) {
	if old != fix {
		cache.DeleteFromLRU(key)
	}

}