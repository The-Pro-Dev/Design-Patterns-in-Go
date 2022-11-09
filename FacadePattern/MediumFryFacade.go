package FacadePattern

func MediumFryFacade() {
	oven := Oven{}

	potato := Potato{"Potato"}
	potato.crush()

	oven.turnOn()
	potato.fry()
	oven.turnOff()

	bucket := Bucket{}
	bucket.fill(BucketItem(potato))
	bucket.pack()
}
