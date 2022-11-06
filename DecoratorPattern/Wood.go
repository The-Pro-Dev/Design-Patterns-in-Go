package DecoratorPattern

import (
	"reflect"
	"strconv"
)

type Wood struct {
	rate float64 `default:"12.00"`
}

// Ensure Default Values are filled before accessing
func set_default_wood(wood *Wood) {
	typ := reflect.TypeOf(*wood)

	if wood.rate == 0 {
		field, _ := typ.FieldByName("rate")
		wood.rate, _ = strconv.ParseFloat(field.Tag.Get("default"), 64)
	}
}

func (wood Wood) getName() string {
	return "Ply Wood"
}

func (wood Wood) getCost(dimension float64) (float64, error) {
	set_default_wood(&wood)
	return dimension * wood.rate, nil
}
