package parser

import "sync"

var (
	registry = map[string]Parser{}
	mu       sync.RWMutex
)

func Register(p Parser) {
	mu.Lock()
	defer mu.Unlock()
	for _, t := range p.SupportedTypes() {
		registry[t] = p
	}
}

func GetByMime(mime string) Parser {
	mu.RLock()
	defer mu.RUnlock()
	return registry[mime]
}
