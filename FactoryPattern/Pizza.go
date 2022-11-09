package FactoryPattern

import "fmt"

type Dough string
type Sauce string
type Topping string

type Pizza struct {
	dough    Dough
	name     string
	sauce    Sauce
	toppings []Topping
}

func (pizza Pizza) prepare() {
	fmt.Println("Preparing", pizza.name)
	fmt.Println("Tossing dough", pizza.dough)
	fmt.Println("Adding sauce", pizza.sauce)
	fmt.Println("Adding toppings", pizza.toppings)
}

func (pizza Pizza) bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (pizza Pizza) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (pizza Pizza) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}
