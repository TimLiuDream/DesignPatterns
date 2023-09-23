package main

// 命令模式
func main() {
	simpleLight := new(light)
	lightOnCmd := &lightOnCommand{l: simpleLight}
	lightOffCmd := &lightOffCommand{l: simpleLight}

	simpleGarageDoor := new(garageDoor)
	doorOpenCmd := &garageDoorOpenCommand{gd: simpleGarageDoor}
	doorCloseCmd := &garageDoorCloseCommand{gd: simpleGarageDoor}

	rc := createSimpleRemoteController()
	rc.setCommand(simpleLight.label(), lightOnCmd, lightOffCmd)
	rc.setCommand(simpleGarageDoor.label(), doorOpenCmd, doorCloseCmd)

	rc.onButtonWasPressed(simpleLight.label())
	rc.offButtonWasPressed(simpleLight.label())
	rc.undo()
	rc.onButtonWasPressed(simpleGarageDoor.label())
	rc.offButtonWasPressed(simpleGarageDoor.label())
	rc.undo()

	rc.onButtonWasPressed("coffeeMachine")
}
