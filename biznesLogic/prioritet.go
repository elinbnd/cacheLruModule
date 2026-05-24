package biznesLogic
import "strings"
type Policy struct {
	cfg CacheConfig
}
func NewPolicy(cfg CacheConfig) *Policy {
	return &Policy{cfg: cfg}
}
func (p *Policy) resolve(domain string, path string) CacheRule {
	best := p.cfg.Global
	for _, r := range p.cfg.Rules {
		match := true
		if r.Domain != "" && !strings.Contains(domain, r.Domain) {
			match = false
		}
		if r.Path != "" && !strings.HasPrefix(path, r.Path) {
			match = false
		}
		if match && r.Priority > best.Priority {
			best = r
		}
	}
	return best
}
func (p *Policy) ShouldCache(status int, size int, domain string, path string) (bool, int) {
	rule := p.resolve(domain, path)
	if (rule.MinSize > 0 && size < rule.MinSize) ||
   		(rule.MaxSize > 0 && size > rule.MaxSize) {
    	return false, 0
	}
	if status >= 200 && status < 300 {
		return true, rule.TTL
	}
	if status >= 300 && status < 400 && rule.Cache3xx {
		return true, rule.RedirectTTL
	}
	if status >= 400 && status < 500 && rule.Cache4xx {
		return true, rule.ErrorTTL
	}
	if status >= 500 && status < 600 && rule.Cache5xx {
		return true, rule.ErrorTTL
	}
	return false, 0
}
