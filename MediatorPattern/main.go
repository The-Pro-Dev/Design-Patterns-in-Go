package MediatorPattern

import "fmt"

func Main() {
	fmt.Println("** Mediator Pattern **")

	h := Actor{Name: "Emraan Hashmi"}
	a := Singer{Name: "Atif Aslam"}

	ch := ChoreoGrapher{}
	ch.add(h)
	ch.add(a)

	ch.perform()
}
