package pikpakapi

import "sync"

type cache struct {
	m sync.Map
}

type tuple[T1 comparable, T2 comparable] struct {
	T1 T1
	T2 T2
}

func newTuple[T1 comparable, T2 comparable](t1 T1, t2 T2) tuple[T1, T2] {
	return tuple[T1, T2]{
		t1,
		t2,
	}
}

func newCache() *cache {
	return &cache{
		m: sync.Map{},
	}
}

func (c *cache) Get(key tuple[string, string]) (string, bool) {
	v, ok := c.m.Load(key)
	if !ok {
		return "", false
	}
	return v.(string), ok
}

func (c *cache) Set(key tuple[string, string], value string) {
	c.m.Store(key, value)
}
