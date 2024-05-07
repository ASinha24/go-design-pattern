package sportsfactory

import (
	"errors"
)

// Abstract Factory Interface
type ISportsfactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func NewSportsFactory(brand string) (ISportsfactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, errors.New("wrong brand type passed")
}
