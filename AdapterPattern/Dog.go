package AdapterPattern

type Dog struct {
	name string
}

func (dog Dog) bark() string {
	return "Bark!"
}

type DogAdapter struct {
	dog Dog
}

func (dogAdapter DogAdapter) name() string {
	return dogAdapter.dog.name
}

func (dogAdapter DogAdapter) kind() string {
	return "dog"
}

func (dogAdapter DogAdapter) sound() string {
	return dogAdapter.dog.bark()
}
