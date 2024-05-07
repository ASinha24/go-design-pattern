package sportsfactory

// Concrete factory
type Nike struct{}

func (n *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			Logo: "nike",
			Size: 15,
		},
	}
}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "nike",
			Size: 15,
		},
	}
}
