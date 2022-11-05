package MethodChaining

import "fmt"

func Main() {
	fmt.Println("** Method Chaining **")

	number := NumberContainer{0}.add(10).mul(3).sub(6).div(4)
	fmt.Println("Result:", number.value)
}
