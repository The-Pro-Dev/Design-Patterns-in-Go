package CommandPattern

import "fmt"

func Main() {
	fmt.Println("** Command Pattern **")

	light := Light{}
	light.print()

	lightOnCommand := LightOnCommand{&light}
	Executor(lightOnCommand)
	light.print()

	lightOffCommand := LightOffCommand{&light}
	Executor(lightOffCommand)
	light.print()
}
