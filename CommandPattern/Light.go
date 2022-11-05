package CommandPattern

import "reflect"

type LightState string

const (
	ON  = LightState("on")
	OFF = LightState("off")
)

type Light struct {
	state LightState `default:"on"`
}

// Ensure Default Values are filled before accessing
func set_default(light *Light) {
	typ := reflect.TypeOf(*light)

	if light.state == "" {
		field, _ := typ.FieldByName("state")
		light.state = LightState(field.Tag.Get("default"))
	}
}

func (light *Light) turnOn() {
	light.state = ON
}

func (light *Light) turnOff() {
	light.state = OFF
}

func (light *Light) print() {
	set_default(light)
	println("The light is currently", light.state)
}
