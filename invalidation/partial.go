package invalidation
import "github.com/elinbnd/cacheLruModule/cache"
func PrefixDelete(cache *cache.LRU, prefixStr string) {
	cache.PrefDelete(prefixStr)
}