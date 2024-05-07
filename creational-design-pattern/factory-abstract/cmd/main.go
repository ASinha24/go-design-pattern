package main

import (
	"fmt"

	sf "github.com/asinha24/go-design-pattern/creational-design-pattern/factory-abstract/sports-factory"
)

func main() {
	adidasFactory, _ := sf.NewSportsFactory("adidas")
	nikeFactory, _ := sf.NewSportsFactory("nike")

	// adidasShirt := adidasFactory.MakeShirt()
	// adidasShoe := adidasFactory.MakeShoe()

	// nikeShirt := nikeFactory.MakeShirt()
	// nikeShoe := nikeFactory.MakeShoe()

	printShirtDetails(adidasFactory)
	printShirtDetails(nikeFactory)

}

func printShirtDetails(s sf.ISportsfactory) {
	fmt.Printf("Logo: %s", s.MakeShirt().GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.MakeShoe().GetSize())
	fmt.Println()
}
