package CommandPattern

type LightOffCommand struct {
	light *Light
}

func (cmd LightOffCommand) execute() {
	cmd.light.turnOff()
}
