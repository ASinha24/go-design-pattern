package cachealgorithms

type EvictionAlgo interface {
	Evict(c *Cache)
}
