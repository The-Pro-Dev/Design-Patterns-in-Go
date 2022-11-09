package FactoryPattern

import "fmt"

func Main() {
	fmt.Println("** Factory Pattern **")

	store := PizzaStore{}

	store.create(CheesePizza)
	store.create(PepperoniPizza)
	store.create(VeggiePizza)
}
