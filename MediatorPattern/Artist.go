package MediatorPattern

type Artist interface {
	perform()
}

type Actor struct {
	Name string
}

func (a Actor) perform() {
	println(a.Name, "is acting")
}

type Singer struct {
	Name string
}

func (a Singer) perform() {
	println(a.Name, "is singing")
}
