package invalidation
import "github.com/elinbnd/cacheLruModule/cache"
type EventsHandler struct {
	TypeEvent string 
	KeyEvent string 
	TagEvent string 
}
var AllEventsChannel = make(chan EventsHandler, 100)
func StartEventsHandler(cache *cache.LRU) {
	go func() {
		for handlerEvent := range AllEventsChannel {
			if handlerEvent.TypeEvent == "delete_key" {
				cache.DeleteFromLRU(handlerEvent.KeyEvent)
			} else if handlerEvent.TypeEvent == "delete_tag" {
				cache.DeleteForTag(handlerEvent.TagEvent)
			} else if handlerEvent.TypeEvent == "clear" {
				cache.ClearCache()
			}
		} 
	}()
}