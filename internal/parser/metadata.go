package parser

type Metadata map[string]string

func (m Metadata) Set(key, val string) {
	m[key] = val
}

func (m Metadata) Get(key string) string {
	return m[key]
}
