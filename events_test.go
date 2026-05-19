package cache

import (
	"testing"
	"time"
)

func TestDeleteEvent(test *testing.T) {
	lru := CreateLRUCache(10)
	StartEventsHandler(lru)
	lru.PutLRU("k1", Item{}, 10)
	AllEventsChannel <- EventsHandler{
		TypeEvent: "delete_key",
		KeyEvent:  "k1",
	}
	time.Sleep(100 * time.Millisecond)
	res, ok := lru.GetCache("k1")
	if ok {
		test.Fatal("event delete fail (TestDeleteEvent) ", res)
	}
}
