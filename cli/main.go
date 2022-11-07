package main

import (
	"github.com/Progyan1997/Design-Patterns-in-Go/AdapterPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/BridgePattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/BuilderPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/CommandPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/DecoratorPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/Facade"
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

	MethodChaining.Main()

	println()

	Facade.Main()
}
