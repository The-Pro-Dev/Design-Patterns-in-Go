package Facade

type Oven struct {
}

func (Oven) turnOn() {
	println("Turning On Oven")
}

func (Oven) turnOff() {
	println("Turning Off Oven")
}
