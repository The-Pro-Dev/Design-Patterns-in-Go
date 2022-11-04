package AdapterPattern

import "fmt"

func sound(animalAdapter AnimalAdapter) {
	fmt.Println(animalAdapter.name(), "is a", animalAdapter.kind(), "and makes sound of", animalAdapter.sound())
}

func Main() {
	fmt.Println("** Adapter Pattern **")

	cat := Cat{name: "Evelyn"}
	catAdapter := CatAdapter{cat: cat}
	sound(catAdapter)

	dog := Dog{name: "Brazzos"}
	dogAdapter := DogAdapter{dog: dog}
	sound(dogAdapter)
}
