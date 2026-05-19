package cache
func Cascade(cache *LRU, keys []string) {
	for i:=0; i < len(keys); i++ {
		cache.DeleteFromLRU(keys[i])
	}
}