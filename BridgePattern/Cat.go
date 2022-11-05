package BridgePattern

type Cat struct {
	name  string
	breed string
	fur   Fur
}

func (cat Cat) getName() string {
	return cat.name
}

func (cat Cat) getBreed() string {
	return cat.breed
}

func (cat Cat) getColors() string {
	return cat.fur.getColors()
}
