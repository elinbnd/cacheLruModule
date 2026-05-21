package cache

import (
	"fmt"
	"log"
	"net/http"
)

func Middleware(lru *LRU, p *Policy) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Method + ":" + r.URL.String()
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
		
			rec := NewRecorder(w)
			next.ServeHTTP(rec, r)
			size := len(rec.Body)
			cacheControl := r.Header.Get("Cache-Control")
			auth := r.Header.Get("Authorization")
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
			fmt.Println("AUTH ", auth)
			fmt.Println("COOKIE ",cookie)
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
				lru.PutLRU(key, Item{
					Status:  rec.Status,
					Body:    rec.Body,
					Headers: headers,
					Tags:    []string{"default"},
				}, ttl)
			}
		})
	}
}
