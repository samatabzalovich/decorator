package main

type CheeseTopping struct {
	doner IDoner
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.doner.getPrice()
	return pizzaPrice + 10
}
