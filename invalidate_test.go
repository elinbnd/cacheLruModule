package cache
import "testing"
func TestDeletePrefix(test *testing.T) {
	c := NewCache(10, 100)
	c.Set("api:user", Item{}, 10)
	c.Set("api:admin", Item{}, 10)
	c.DeletePrefix("api:")
	res, ok := c.Get("api:user")
	if ok {
		test.Fatal("prefix delete fail (TestDeletePrefix)", res)
	}
}
func TestDeleteRegex(test *testing.T) {
	c := NewCache(10, 100)
	c.Set("user:1", Item{}, 10)
	c.DeleteRegex(`user:\d+`)
	res, ok := c.Get("user:1")
	if ok {
		test.Fatal("regex delete fail (TestDeleteRegex) ", res)
	}
}
func TestDeleteByTag(test *testing.T) {
	c := NewCache(10, 100)
	c.Set("k1", Item{
		Tags: []string{"users"},
	}, 10)
	c.DeleteByTag("users")
	res, ok := c.Get("k1")
	if ok {
		test.Fatal("tag delete fail (TestDeleteByTag) ", res)
	}
}
func TestClear(test *testing.T) {
	c := NewCache(10, 100)
	c.Set("k1", Item{}, 10)
	c.Clear()
	res, ok := c.Get("k1")
	if ok {
		test.Fatal("clear fail (TestClear) ", res)
	}
}
