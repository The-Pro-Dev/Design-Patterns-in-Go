package DecoratorPattern

type Material interface {
	getCost(dimension float64) (float64, error)
}
