package tests

import ("testing"
"github.com/elinbnd/cacheLruModule/biznesLogic")

func TestShouldCacheTwoXRange(test *testing.T) {
	cfg := biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{
			TTL: 10,
		},
	}
	p := biznesLogic.NewPolicy(cfg)
	ok, ttl := p.ShouldCache(200, 100, "", "/")
	if !ok || ttl != 10 {
		test.Fatal("2x Range should cache (TestShouldCacheTwoXRange)")
	}
}
func TestShouldCacheFourXRange(test *testing.T) {
	cfg := biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{
			Cache4xx: true,
			ErrorTTL: 20,
		},
	}
	p := biznesLogic.NewPolicy(cfg)
	ok, ttl := p.ShouldCache(404, 100, "", "/")
	if !ok || ttl != 20 {
		test.Fatal("4x Range should cache (TestShouldCache4xx)")
	}
}
func TestShouldCacheSizeFilter(test *testing.T) {
	cfg := biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{
			TTL:     10,
			MinSize: 50,
		},
	}
	p := biznesLogic.NewPolicy(cfg)
	ok, _ := p.ShouldCache(200, 10, "", "/")
	if ok {
		test.Fatal("small body should not cache")
	}
}
func TestResolvePriority(test *testing.T) {
	cfg := biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{
			TTL:      5,
			Priority: 1,
		},
		Rules: []biznesLogic.CacheRule{
			{
				Path:     "/api",
				TTL:      50,
				Priority: 10,
			},
		},
	}
	p := biznesLogic.NewPolicy(cfg)
	ok, ttl := p.ShouldCache(200, 100, "", "/api/test")
	if !ok || ttl != 50 {
		test.Fatal("priority broken (TestResolvePriority)")
	}
}
func TestShouldCache3xx(test *testing.T) {
	cfg := biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{
			Cache3xx:    true,
			RedirectTTL: 5,
		},
	}
	p := biznesLogic.NewPolicy(cfg)
	ok, ttl := p.ShouldCache(302, 100, "", "/")
	if !ok || ttl != 5 {
		test.Fatal("3xx cache broken (TestShouldCache3xx)")
	}
}
func TestShouldCache5xx(test *testing.T) {
	cfg := biznesLogic.CacheConfig{
		Global: biznesLogic.CacheRule{
			Cache5xx: true,
			ErrorTTL: 7,
		},
	}
	p := biznesLogic.NewPolicy(cfg)
	ok, ttl := p.ShouldCache(500, 100, "", "/")
	if !ok || ttl != 7 {
		test.Fatal("5xx cache broken (TestShouldCache5xx)")
	}
}
