package main

import "fmt"

func main() {

	doner := &concreteDoner{}

	//Add cheese topping
	donerWithCheese := &CheeseTopping{
		doner: doner,
	}

	//Add tomato topping
	donerWithCheeseAndTomato := &TomatoTopping{
		doner: donerWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", donerWithCheeseAndTomato.getPrice())
}
