package tests
import ("testing" 
 "github.com/elinbnd/cacheLruModule/cache"

 )

func TestDeletePrefix(test *testing.T) {
	c := cache.NewCache(10, 100)
	c.Set("api:user", cache.Item{}, 10)
	c.Set("api:admin", cache.Item{}, 10)
	cache.DeletePrefix(c, "api:")
	res, ok := c.Get("api:user")
	if ok {
		test.Fatal("prefix delete fail (TestDeletePrefix)", res)
	}
}
func TestDeleteRegex(test *testing.T) {
	c := cache.NewCache(10, 100)
	c.Set("user:1", cache.Item{}, 10)
	cache.DeleteRegex(c, `user:\d+`)
	res, ok := c.Get("user:1")
	if ok {
		test.Fatal("regex delete fail (TestDeleteRegex) ", res)
	}
}
func TestDeleteByTag(test *testing.T) {
	c := cache.NewCache(10, 100)
	c.Set("k1", cache.Item{
		Tags: []string{"users"},
	}, 10)
	cache.DeleteByTag(c, "users")
	res, ok := c.Get("k1")
	if ok {
		test.Fatal("tag delete fail (TestDeleteByTag) ", res)
	}
}
func TestClear(test *testing.T) {
	c := cache.NewCache(10, 100)
	c.Set("k1", cache.Item{}, 10)
	c.Clear()
	res, ok := c.Get("k1")
	if ok {
		test.Fatal("clear fail (TestClear) ", res)
	}
}
