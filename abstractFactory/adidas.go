package main

type adidas struct{}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 8,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 28,
		},
	}
}