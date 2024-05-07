package cachealgorithms

import "fmt"

type Lfu struct{}

func (lf *Lfu) Evict(c *Cache) {
	fmt.Println("evicted by least frequently used eviction algorithm")
}
