package cache
func FindChanges(cache *LRU, key string, old string, fix string) {
	if old != fix {
		cache.DeleteFromLRU(key)
	}

}