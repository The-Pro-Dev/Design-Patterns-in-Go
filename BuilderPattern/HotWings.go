package BuilderPattern

type HotWings struct {
}

func (hotWings HotWings) getName() string {
	return "Hot Wings"
}

func (hotWings HotWings) getPrice() float64 {
	return 70.00
}
