package cache
import (
	"container/list"
	"strings"
	"sync"
	"time"
)
type Elem struct {
	key string 
	itemInBox Item
}
type LRU struct {
	info map[string]*list.Element
	lst *list.List
	lruCap int 
	mutex sync.Mutex
}
func CreateLRUCache(lruCap int) *LRU {
	return &LRU{
		info: make(map[string]*list.Element),
		lst: list.New(),
		lruCap: lruCap,
	}
}
func (cache *LRU) GetCache(key string) (Item, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	element, ok := cache.info[key]
	if !ok {
		return Item{}, false
	}
	elem := element.Value.(Elem)
	if elem.itemInBox.Expire < time.Now().Unix() {
		cache.lst.Remove(element)
		delete(cache.info, key)
		return Item{}, false
	}
	cache.lst.MoveToFront(element)
	return elem.itemInBox, true
}
func (cache *LRU) PutLRU(key string, item Item, ttl int) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	item.Expire = time.Now().Unix() + int64(ttl)
	// item.Expire = time.Now().Unix() + int(ttl)
	if element, ok := cache.info[key]; ok {
		element.Value = Elem {
			key: key,
			itemInBox: item,
		}
		cache.lst.MoveToFront(element)
		return
	}
	newElemInCache := cache.lst.PushFront(Elem{
		key: key,
		itemInBox: item,
	})
	cache.info[key] = newElemInCache
	if cache.lst.Len() > cache.lruCap {
		oldElement := cache.lst.Back()
		if oldElement != nil {
			oldElementValue := oldElement.Value.(Elem)
			delete(cache.info, oldElementValue.key)
			cache.lst.Remove(oldElement)
		}
	}
}
func (cache *LRU) DeleteFromLRU(kDelete string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	element, ok := cache.info[kDelete]
	if !ok {
		return
	}
	cache.lst.Remove(element)
	delete(cache.info, kDelete)
}
func (cache *LRU) ClearCache() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.info = make(map[string]*list.Element)
	cache.lst.Init()
}
func (cache *LRU) PrefDelete(prefixStr string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for key, element := range cache.info {
		if strings.HasPrefix(key, prefixStr) {
			cache.lst.Remove(element)
			delete(cache.info, key)
		}
	}
}
func (cache *LRU) DeleteForTag(str string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for key, elem := range cache.info {
		cur := elem.Value.(Elem)
		for _, current := range cur.itemInBox.Tags {
			if current == str {
				cache.lst.Remove(elem)
				delete(cache.info, key)
				break
			}
		}
	}
}
func (cache *LRU) DeleteExpired() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	nowTime := time.Now().Unix()
	for key, elem := range cache.info {
		cur := elem.Value.(Elem)
		if cur.itemInBox.Expire < nowTime {
			cache.lst.Remove(elem)
			delete(cache.info, key)
		}
	}
}