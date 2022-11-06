package DecoratorPattern

import (
	"errors"
	"reflect"
	"strconv"
)

type Color struct {
	color    string `default:"black"`
	material Material
	rate     float64 `default:"3.50"`
}

// Ensure Default Values are filled before accessing
func set_default_color(color *Color) {
	typ := reflect.TypeOf(*color)

	if color.color == "" {
		field, _ := typ.FieldByName("color")
		color.color = field.Tag.Get("default")
	}

	if color.rate == 0 {
		field, _ := typ.FieldByName("rate")
		color.rate, _ = strconv.ParseFloat(field.Tag.Get("default"), 64)
	}
}

func (color Color) getColor() string {
	set_default_color(&color)
	return color.color
}

func (color Color) getCost(dimension float64) (float64, error) {
	set_default_color(&color)

	if color.material == nil {
		return -1, errors.New("material required for painting")
	}

	material_cost, _ := color.material.getCost(dimension)

	return material_cost + dimension*color.rate, nil
}
