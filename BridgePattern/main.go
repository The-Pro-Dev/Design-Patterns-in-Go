package BridgePattern

import "fmt"

func knowCat(cat Cat) {
	fmt.Println(cat.getName(), "is a", cat.getBreed(), "cat and has a fur coat of color", cat.getColors())
}

func Main() {
	fmt.Println("** Bridge Pattern **")

	knowCat(BombayCat)
	knowCat(SiameseCat)
}
