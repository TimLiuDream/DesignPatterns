package main

import "fmt"

const (
	labelNoCommand = "on_command"
)

type noCommand struct{}

func (n *noCommand) execute() {
	fmt.Println("no command executed")
}

func (n *noCommand) undo() {
	fmt.Println("no command undo")
}
