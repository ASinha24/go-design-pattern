package cachealgorithms

import "fmt"

type Fifo struct{}

func (f *Fifo) Evict(c *Cache) {
	fmt.Println("Evicting using fifo eviction algorithm")
}
