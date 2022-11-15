package main

import (
	"github.com/Progyan1997/Design-Patterns-in-Go/AdapterPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/BridgePattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/BuilderPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/CommandPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/DecoratorPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/FacadePattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/FactoryPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/FlyweightPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/MethodChaining"
)

func main() {
	AdapterPattern.Main()

	println()

	BridgePattern.Main()

	println()

	BuilderPattern.Main()

	println()

	CommandPattern.Main()

	println()

	DecoratorPattern.Main()

	println()

	FacadePattern.Main()

	println()

	FactoryPattern.Main()

	println()

	FlyweightPattern.Main()

	println()

	MethodChaining.Main()
}
