package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elinbnd/cacheLruModule/biznesLogic"
	"github.com/elinbnd/cacheLruModule/cache"
	"github.com/elinbnd/cacheLruModule/invalidation"
)

func Middleware(lru *cache.LRU, p *biznesLogic.Policy) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				next.ServeHTTP(w, r)
				return
			}
			key := r.Method + ":" + r.URL.Path + "?" + r.URL.RawQuery
			if item, ok := lru.GetCache(key); ok {
				log.Println("cache", key)
				log.Println("another ", key)
				for k, v := range item.Headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(item.Status)
				w.Write(item.Body)
				return
			}
			auth := r.Header.Get("Authorization")
			cacheControl := r.Header.Get("Cache-Control")
			if auth != "" {

				return
			}
			cookie := r.Header.Get("Cookie")
			if cookie != "" {

				return
			}
			if cacheControl == "no-cache" {

				return
			}
			rec := invalidation.NewRecorder(w)
			next.ServeHTTP(rec, r)
			size := len(rec.Body)

			fmt.Println("AUTH ", auth)
			fmt.Println("COOKIE ", cookie)
			fmt.Println("KEY ", key)
			ok, ttl := p.ShouldCache(rec.Status, size, r.Host, r.URL.Path)
			if ok {
				headers := make(map[string]string)
				for k := range rec.Header() {
					headers[k] = rec.Header().Get(k)
				}
				if rec.Header().Get("Cache-Control") == "no-cache" {
					return
				}
				lru.PutLRU(key, cache.Item{
					Status:  rec.Status,
					Body:    rec.Body,
					Headers: headers,
					Tags:    []string{"default"},
				}, ttl)
			}
		})
	}
}
