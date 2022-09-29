package main

type TomatoTopping struct {
	doner IDoner
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.doner.getPrice()
	return pizzaPrice + 7
}
