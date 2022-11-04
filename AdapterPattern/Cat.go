package AdapterPattern

type Cat struct {
	name string
}

func (cat Cat) meow() string {
	return "Meow!"
}

type CatAdapter struct {
	cat Cat
}

func (catAdapter CatAdapter) name() string {
	return catAdapter.cat.name
}

func (catAdapter CatAdapter) kind() string {
	return "cat"
}

func (catAdapter CatAdapter) sound() string {
	return catAdapter.cat.meow()
}
