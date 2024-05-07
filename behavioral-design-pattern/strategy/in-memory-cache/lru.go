package cachealgorithms

import "fmt"

type Lru struct{}

func (lr *Lru) Evict(c *Cache) {
	fmt.Println("evecting using least recently used eviction algoritm")
}
