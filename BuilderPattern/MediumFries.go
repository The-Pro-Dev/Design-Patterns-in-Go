package BuilderPattern

type MediumFries struct {
}

func (mFried MediumFries) getName() string {
	return "Medium Fries"
}

func (mFried MediumFries) getPrice() float64 {
	return 90.00
}
