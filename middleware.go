package cache

import "net/http"

func Middleware(c *Cache, p *Policy) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Method + ":" + r.URL.String()
			if item, ok := c.Get(key); ok {
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
				next.ServeHTTP(w, r)
				return
			}
			cookie := r.Header.Get("Cookie")
			if cookie != "" {
				next.ServeHTTP(w, r)
				return
			}
			if cacheControl == "no-cache" {
				next.ServeHTTP(w, r)
				return
			}
			ok, ttl := p.ShouldCache(rec.Status, size, r.Host, r.URL.Path)
			if ok {
				headers := make(map[string]string)
				for k := range rec.Header() {
					headers[k] = rec.Header().Get(k)
				}
				if rec.Header().Get("Cache-Control") == "no-cache" {
					return
				}
				c.Set(key, Item{
					Status:  rec.Status,
					Body:    rec.Body,
					Headers: headers,
					Tags:    []string{"default"},
				}, ttl)
			}
		})
	}
}
