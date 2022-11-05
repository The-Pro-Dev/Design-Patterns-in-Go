package CommandPattern

type LightOnCommand struct {
	light *Light
}

func (cmd LightOnCommand) execute() {
	cmd.light.turnOn()
}
