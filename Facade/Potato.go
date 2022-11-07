package Facade

type Potato struct {
	name string
}

func (Potato) crush() {
	println("Crushing Potato")
}

func (Potato) fry() {
	println("Frying Potato")
}
