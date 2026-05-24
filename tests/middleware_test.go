package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elinbnd/cacheLruModule/biznesLogic"
	"github.com/elinbnd/cacheLruModule/cache"
	"github.com/elinbnd/cacheLruModule/middleware"
)

func TestNoCacheHeader(test *testing.T) {
	c := cache.CreateLRUCache(10)
	cfg := biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{
			TTL: 10,
		},
	}
	policy := biznesLogic.NewPolicy(cfg)
	handler := middleware.Middleware(c, policy)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		w.Write([]byte("data"))
	}))
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	_, ok := c.GetCache("GET:/test")
	if ok {
		test.Error("should not cache")
	}
}
func TestMiddlewareCacheHit(test *testing.T) {
	c := cache.CreateLRUCache(10)
	c.PutLRU("GET:/test", cache.Item{
		Status: 200,
		Body:   []byte("cached"),
	}, 10)
	handler := middleware.Middleware(c, biznesLogic.NewPolicy(biznesLogic.CacheConfig{}))(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			test.Fatal("handler should not execute (TestMiddlewareCacheHit)")
		}),
	)
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Body.String() != "cached" {
		test.Fatal("cache miss (TestMiddlewareCacheHit)")
	}
}

func TestMiddlewareAuthorizationSkip(test *testing.T) {
	c := cache.CreateLRUCache(10)
	handler := middleware.Middleware(c, biznesLogic.NewPolicy(biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{TTL: 10},
	}))(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		}),
	)
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res, ok := c.GetCache("GET:/test")
	if ok {
		test.Fatal("authorized request should not cache (TestMiddlewareAuthorizationSkip) ", res)
	}
}
func TestMiddlewareCookieSkip(t *testing.T) {
	c := cache.CreateLRUCache(10)
	handler := middleware.Middleware(c, biznesLogic.NewPolicy(biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{TTL: 10},
	}))(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		}),
	)
	req := httptest.NewRequest("GET", "/cookie", nil)
	req.Header.Set("Cookie", "session=123")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res, ok := c.GetCache("GET:/cookie")
	if ok {
		t.Fatal("cookie request should not cache (TestMiddlewareCookieSkip) ", res)
	}
}
