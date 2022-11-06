package DecoratorPattern

import (
	"errors"
	"reflect"
	"strconv"
)

type Stencil struct {
	code     string `default:"bird"`
	material Material
	rate     float64 `default:"1.25"`
}

// Ensure Default Values are filled before accessing
func set_default_stencil(stencil *Stencil) {
	typ := reflect.TypeOf(*stencil)

	if stencil.code == "" {
		field, _ := typ.FieldByName("code")
		stencil.code = field.Tag.Get("default")
	}

	if stencil.rate == 0 {
		field, _ := typ.FieldByName("rate")
		stencil.rate, _ = strconv.ParseFloat(field.Tag.Get("default"), 64)
	}
}

func (stencil Stencil) getCode() string {
	set_default_stencil(&stencil)
	return stencil.code
}

func (stencil Stencil) getCost(dimension float64) (float64, error) {
	set_default_stencil(&stencil)

	if stencil.material == nil {
		return -1, errors.New("material required for painting")
	}

	material_cost, _ := stencil.material.getCost(dimension)

	return material_cost + dimension*stencil.rate, nil
}
