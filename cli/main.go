package main

import (
	"github.com/0xTheProDev/Design-Patterns-in-Go/AdapterPattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/BridgePattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/BuilderPattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/CommandPattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/DecoratorPattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/FacadePattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/FactoryPattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/FlyweightPattern"
	"github.com/0xTheProDev/Design-Patterns-in-Go/MethodChaining"
	"github.com/0xTheProDev/Design-Patterns-in-Go/SingletonPattern"
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

	println()

	SingletonPattern.Main()
}
