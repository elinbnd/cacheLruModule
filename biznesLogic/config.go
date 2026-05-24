package biznesLogic
type CacheRule struct {
	Domain string
	Path string
	TTL int
	ErrorTTL int
	RedirectTTL int
	Cache3xx bool
	Cache4xx bool
	Cache5xx bool
	MinSize int
	MaxSize int
	Priority int
}
type CacheConfig struct {
	Global CacheRule
	Rules []CacheRule
}