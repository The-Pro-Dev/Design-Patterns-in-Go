package MethodChaining

type NumberContainer struct {
	value float64
}

func (container NumberContainer) add(num float64) NumberContainer {
	container.value += num
	return container
}

func (container NumberContainer) sub(num float64) NumberContainer {
	container.value -= num
	return container
}

func (container NumberContainer) mul(num float64) NumberContainer {
	container.value *= num
	return container
}

func (container NumberContainer) div(num float64) NumberContainer {
	container.value /= num
	return container
}
