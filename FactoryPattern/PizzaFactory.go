package FactoryPattern

type PizzaKind string

const (
	CheesePizza    = PizzaKind("Cheese")
	ClamPizza      = PizzaKind("Clam")
	PepperoniPizza = PizzaKind("Pepperoni")
	VeggiePizza    = PizzaKind("Veggie")
)

func createCheesePizza() Pizza {
	return Pizza{
		name:     "NY Style Sauce and Cheese Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara sauce",
		toppings: []Topping{"Grated Reggiano Cheese"},
	}
}

func createClamPizza() Pizza {
	return Pizza{
		name:     "NY Style Sauce and Clam Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara sauce",
		toppings: []Topping{"Grated Reggiano Cheese", "Clams"},
	}
}

func createPepperoniPizza() Pizza {
	return Pizza{
		name:     "NY Style Sauce and Pepperoni Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara sauce",
		toppings: []Topping{"Grated Reggiano Cheese", "Pepperoni"},
	}
}

func createVeggiePizza() Pizza {
	return Pizza{
		name:     "NY Style Sauce and Veggie Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara sauce",
		toppings: []Topping{"Grated Reggiano Cheese", "Paprika", "Tomato", "Beans"},
	}
}
