package FactoryPattern

type PizzaStore struct {
}

func (store *PizzaStore) create(kind PizzaKind) {
	var pizza Pizza

	switch kind {
	case CheesePizza:
		pizza = createCheesePizza()
	case ClamPizza:
		pizza = createClamPizza()
	case PepperoniPizza:
		pizza = createPepperoniPizza()
	case VeggiePizza:
		pizza = createVeggiePizza()
	}

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	println()
}
