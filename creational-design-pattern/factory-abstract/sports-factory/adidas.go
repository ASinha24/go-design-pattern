package sportsfactory

// Concrete factory
type Adidas struct{}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShirt{
		Shirt: Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
