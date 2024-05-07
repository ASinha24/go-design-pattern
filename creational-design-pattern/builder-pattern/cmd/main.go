package main

import (
	"fmt"

	bp "github.com/asinha24/go-design-pattern/creational-design-pattern/builder-pattern"
)

func main() {
	normalBuilder := bp.NewIBuilder("normal")

	director := bp.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse("wooden window", "wooden door", 2)

	fmt.Println("normal house", normalHouse)

	iglooBuilder := bp.NewIBuilder("igloo")

	director = bp.NewDirector(iglooBuilder)
	iglooHouse := director.BuildHouse("snow window", "snow door", 1)

	fmt.Println("igloo house", iglooHouse)
}
