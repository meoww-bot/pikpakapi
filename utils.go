package pikpakapi

type cache[K comparable, V any] struct {
	m map[K]V
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

func newCache[K comparable, V any]() *cache[K, V] {
	return &cache[K, V]{
		m: make(map[K]V),
	}
}

func (c *cache[K, V]) Get(key K) (V, bool) {
	v, ok := c.m[key]
	return v, ok
}

func (c *cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}
