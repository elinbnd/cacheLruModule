package cache
func PrefixDelete(cache *LRU, prefixStr string) {
	cache.PrefDelete(prefixStr)
}