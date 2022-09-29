package main

type concreteDoner struct {
}

func (p *concreteDoner) getPrice() int {
	return 15
}
