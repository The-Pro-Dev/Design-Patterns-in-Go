package main

import (
	"fmt"

	"github.com/Progyan1997/Design-Patterns-in-Go/AdapterPattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/BridgePattern"
	"github.com/Progyan1997/Design-Patterns-in-Go/BuilderPattern"
)

func main() {
	AdapterPattern.Main()

	fmt.Println()

	BridgePattern.Main()

	fmt.Println()

	BuilderPattern.Main()
}
