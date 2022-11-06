package DecoratorPattern

import "fmt"

func Main() {
	fmt.Println("** Decorator Pattern **")

	plyWood := Wood{}
	plyWoodWithRedPaint := Color{color: "Red", material: plyWood}
	plyWoodWithRedPaintAndStencil := Stencil{material: plyWoodWithRedPaint}

	cost, _ := plyWoodWithRedPaintAndStencil.getCost(12.0 * 11.0)

	fmt.Println("Cost of", plyWood.getName(), "product [12 x 11] with", plyWoodWithRedPaint.getColor(), "paint and", plyWoodWithRedPaintAndStencil.getCode(), "stencil is:", cost)
}
