package AdapterPattern

type AnimalAdapter interface {
	name() string
	kind() string
	sound() string
}
