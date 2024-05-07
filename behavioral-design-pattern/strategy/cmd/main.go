package main

import (
	ca "github.com/asinha24/go-design-pattern/behavioral-design-pattern/strategy/in-memory-cache"
)

func main() {
	c := ca.NewCache(&ca.Lru{})
	c.Add("a", "97")
	c.Add("b", "98")
	c.Add("c", "99")
	c.SetEvictionAlgo(&ca.Lfu{})
	c.Add("d", "100")
	c.SetEvictionAlgo(&ca.Fifo{})
	c.Add("e", "101")
}
