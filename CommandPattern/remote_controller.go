package main

type remoteController struct {
	onCommands  map[string]command
	offCommands map[string]command
	preCmd      command
}

func createSimpleRemoteController() *remoteController {
	rc := &remoteController{
		onCommands:  make(map[string]command),
		offCommands: make(map[string]command),
	}
	// 这里处理为了没有找到插槽时的表现
	noCMD := new(noCommand)
	rc.onCommands[labelNoCommand] = noCMD
	rc.offCommands[labelNoCommand] = noCMD

	// 初始化的时候，没有上一次点击的按钮，所以使用 noCommand
	rc.preCmd = noCMD
	return rc
}

func (rc *remoteController) setCommand(label string, onCMD, offCMD command) {
	rc.onCommands[label] = onCMD
	rc.offCommands[label] = offCMD
}

func (rc *remoteController) onButtonWasPressed(label string) {
	cmd := rc.onCommands[label]
	if cmd == nil {
		cmd = rc.onCommands[labelNoCommand]
	}
	cmd.execute()
	rc.preCmd = cmd
}

func (rc *remoteController) offButtonWasPressed(label string) {
	cmd := rc.offCommands[label]
	if cmd == nil {
		cmd = rc.offCommands[labelNoCommand]
	}
	cmd.execute()
	rc.preCmd = cmd
}

func (rc *remoteController) undo() {
	rc.preCmd.undo()
}
