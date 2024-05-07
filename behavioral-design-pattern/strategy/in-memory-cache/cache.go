package cachealgorithms

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func NewCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:      map[string]string{},
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Evict()
	}

	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}

func (c *Cache) Evict() {
	c.evictionAlgo.Evict(c)
	c.capacity--
}
